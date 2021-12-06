package coinbase

import (
	"github.com/cryptometrics/cql/client"

	"github.com/cryptometrics/cql/model"
)

// CoinbaseAccounts is an object used to query coinbase account data.
type CoinbaseAccounts struct {
	client.Parent
}

// NewCoinbaseAccounts will return an object to query coinbase account data.
// This is not the same as trading accounts, rather account associated with
// coinbase wallets.
func NewCoinbaseAccounts(conn client.Connector) *CoinbaseAccounts {
	coinbaseAccounts := new(CoinbaseAccounts)
	client.ConstructParent(&coinbaseAccounts.Parent, conn)
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
func (accounts *CoinbaseAccounts) GenerateCryptoAddress(walletId string) (m *model.CoinbaseCryptoAddress, err error) {
	req := accounts.Post(AddressesEndpoint)
	return m, req.PathParam("account_id", walletId).Fetch().Assign(&m).JoinMessages()
}

// All lists all the user's available Coinbase wallets (These are the
// wallets/accounts that are used for buying and selling on www.coinbase.com)
//
// * source: https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_getcoinbaseaccounts
func (accounts *CoinbaseAccounts) Wallets() (m []*model.CoinbaseWallet, err error) {
	req := accounts.Get(WalletsEndpoint)
	return m, req.Fetch().Assign(&m).JoinMessages()
}
