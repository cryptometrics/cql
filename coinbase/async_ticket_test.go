package coinbase

import (
	"context"
	"testing"
	"time"

	"github.com/cryptometrics/cql/model"
	. "github.com/franela/goblin"
)

func TestWebsocket(t *testing.T) {
	g := Goblin(t)
	g.Describe("AsyncTicket#Close", func() {
		g.It("The connection should close wihtout error", func() {
			// try closing 100 times
			treshold := 100
			for i := 0; i < treshold; i++ {
				mockC, _ := newmockWebsocketConnector()
				asyncTicket := newAsyncTicker(context.Background(), mockC)
				go func() {
					tickers := []model.CoinbaseWebsocketTicker{}
					for ticker := range asyncTicket.startStream().channel {
						tickers = append(tickers, ticker)
					}
				}()
				time.Sleep(1 * time.Millisecond)
				asyncTicket.Close()
			}
		})
	})
}
