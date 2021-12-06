package coinbase

import (
	"context"
	"cql/model"

	"golang.org/x/sync/errgroup"
)

type TickerChannel chan model.CoinbaseWebsocketTicker

type AsyncTicker struct {
	Errors *errgroup.Group

	channel TickerChannel
	conn    WebsocketConnector
	message *WebsocketMessage
}

func newAsyncTicker(ctx context.Context, conn WebsocketConnector, products ...string) *AsyncTicker {
	ticker := new(AsyncTicker)
	ticker.Errors, _ = errgroup.WithContext(ctx)
	ticker.channel = make(TickerChannel)
	ticker.conn = conn

	channels := []WebsocketChannel{{Name: "ticker"}}
	msg, _ := NewWebsocketMessage(products, channels)
	msg.Subscribe(conn)
	ticker.message = msg

	return ticker
}

// startStream starts the websocket stream, streaming it into the channel
func (ticker *AsyncTicker) startStream() *AsyncTicker {
	ticker.Errors.Go(func() (err error) {
		for true {
			var row model.CoinbaseWebsocketTicker
			if err = ticker.conn.ReadJSON(&row); err != nil {
				break
			}
			ticker.channel <- row
		}
		return err
	})
	return ticker
}

// Channel returns the ticker channel for streaming
func (ticker *AsyncTicker) Channel() TickerChannel {
	return ticker.channel
}

// Close unsubscribes the message from the websocket and closes the channel
func (ticker *AsyncTicker) Close() error {
	if err := ticker.message.Unsubscribe(ticker.conn); err != nil {
		return err
	}
	close(ticker.channel)
	return nil
}
