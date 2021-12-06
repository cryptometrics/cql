package coinbase

type WebsocketConnector interface {
	ReadJSON(interface{}) error
	WriteJSON(interface{}) error
}

// DefaultWebsocketConnector returns a new websocket as the default connection.
func DefaultWebsocketConnector() (WebsocketConnector, error) {
	return NewWebsocket()
}
