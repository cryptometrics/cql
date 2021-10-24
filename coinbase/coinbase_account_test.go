package coinbase

import (
	"cql/client"
	"cql/client2"
	"cql/model"
	"fmt"
	"io/ioutil"
	"testing"

	. "github.com/franela/goblin"
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
)

func TestCoinbaseAccountsWallets(t *testing.T) {
	logrus.SetOutput(ioutil.Discard)

	g := Goblin(t)
	g.Describe("CoinbaseAccounts#Wallet", func() {
		clientCtrl := gomock.NewController(t)
		defer clientCtrl.Finish()

		mocker1 := &mockC{
			desc:     "should correctly assign slice",
			g:        g,
			ctrl:     clientCtrl,
			endpoint: ENDPOINT_COINBASE_ACCOUNTS,
			data:     []byte(`[{"id":"test"}]`),
			assertReq: func(g *G, r client2.Request) {
				g.Assert(r.Method).Eql(client.GET)
				g.Assert(r.EndpointArgs()).IsZero()
			},
			assert: func(g *G, c client2.Connector) {
				ca := NewCoinbaseAccounts(c)
				m, err := ca.Wallets()
				if err != nil {
					panic(err)
				}
				g.Assert(err).IsNil()

				expected := []*model.CoinbaseWallet{{ID: "test"}}
				g.Assert(m).Eql(expected)
			},
		}
		mocker1.run()
	})
}

func TestCoinbaseAccountsGenerateCryptoAddress(t *testing.T) {
	logrus.SetOutput(ioutil.Discard)

	g := Goblin(t)
	g.Describe("CoinbaseAccounts#GenerateCryptoAddress", func() {
		clientCtrl := gomock.NewController(t)
		defer clientCtrl.Finish()

		id1 := "78e6166a-717c-5beb-b095-043601d66f30"
		depositID := "test"
		mocker1 := &mockC{
			desc:     "should correctly assign object",
			g:        g,
			ctrl:     clientCtrl,
			endpoint: ENDPOINT_COINBASE_ACCOUNT_ADDRESSES,
			data:     []byte(fmt.Sprintf(`{"id":"%s"}`, depositID)),
			assertReq: func(g *G, r client2.Request) {
				g.Assert(r.Method).Eql(client.POST)
				g.Assert(*r.EndpointArgs()["id"].PathParam).Eql(id1)
			},
			assert: func(g *G, c client2.Connector) {
				ca := NewCoinbaseAccounts(c)
				m, err := ca.GenerateCryptoAddress(id1)
				if err != nil {
					panic(err)
				}
				g.Assert(err).IsNil()

				expected := &model.CoinbaseDepositAddress{ID: depositID}
				g.Assert(m).Eql(expected)
			},
		}
		mocker1.run()
	})
}
