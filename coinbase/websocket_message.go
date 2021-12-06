package coinbase

type WebsocketMessage struct {
	Type     string             `json:"type"`
	Products []string           `json:"product_ids"`
	Channels []WebsocketChannel `json:"channels"`
}

// NewWebsocket will return a new coinbase pro websocket
// subscription that can be used to get a feed of real-time market data.
func NewWebsocketMessage(products []string, channels []WebsocketChannel) (*WebsocketMessage, error) {
	sub := new(WebsocketMessage)
	sub.Products = products
	sub.Channels = channels
	return sub, nil
}

func (msg *WebsocketMessage) isSub()   { msg.Type = "subscribe" }
func (msg *WebsocketMessage) isUnsub() { msg.Type = "unsubscribe" }

// Subscribe will use the websocket message to subscribe to a websocket
// connections, returning that connection or any errors that occured attempting
// to make it.
func (msg *WebsocketMessage) Subscribe(conn WebsocketConnector) error {
	msg.isSub()
	if err := conn.WriteJSON(msg); err != nil {
		return err
	}
	return nil
}

// Unsubscribe will use the websocket message to unsubscribe to a connection.
func (msg *WebsocketMessage) Unsubscribe(conn WebsocketConnector) error {
	msg.isUnsub()
	if err := conn.WriteJSON(msg); err != nil {
		return err
	}
	return nil
}
