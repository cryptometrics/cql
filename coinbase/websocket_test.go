package coinbase

import (
	"fmt"
	"testing"
)

func TestWebsocket(t *testing.T) {
	ws := NewProductWebsocket(DefaultWebsocketConnector)
	for ticker := range ws.Ticker("ETH-USD").Channel() {
		fmt.Printf("%+v\n", ticker)
	}
}
