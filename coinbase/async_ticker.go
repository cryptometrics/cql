package coinbase

import (
	"context"
	"fmt"

	"github.com/cryptometrics/cql/model"

	"golang.org/x/sync/errgroup"
)

type TickerChannel chan model.CoinbaseWebsocketTicker

type AsyncTicker struct {
	Errors *errgroup.Group

	channel         TickerChannel
	closed, closing chan struct{}
	conn            WebsocketConnector
	message         *WebsocketMessage
	streaming       bool
}

func newAsyncTicker(ctx context.Context, conn WebsocketConnector, products ...string) *AsyncTicker {
	ticker := new(AsyncTicker)
	ticker.Errors, _ = errgroup.WithContext(ctx)
	ticker.conn = conn

	channels := []WebsocketChannel{{Name: "ticker"}}
	msg, _ := NewWebsocketMessage(products, channels)
	msg.Subscribe(conn)
	ticker.message = msg

	return ticker
}

// StartStream starts the websocket stream, streaming it into the
// AsyncTicker.channel.  This method is idempotent, so if you call it multiple
// times successively without closing the websocket it will close the ws for you
// in each successive run and re-make the channels to stream over.
func (ticker *AsyncTicker) StartStream() *AsyncTicker {
	if ticker.streaming {
		return ticker
	}

	ticker.channel = make(TickerChannel)
	ticker.closed = make(chan struct{})
	ticker.closing = make(chan struct{})

	ticker.Errors.Go(func() (err error) {
		defer func() {
			fmt.Println("m")
			close(ticker.closed)
			close(ticker.channel)
			ticker.streaming = false
		}()
		for {
			var row model.CoinbaseWebsocketTicker
			if err = ticker.conn.ReadJSON(&row); err != nil {
				return err
			}
			select {
			case <-ticker.closing:
				return
			case ticker.channel <- row:
			}
		}
	})
	return ticker
}

// Channel returns the ticker channel for streaming
func (ticker *AsyncTicker) Channel() TickerChannel {
	return ticker.channel
}

// Close unsubscribes the message from the websocket and closes the channel.
// The Close routine can be called multiple times safely.
func (ticker *AsyncTicker) Close() error {
	if !ticker.streaming {
		return nil
	}
	// first unsubscribe from the websocket
	if err := ticker.message.Unsubscribe(ticker.conn); err != nil {
		return err
	}

	select {
	case ticker.closing <- struct{}{}:
		// wait for the stream go routine to close the 'closed' channel
		<-ticker.closed
	case <-ticker.closed:
	}
	return nil
}
