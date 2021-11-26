package coinbase

// import (
// 	"cql/client"
// 	"cql/client"
// 	"cql/model"
// 	"io/ioutil"
// 	"testing"

// 	. "github.com/franela/goblin"
// 	"github.com/golang/mock/gomock"
// 	"github.com/sirupsen/logrus"
// 	"gopkg.in/guregu/null.v3"
// )

// func TestCoinbaseConversionFind(t *testing.T) {
// 	logrus.SetOutput(ioutil.Discard)

// 	g := Goblin(t)
// 	g.Describe("CoinbaseConversion#Find", func() {
// 		clientCtrl := gomock.NewController(t)
// 		defer clientCtrl.Finish()

// 		id1 := "id"
// 		pid1 := "pid"
// 		mocker1 := &mockC{
// 			desc:     "should correctly assign object",
// 			g:        g,
// 			ctrl:     clientCtrl,
// 			endpoint: ENDPOINT_COINBASE_CONVERSION,
// 			data: []byte(`{
// 				"id": "85e5fb8e-13c0-471f-b9a7-367076b443c7",
// 				"to": "",
// 				"from": "",
// 				"to_account_id": "efd6e38b-0449-484d-aa1e-daa0a04ab49a",
// 				"from_account_id": "ccb0db90-3476-4573-9a71-5c88b3f38b6b"
// 			}`),
// 			assertReq: func(g *G, r client.Request) {
// 				g.Assert(r.Method).Eql(client.GET)
// 				g.Assert(r.Endpoint).Eql(ENDPOINT_COINBASE_CONVERSION)
// 				g.Assert(*r.EndpointArgs()["id"].PathParam).Eql(id1)
// 				g.Assert(*r.EndpointArgs()["profile_id"].QueryParam).Eql(pid1)
// 			},
// 			assert: func(g *G, c client.Connector) {
// 				conv := NewConversion(c)
// 				m, err := conv.Find(id1, &model.CoinbaseCurrencyConversionOpts{
// 					ProfileID: null.StringFrom(pid1).Ptr(),
// 				})
// 				if err != nil {
// 					panic(err)
// 				}
// 				g.Assert(err).IsNil()

// 				expected := model.CoinbaseCurrencyConversion{
// 					ID:            "85e5fb8e-13c0-471f-b9a7-367076b443c7",
// 					ToAccountID:   "efd6e38b-0449-484d-aa1e-daa0a04ab49a",
// 					FromAccountID: "ccb0db90-3476-4573-9a71-5c88b3f38b6b",
// 				}
// 				g.Assert(*m).Eql(expected)
// 			},
// 		}
// 		mocker1.run()
// 	})

// 	g.Describe("CoinbaseConversion#Make", func() {
// 		clientCtrl := gomock.NewController(t)
// 		defer clientCtrl.Finish()
// 		var (
// 			to     = "to"
// 			from   = "from"
// 			amount = 10.0
// 			pid    = "pid"
// 			nonce  = "nonce"
// 		)
// 		mocker1 := &mockC{
// 			desc:     "should correctly assign object",
// 			g:        g,
// 			ctrl:     clientCtrl,
// 			endpoint: ENDPOINT_COINBASE_CONVERSIONS,
// 			data: []byte(`{
// 				"id": "85e5fb8e-13c0-471f-b9a7-367076b443c7",
// 				"to": "",
// 				"from": "",
// 				"to_account_id": "efd6e38b-0449-484d-aa1e-daa0a04ab49a",
// 				"from_account_id": "ccb0db90-3476-4573-9a71-5c88b3f38b6b"
// 			}`),
// 			assertReq: func(g *G, r client.Request) {
// 				g.Assert(r.Method).Eql(client.POST)
// 				g.Assert(r.Endpoint).Eql(ENDPOINT_COINBASE_CONVERSIONS)
// 				g.Assert(r.EndpointArgs()).IsZero()
// 			},
// 			assert: func(g *G, c client.Connector) {
// 				conv := NewConversion(c)
// 				m, err := conv.Make(from, to, amount, &model.CoinbaseCurrencyConversionOpts{
// 					ProfileID: null.StringFrom(pid).Ptr(),
// 					Nonce:     null.StringFrom(nonce).Ptr(),
// 				})
// 				if err != nil {
// 					panic(err)
// 				}
// 				g.Assert(err).IsNil()

// 				expected := model.CoinbaseCurrencyConversion{
// 					ID:            "85e5fb8e-13c0-471f-b9a7-367076b443c7",
// 					ToAccountID:   "efd6e38b-0449-484d-aa1e-daa0a04ab49a",
// 					FromAccountID: "ccb0db90-3476-4573-9a71-5c88b3f38b6b",
// 				}
// 				g.Assert(*m).Eql(expected)
// 			},
// 		}
// 		mocker1.run()
// 	})
// }
