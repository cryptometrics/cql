package coinbase

import (
	"context"
)

type ProductWebsocket struct {
	conn WebsocketConnector
}

// NewProductWebsocket will create a connection to the coinbase websocket and
// return a singleton that can be used to open channels that stream product
// data via a websocket.
func NewProductWebsocket(ws WebsocketCreator) *ProductWebsocket {
	productWebsocket := new(ProductWebsocket)
	productWebsocket.conn, _ = ws()
	return productWebsocket
}

// Ticker ticker uses the ProductWebsocket connection to query coinbase for
// ticket data, then it puts that data onto a channel for
// model.CoinbaseWebsocketTicker
func (productWebsocket *ProductWebsocket) Ticker(products ...string) *AsyncTicker {
	ticker := newAsyncTicker(context.Background(), productWebsocket.conn, products...)
	return ticker.startStream()
}
