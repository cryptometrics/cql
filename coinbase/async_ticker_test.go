package coinbase

import (
	"context"
	"testing"
	"time"

	"github.com/cryptometrics/cql/model"
	. "github.com/franela/goblin"
)

func TestAsyncTickerStream(t *testing.T) {
	g := Goblin(t)
	g.Describe("ticker#Close", func() {
		g.It("should close without error", func() {
			treshold := 100
			for i := 0; i < treshold; i++ {
				mockC, _ := newmockWebsocketConnector()
				ticker := newAsyncTicker(context.Background(), mockC)
				ticker.StartStream()
				go func() {
					tickers := []model.CoinbaseWebsocketTicker{}
					for ticker := range ticker.channel {
						tickers = append(tickers, ticker)
					}
				}()
				ticker.Close()
			}
		})

		g.It("should close wihtout error on long runtime", func() {
			mockC, _ := newmockWebsocketConnector()
			ticker := newAsyncTicker(context.Background(), mockC)
			ticker.StartStream()
			go func() {
				tickers := []model.CoinbaseWebsocketTicker{}
				for ticker := range ticker.channel {
					tickers = append(tickers, ticker)
				}
			}()
			time.Sleep(2 * time.Second)
			ticker.Close()
		})

		g.It("should do nothing when there is no stream", func() {
			mockC, _ := newmockWebsocketConnector()
			ticker := newAsyncTicker(context.Background(), mockC)
			ticker.Close()
		})
	})

	g.Describe("ticker#StartStream", func() {
		g.It("should re-initialize channel data after each close", func() {
			treshold := 100
			mockC, _ := newmockWebsocketConnector()
			ticker := newAsyncTicker(context.Background(), mockC)
			for i := 0; i < treshold; i++ {
				ticker.StartStream()
				go func() {
					tickers := []model.CoinbaseWebsocketTicker{}
					for ticker := range ticker.channel {
						tickers = append(tickers, ticker)
					}
				}()
				ticker.Close()
			}
		})

		g.It("should error if data is alreading streaming", func() {
			mockC, _ := newmockWebsocketConnector()
			ticker := newAsyncTicker(context.Background(), mockC)
			ticker.StartStream()
			go func() {
				tickers := []model.CoinbaseWebsocketTicker{}
				for ticker := range ticker.channel {
					tickers = append(tickers, ticker)
				}
			}()
			time.Sleep(1 * time.Microsecond)
			_, err := ticker.StartStream()
			g.Assert(err).IsNotNil()
			ticker.Close()
		})
	})
}
