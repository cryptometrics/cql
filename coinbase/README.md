# Coinbase

`coinbase` is a package meant to be used as an SDK for creating a third-party connection in your code base to read and write information to coinbase pro using auth credentials.  For example:

```go
func CreateOrder() {
	requestOptions := &model.CoinbaseNewOrderOptions{
		Type:        scalar.OrderTypeLimit,
		Side:        scalar.OrderSideBuy,
		Stp:         scalar.OrderSTP_DC,
		Stop:        scalar.OrderStopLoss,
		TimeInForce: scalar.TimeInForceGTC,
		CancelAfter: scalar.CancelAfterMin,
		Price:       1.0,
		Size:        10.0,
		ProductId:   "BTC-USD",
		StopPrice:   500.0,
	}
	coinbase.NewOrders(coinbase.DefaultClient).Create(requestOptions)
}
```

A slightly more complex example, using websockets:

```go
func Open() {
	ws := coinbase.NewProductWebsocket(coinbase.DefaultWebsocketConnector)
	ticker := ws.Ticker("ETH-USD")
	ticker.Open()
	go func() {
		for row := range ticker.Channel() {
			fmt.Println(row.ProductId, row.Time, row.Price)
		}
	}()

	time.Sleep(5 * time.Second)
	ticker.Close()
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
_conn_, err := NewClient("url", "key", "passphrase", "secret")
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
	ticker.Open()
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
