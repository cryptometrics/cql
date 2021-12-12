package coinbase

// Create a mock websocket connection
type mockWebsocketConnector struct{}

func newmockWebsocketConnector() (WebsocketConnector, error) {
	return &mockWebsocketConnector{}, nil
}

func (conn *mockWebsocketConnector) ReadJSON(_ interface{}) error {
	return nil
}

func (conn *mockWebsocketConnector) WriteJSON(_ interface{}) error {
	return nil
}
