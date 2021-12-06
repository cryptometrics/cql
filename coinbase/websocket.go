package coinbase

import (
	"net/http"

	"github.com/gorilla/websocket"
)

// Websocket holds the connection and response data for a coinbase websocket.
type Websocket struct {
	*websocket.Conn
	Response *http.Response
}

type WebsocketCreator func() (WebsocketConnector, error)

// NewWebsocket will return a new coinbase websocket connection
func NewWebsocket() (WebsocketConnector, error) {
	conn := new(Websocket)
	url := "wss://ws-feed.pro.coinbase.com"
	var dialer websocket.Dialer
	var err error
	conn.Conn, conn.Response, err = dialer.Dial(url, nil)
	return conn, err
}
