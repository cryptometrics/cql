package coinbase

import (
	"cql/client"
	"cql/model"
)

type CoinbaseAccounts struct{}

func NewCoinbaseAccounts() *CoinbaseAccounts { return &CoinbaseAccounts{} }

func (account *CoinbaseAccounts) wallets(conn client.Connector) (m []*model.CoinbaseWallet, err error) {
	return m, conn.Decode(&client.Request{Method: client.GET,
		Endpoint: CoinbaseAccountsEP}, &m)
}

// All lists all the user's available Coinbase wallets (These are the
// wallets/accounts that are used for buying and selling on www.coinbase.com)
func (accounts *CoinbaseAccounts) Wallets() ([]*model.CoinbaseWallet, error) {
	return accounts.wallets(newClient)
}
