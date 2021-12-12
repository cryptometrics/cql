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
	ticker.streaming = false

	channels := []WebsocketChannel{{Name: "ticker"}}
	msg, _ := NewWebsocketMessage(products, channels)
	msg.Subscribe(conn)
	ticker.message = msg

	return ticker
}

func (ticker *AsyncTicker) makeChannels() {
	ticker.channel = make(TickerChannel)
	ticker.closed = make(chan struct{})
	ticker.closing = make(chan struct{})
}

// StartStream starts the websocket stream, streaming it into the
// AsyncTicker.channel
func (ticker *AsyncTicker) StartStream() (*AsyncTicker, error) {
	if ticker.streaming {
		return nil, fmt.Errorf("data already streaming for AsyncTicker object")
	}
	ticker.streaming = true
	ticker.makeChannels()
	ticker.Errors.Go(func() (err error) {
		defer func() {
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
	return ticker, nil
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
	if err := ticker.message.Unsubscribe(ticker.conn); err != nil {
		return err
	}
	select {
	case ticker.closing <- struct{}{}:
		<-ticker.closed
	case <-ticker.closed:
	}
	return nil
}
