# Coinbase

## Async Ticker

The async ticker runs the coinbase product websocket asyncronously, connecting outside processes by sending the websocket message over a `AsyncTicket.channel`.

```go
	var (
		conn 		= coinbase.DefaultWebsocketConnector
		ws 			= coinbase.NewProductWebsocket(conn)
	)

	// initialize the ticker object to channel product messages
	ticker := ws.Ticker("ETH-USD")

	// StartStream this will start go routine that passes product messages
	// concerning ETH-USD currency pair to a channel on the ticker struct.
	ticker.StartStream()
	go func() {

		// Next we range over the product message channel and print the product
		// messages
		for productMsg := range ticker.Channel() {
			fmt.Println(productMsg)
		}
	}()

	// Let the product messages print for 5 seconds
	time.Sleep(5 * time.Seconds)

	// then close the ticker channel, this will unsubscribe from the websocket
	// and close the underlying channel that the messages read to..
	ticker.Close()
```
