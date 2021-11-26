package coinbase

// import (
// 	. "github.com/franela/goblin"
// )

// func TestCoinbaseCurrencyFind(t *testing.T) {
// 	logrus.SetOutput(ioutil.Discard)

// 	g := Goblin(t)
// 	g.Describe("CoinbaseCurrency#Find", func() {
// 		clientCtrl := gomock.NewController(t)
// 		defer clientCtrl.Finish()

// 		id1 := "id"
// 		mocker1 := &mockC{
// 			desc:     "should correctly assign object",
// 			g:        g,
// 			ctrl:     clientCtrl,
// 			endpoint: ENDPOINT_CURRENCY,
// 			data: []byte(`{
// 				"id": "USD",
// 				"name": "United States Dollar",
// 				"min_size": "0.01",
// 				"max_precision": "0.01",
// 				"status": "online",
// 				"details": {
// 					"type": "fiat",
// 					"symbol": "$",
// 					"sort_order": 1,
// 					"push_payment_methods": [
// 						"bank_wire",
// 						"fedwire",
// 						"swift_bank_account",
// 						"intra_bank_account"
// 					],
// 					"display_name": "US Dollar",
// 					"group_types": [
// 						"fiat",
// 						"usd"
// 					]
// 				}
// 			}`),
// 			assertReq: func(g *G, r client.Request) {
// 				g.Assert(r.Method).Eql(client.GET)
// 				g.Assert(r.Endpoint).Eql(ENDPOINT_CURRENCY)
// 				g.Assert(*r.EndpointArgs()["id"].PathParam).Eql(id1)
// 			},
// 			assert: func(g *G, c client.Connector) {
// 				conv := NewCurrency(c)
// 				m, err := conv.Find(id1)
// 				if err != nil {
// 					panic(err)
// 				}
// 				g.Assert(err).IsNil()

// 				expected := model.CoinbaseCurrency{
// 					ID:           "USD",
// 					Name:         "United States Dollar",
// 					MinSize:      0.01,
// 					MaxPrecision: 0.01,
// 					Status:       "online",
// 					Details: model.CoinbaseCurrencyDetails{
// 						Type:      "fiat",
// 						Symbol:    "$",
// 						SortOrder: 1,
// 						PushPaymentMethods: []scalar.PaymentMethod{
// 							scalar.PaymentMethodBankWire,
// 							scalar.PaymentMethodFedwire,
// 							scalar.PaymentMethodSWIFTBankAccount,
// 							scalar.PaymentMethodIntraBankAccount,
// 						},
// 						DisplayName: "US Dollar",
// 						GroupTypes: []string{
// 							"fiat",
// 							"usd",
// 						},
// 					},
// 				}
// 				g.Assert(*m).Eql(expected)
// 			},
// 		}
// 		mocker1.run()
// 	})

// 	g.Describe("CoinbaseCurrency#All", func() {
// 		clientCtrl := gomock.NewController(t)
// 		defer clientCtrl.Finish()
// 		mocker1 := &mockC{
// 			desc:     "should correctly assign slice",
// 			g:        g,
// 			ctrl:     clientCtrl,
// 			endpoint: ENDPOINT_CURRENCIES,
// 			data:     []byte(`[{"id":"USD"},{"id":"BTC"}]`),
// 			assertReq: func(g *G, r client.Request) {
// 				g.Assert(r.Method).Eql(client.GET)
// 				g.Assert(r.Endpoint).Eql(ENDPOINT_CURRENCIES)
// 				g.Assert(r.EndpointArgs).IsZero()
// 			},
// 			assert: func(g *G, c client.Connector) {
// 				conv := NewCurrency(c)
// 				m, err := conv.All()
// 				if err != nil {
// 					panic(err)
// 				}
// 				g.Assert(err).IsNil()

// 				g.Assert(*m[0]).Eql(model.CoinbaseCurrency{ID: "USD"})
// 				g.Assert(*m[1]).Eql(model.CoinbaseCurrency{ID: "BTC"})
// 			},
// 		}
// 		mocker1.run()
// 	})
// }
