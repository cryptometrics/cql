package coinbase

// import (
// 	. "github.com/franela/goblin"
// )

// TODO -- get rid of these tests, just need to test decode.  How they actually
// TODO -- get decoded is in the model by unmarshal methods.
// ! Think this through though!

// func TestMarketDataCurrencies(t *testing.T) {
// 	g := Goblin(t)
// 	g.Describe("Test MarketData#Currencies", func() {
// 		test := func(desc string, g *G, ctrl *gomock.Controller, data []byte,
// 			expected []*model.CoinbaseCurrency) {
// 			g.It(desc, func() {
// 				mockC := client.NewMockC(ctrl)

// 				mockC.EXPECT().Connect().Return(nil)

// 				mockC.EXPECT().Get(gomock.Any()).DoAndReturn(
// 					func(endpoint string) (*http.Response, error) {
// 						g.Assert(endpoint).Equal(CurrenciesEP.Get())

// 						stringReader := bytes.NewReader(data)
// 						stringReadCloser := ioutil.NopCloser(stringReader)
// 						return &http.Response{Body: stringReadCloser}, nil
// 					})

// 				var gen = func(kind client.Kind) (client.C, error) {
// 					return mockC, nil
// 				}

// 				md := NewMarketData()
// 				m, err := md.currencies(gen)
// 				if err != nil {
// 					panic(err)
// 				}
// 				g.Assert(err).IsNil()
// 				g.Assert(m).Eql(expected)
// 			})
// 		}

// 		clientCtrl := gomock.NewController(t)
// 		defer clientCtrl.Finish()

// 		var buf []byte
// 		var expected []*model.CoinbaseCurrency

// 		buf = []byte(`[]`)
// 		expected = []*model.CoinbaseCurrency{}
// 		test("no values in the response body", g, clientCtrl, buf, expected)

// 		buf = []byte(`[{"id":"BTC"}]`)
// 		expected = []*model.CoinbaseCurrency{
// 			{ID: "BTC"},
// 		}
// 		test("one value in the response body", g, clientCtrl, buf, expected)

// 		buf = []byte(`[{"id":"BTC"},{"id":"ETH"}]`)
// 		expected = []*model.CoinbaseCurrency{
// 			{ID: "BTC"}, {ID: "ETH"},
// 		}
// 		test("two values in the response body", g, clientCtrl, buf, expected)

// 		buf = []byte(`[{"details":{"type":"crypto"}}]`)
// 		expected = []*model.CoinbaseCurrency{
// 			{Details: model.CoinbaseCurrencyDetails{Type: "crypto"}},
// 		}
// 		test("details should parse correctly", g, clientCtrl, buf, expected)

// 		buf = []byte(`[{"min_size":"0.01"}]`)
// 		expected = []*model.CoinbaseCurrency{
// 			{MinSize: 0.01},
// 		}
// 		test("strings-floats should convert to floats", g, clientCtrl, buf, expected)
// 	})
// }

// func TestMarketDataCurrency(t *testing.T) {
// 	g := Goblin(t)
// 	g.Describe("Test MarketData#Currency", func() {
// 		test := func(desc string, g *G, ctrl *gomock.Controller, data []byte, id string,
// 			expected *model.CoinbaseCurrency) {
// 			g.It(desc, func() {
// 				mockC := client.NewMockC(ctrl)

// 				mockC.EXPECT().Connect().Return(nil)

// 				mockC.EXPECT().Get(gomock.Any()).DoAndReturn(
// 					func(endpoint string) (*http.Response, error) {
// 						g.Assert(endpoint).Equal(CurrencyEP.Get(id))

// 						stringReader := bytes.NewReader(data)
// 						stringReadCloser := ioutil.NopCloser(stringReader)
// 						return &http.Response{Body: stringReadCloser}, nil
// 					})

// 				var gen = func(kind client.Kind) (client.C, error) {
// 					return mockC, nil
// 				}

// 				md := NewMarketData()
// 				m, err := md.currency(gen, id)
// 				if err != nil {
// 					panic(err)
// 				}
// 				g.Assert(err).IsNil()
// 				g.Assert(m).Eql(expected)
// 			})
// 		}

// 		clientCtrl := gomock.NewController(t)
// 		defer clientCtrl.Finish()

// 		var buf []byte
// 		var expected *model.CoinbaseCurrency

// 		buf = []byte(`{}`)
// 		expected = &model.CoinbaseCurrency{}
// 		test("no values in the response body", g, clientCtrl, buf, "btc", expected)

// 		buf = []byte(`{"id":"BTC"}`)
// 		expected = &model.CoinbaseCurrency{ID: "BTC"}

// 		test("value in response body", g, clientCtrl, buf, "btc", expected)
// 	})
// }

// func TestMarketDataProducts(t *testing.T) {
// 	g := Goblin(t)
// 	g.Describe("Test MarketData#Products", func() {
// 		test := func(desc string, g *G, ctrl *gomock.Controller, data []byte,
// 			expected []*model.CoinbaseProduct) {
// 			g.It(desc, func() {
// 				mockC := client.NewMockC(ctrl)

// 				mockC.EXPECT().Connect().Return(nil)

// 				mockC.EXPECT().Get(gomock.Any()).DoAndReturn(
// 					func(endpoint string) (*http.Response, error) {
// 						g.Assert(endpoint).Equal(ProductsEP.Get())

// 						stringReader := bytes.NewReader(data)
// 						stringReadCloser := ioutil.NopCloser(stringReader)
// 						return &http.Response{Body: stringReadCloser}, nil
// 					})

// 				var gen = func(kind client.Kind) (client.C, error) {
// 					return mockC, nil
// 				}

// 				md := NewMarketData()
// 				m, err := md.products(gen)
// 				if err != nil {
// 					panic(err)
// 				}
// 				g.Assert(err).IsNil()
// 				g.Assert(m).Eql(expected)
// 			})
// 		}

// 		clientCtrl := gomock.NewController(t)
// 		defer clientCtrl.Finish()

// 		var buf []byte
// 		var expected []*model.CoinbaseProduct

// 		buf = []byte(`[]`)
// 		expected = []*model.CoinbaseProduct{}
// 		test("no values in the response body", g, clientCtrl, buf, expected)

// 		buf = []byte(`[{"id":"BTC"}]`)
// 		expected = []*model.CoinbaseProduct{
// 			{ID: "BTC"},
// 		}
// 		test("one value in the response body", g, clientCtrl, buf, expected)

// 		buf = []byte(`[{"id":"BTC"},{"id":"ETH"}]`)
// 		expected = []*model.CoinbaseProduct{
// 			{ID: "BTC"}, {ID: "ETH"},
// 		}
// 		test("two values in the response body", g, clientCtrl, buf, expected)

// 		buf = []byte(`[{"quote_increment":"0.01"}]`)
// 		expected = []*model.CoinbaseProduct{
// 			{QuoteIncrement: 0.01},
// 		}
// 		test("strings-floats should convert to floats", g, clientCtrl, buf, expected)
// 	})
// }

// func TestMarketDataProduct(t *testing.T) {
// 	g := Goblin(t)
// 	g.Describe("Test MarketData#Product", func() {
// 		test := func(desc string, g *G, ctrl *gomock.Controller, data []byte, id string,
// 			expected *model.CoinbaseProduct) {
// 			g.It(desc, func() {
// 				mockC := client.NewMockC(ctrl)

// 				mockC.EXPECT().Connect().Return(nil)

// 				mockC.EXPECT().Get(gomock.Any()).DoAndReturn(
// 					func(endpoint string) (*http.Response, error) {
// 						g.Assert(endpoint).Equal(ProductEP.Get(id))

// 						stringReader := bytes.NewReader(data)
// 						stringReadCloser := ioutil.NopCloser(stringReader)
// 						return &http.Response{Body: stringReadCloser}, nil
// 					})

// 				var gen = func(kind client.Kind) (client.C, error) {
// 					return mockC, nil
// 				}

// 				md := NewMarketData()
// 				m, err := md.product(gen, id)
// 				if err != nil {
// 					panic(err)
// 				}
// 				g.Assert(err).IsNil()
// 				g.Assert(m).Eql(expected)
// 			})
// 		}

// 		clientCtrl := gomock.NewController(t)
// 		defer clientCtrl.Finish()

// 		var buf []byte
// 		var expected *model.CoinbaseProduct

// 		buf = []byte(`{}`)
// 		expected = &model.CoinbaseProduct{}
// 		test("no values in the response body", g, clientCtrl, buf, "btc", expected)

// 		buf = []byte(`{"id":"BTC"}`)
// 		expected = &model.CoinbaseProduct{ID: "BTC"}

// 		test("value in response body", g, clientCtrl, buf, "btc", expected)
// 	})
// }
