# Coinbase

`coinbase` is a package meant to be used as an SDK for creating a third-party connection in your code base to read and write information to coinbase pro using auth credentials.  For example:

```go
package alert

import (
	"github.com/cryptometrics/cql/coinbase"
)

func StartStream
	ws := coinbase.NewProductWebsocket(coinbase.DefaultWebsocketConnector)
	ticker := ws.Ticker("ETH-USD")
	ticker.StartStream()
	go func() {
		for row := range ticker.Channel() {
			fmt.Println(row.ProductId, row.Time, row.Price)
		}
	}()
}
```

## Connecting

To use nealy any of the coinbase accessors, you first need to establish a valid connection to coinbase pro using your auth credentials.  There are curretly two ways of doing this:

### .env File

You can create a .env file anywhere on your system and import the env variables into go to create a valid connection

```sh
CB_PRO_URL=
CB_PRO_ACCESS_PASSPHRASE=
CB_PRO_ACCESS_KEY=
CB_PRO_SECRET=
```

```go
// initialize the connection
conn, err := NewClientEnv("/path/to/.env")
if err != nil {
	panic(err)
}

// then use it for something
accounts := NewAccounts(conn)
allAccounts, _ := accounts.All()
fmt.Println(allAccounts)
```

### Direct Input

You can also pass the auth credentials directly

```go
// initialize the connection
_conn_, err := NewClient("url", "key", "passphrase", "secret",)
if err != nil {
	panic(err)
}
```
## WebSocket

### Async Ticker

The async ticker runs the coinbase product websocket asyncronously, connecting outside processes by sending the websocket message over a `AsyncTicket.channel`.

```go
	conn := coinbase.DefaultWebsocketConnector
	ws := coinbase.NewProductWebsocket(conn)

	// initialize the ticker object to channel product messages
	ticker := ws.Ticker("ETH-USD")

	// start a go routine that passes product messages concerning ETH-USD currency
	// pair to a channel on the ticker struct
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
	// and close the underlying channel that the messages read to.
	ticker.Close()
```

## Resources

[WebSocket](https://en.wikipedia.org/wiki/WebSocket)
