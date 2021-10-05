package coinbase

import (
	"cql/client"
	"cql/model"
)

type CoinbaseAccounts struct{}

func NewCoinbaseAccounts() *CoinbaseAccounts { return &CoinbaseAccounts{} }

func (accounts *CoinbaseAccounts) generateCryptoAddress(conn client.Connector, id string) (m *model.CoinbaseDepositAddress, err error) {
	return m, conn.Decode(&client.Request{
		Method:   client.POST,
		Endpoint: CoinbaseAddressesEP,
		EndpointArgs: client.EndpointArgs{
			"id": &client.EndpointArg{PathParam: &id}},
		// Body: client.JSONBody(map[string]null.Interface{}),
	}, &m)
}

func (account *CoinbaseAccounts) wallets(conn client.Connector) (m []*model.CoinbaseWallet, err error) {
	return m, conn.Decode(&client.Request{Method: client.GET,
		Endpoint: CoinbaseAccountsEP}, &m)
}

// GenerateCryptoAddress will generates a one-time crypto address for depositing
// crypto.
//
// You can generate an address for crypto deposits. See the Coinbase Accounts
// section for information on how to retrieve your coinbase account ID:
// https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_getcoinbaseaccounts
//
// This endpoint requires the "transfer" permission. API key must belong to
// default profile.
//
// source: https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_postcoinbaseaccountaddresses
func (accounts *CoinbaseAccounts) GenerateCryptoAddress(id string) (*model.CoinbaseDepositAddress, error) {
	return accounts.generateCryptoAddress(newClient, id)
}

// All lists all the user's available Coinbase wallets (These are the
// wallets/accounts that are used for buying and selling on www.coinbase.com)
//
// source: https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_getcoinbaseaccounts
func (accounts *CoinbaseAccounts) Wallets() ([]*model.CoinbaseWallet, error) {
	return accounts.wallets(newClient)
}
