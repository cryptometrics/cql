package coinbase

import (
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/cryptometrics/cql/model"
	. "github.com/franela/goblin"
)

func TestAsyncTickerStream(t *testing.T) {
	g := Goblin(t)
	g.Describe("mock third-party usage", func() {
		g.It("should start and close without error", func() {
			// some third-party goroutines
			mockC, _ := newmockWebsocketConnector()
			ticker := newAsyncTicker(mockC)
			numThirdParties := 2
			wg := sync.WaitGroup{}
			wg.Add(numThirdParties)
			for i := 0; i < numThirdParties; i++ {
				go func() {
					ticker.StartStream()
					r := 1 + rand.Intn(1)
					time.Sleep(time.Duration(r) * time.Second)
					ticker.Close()
					wg.Done()
				}()
			}
			wg.Wait()
		})
	})

	g.Describe("ticker#Close", func() {
		g.It("should close without error", func() {
			treshold := 100
			for i := 0; i < treshold; i++ {
				mockC, _ := newmockWebsocketConnector()
				ticker := newAsyncTicker(mockC)
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

		g.It("should closew without error on long runtime", func() {
			mockC, _ := newmockWebsocketConnector()
			ticker := newAsyncTicker(mockC)
			ticker.StartStream()
			go func() {
				tickers := []model.CoinbaseWebsocketTicker{}
				for ticker := range ticker.channel {
					tickers = append(tickers, ticker)
				}
			}()
			time.Sleep(2 * time.Second)
			ticker.Close()
			time.Sleep(2 * time.Millisecond)
		})

		g.It("should do nothing when there is no stream", func() {
			mockC, _ := newmockWebsocketConnector()
			ticker := newAsyncTicker(mockC)
			ticker.Close()
		})
	})

	g.Describe("ticker#StartStream", func() {
		g.It("should re-initialize channel data after each close", func() {
			treshold := 100
			mockC, _ := newmockWebsocketConnector()
			ticker := newAsyncTicker(mockC)
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

		g.It("should be able to start stream over again", func() {
			mockC, _ := newmockWebsocketConnector()
			ticker := newAsyncTicker(mockC)
			ticker.StartStream()
			go func() {
				tickers := []model.CoinbaseWebsocketTicker{}
				for ticker := range ticker.channel {
					tickers = append(tickers, ticker)
				}
			}()
			time.Sleep(2 * time.Microsecond)
			ticker.StartStream()
			go func() {
				tickers := []model.CoinbaseWebsocketTicker{}
				for ticker := range ticker.channel {
					tickers = append(tickers, ticker)
				}
			}()
			time.Sleep(1 * time.Microsecond)
			ticker.Close()
		})

		g.It("should not fatal if you start streams concurrently", func() {
			mockC, _ := newmockWebsocketConnector()
			ticker := newAsyncTicker(mockC)
			treshold := 1000
			for j := 0; j < treshold; j++ {
				concurrentStreams := 100
				for i := 0; i < concurrentStreams; i++ {
					go ticker.StartStream()
				}
				go func() {
					tickers := []model.CoinbaseWebsocketTicker{}
					for ticker := range ticker.channel {
						tickers = append(tickers, ticker)
					}
				}()
				time.Sleep(1 * time.Microsecond)
				ticker.Close()
			}
		})
	})
}
