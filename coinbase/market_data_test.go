package coinbase

import (
	"bytes"
	"cql/client"
	"cql/model"
	"io/ioutil"
	"net/http"
	"testing"

	. "github.com/franela/goblin"
	"github.com/golang/mock/gomock"
)

func TestMarketDataCurrencies(t *testing.T) {
	g := Goblin(t)
	g.Describe("Test MarketData#Currencies", func() {
		test := func(desc string, g *G, ctrl *gomock.Controller, data []byte,
			expected []*model.CoinbaseMarketDataCurrency) {
			g.It(desc, func() {
				mockC := client.NewMockC(ctrl)

				mockC.EXPECT().Connect().Return(nil)

				mockC.EXPECT().Get(gomock.Any()).DoAndReturn(
					func(endpoint string) (*http.Response, error) {
						g.Assert(endpoint).Equal(CurrenciesEP.Get())

						stringReader := bytes.NewReader(data)
						stringReadCloser := ioutil.NopCloser(stringReader)
						return &http.Response{Body: stringReadCloser}, nil
					})

				var gen = func(kind client.Kind) (client.C, error) {
					return mockC, nil
				}

				md := NewMarketData()
				m, err := md.currencies(gen)
				if err != nil {
					panic(err)
				}
				g.Assert(err).IsNil()
				g.Assert(m).Eql(expected)
			})
		}

		clientCtrl := gomock.NewController(t)
		defer clientCtrl.Finish()

		var buf []byte
		var expected []*model.CoinbaseMarketDataCurrency

		buf = []byte(`[]`)
		expected = []*model.CoinbaseMarketDataCurrency{}
		test("no values in the response body", g, clientCtrl, buf, expected)

		buf = []byte(`[{"id":"BTC"}]`)
		expected = []*model.CoinbaseMarketDataCurrency{
			{ID: "BTC"},
		}
		test("one value in the response body", g, clientCtrl, buf, expected)

		buf = []byte(`[{"id":"BTC"},{"id":"ETH"}]`)
		expected = []*model.CoinbaseMarketDataCurrency{
			{ID: "BTC"}, {ID: "ETH"},
		}
		test("two values in the response body", g, clientCtrl, buf, expected)

		buf = []byte(`[{"details":{"type":"crypto"}}]`)
		expected = []*model.CoinbaseMarketDataCurrency{
			{Details: model.CoinbaseMarketDataCurrencyDetails{Type: "crypto"}},
		}
		test("details should parse correctly", g, clientCtrl, buf, expected)
	})
}

func TestMarketDataCurrency(t *testing.T) {
	g := Goblin(t)
	g.Describe("Test MarketData#Currency", func() {
		test := func(desc string, g *G, ctrl *gomock.Controller, data []byte, id string,
			expected *model.CoinbaseMarketDataCurrency) {
			g.It(desc, func() {
				mockC := client.NewMockC(ctrl)

				mockC.EXPECT().Connect().Return(nil)

				mockC.EXPECT().Get(gomock.Any()).DoAndReturn(
					func(endpoint string) (*http.Response, error) {
						g.Assert(endpoint).Equal(CurrencyEP.Get(id))

						stringReader := bytes.NewReader(data)
						stringReadCloser := ioutil.NopCloser(stringReader)
						return &http.Response{Body: stringReadCloser}, nil
					})

				var gen = func(kind client.Kind) (client.C, error) {
					return mockC, nil
				}

				md := NewMarketData()
				m, err := md.currency(gen, id)
				if err != nil {
					panic(err)
				}
				g.Assert(err).IsNil()
				g.Assert(m).Eql(expected)
			})
		}

		clientCtrl := gomock.NewController(t)
		defer clientCtrl.Finish()

		var buf []byte
		var expected *model.CoinbaseMarketDataCurrency

		buf = []byte(`{}`)
		expected = &model.CoinbaseMarketDataCurrency{}
		test("no values in the response body", g, clientCtrl, buf, "btc", expected)

		buf = []byte(`{"id":"BTC"}`)
		expected = &model.CoinbaseMarketDataCurrency{ID: "BTC"}

		test("value in response body", g, clientCtrl, buf, "btc", expected)
	})
}
