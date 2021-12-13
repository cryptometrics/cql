package coinbase

import "github.com/cryptometrics/cql/websocket"

type ProductWebsocket struct {
	conn websocket.Connector
}

// NewProductWebsocket will create a connection to the coinbase websocket and
// return a singleton that can be used to open channels that stream product
// data via a websocket.
func NewProductWebsocket(ws websocket.Creator) *ProductWebsocket {
	productWebsocket := new(ProductWebsocket)
	productWebsocket.conn, _ = ws(WebsocketURL)
	return productWebsocket
}

// Ticker ticker uses the ProductWebsocket connection to query coinbase for
// ticket data, then it puts that data onto a channel for
// model.CoinbaseWebsocketTicker
func (productWebsocket *ProductWebsocket) Ticker(products ...string) *AsyncTicker {
	asyncTicker := newAsyncTicker(productWebsocket.conn, products...)
	return asyncTicker
}
