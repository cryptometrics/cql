package coinbase

import (
	"cql/client2"
	"cql/model"
)

// CoinbaseAccounts is an object used to query coinbase account data.
type CoinbaseAccounts struct {
	parent
}

// NewCoinbaseAccounts will return an object to query coinbase account data.
// This is not the same as trading accounts, rather account associated with
// coinbase wallets.
func NewCoinbaseAccounts(conn client2.Connector) *CoinbaseAccounts {
	coinbaseAccounts := new(CoinbaseAccounts)
	construct(&coinbaseAccounts.parent, conn)
	return coinbaseAccounts
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
// * source: https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_postcoinbaseaccountaddresses
func (accounts *CoinbaseAccounts) GenerateCryptoAddress(id string) (m *model.CoinbaseDepositAddress, err error) {
	req := accounts.get(ENDPOINT_COINBASE_ACCOUNT_ADDRESSES)
	return m, req.PathParam("id", id).Fetch().Assign(&m).JoinMessages()
}

// All lists all the user's available Coinbase wallets (These are the
// wallets/accounts that are used for buying and selling on www.coinbase.com)
//
// * source: https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_getcoinbaseaccounts
func (accounts *CoinbaseAccounts) Wallets() (m []*model.CoinbaseWallet, err error) {
	req := accounts.get(ENDPOINT_COINBASE_ACCOUNTS)
	return m, req.Fetch().Assign(&m).JoinMessages()
}
