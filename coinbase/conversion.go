package coinbase

import "cql/client"

// Conversions is an object used to converts funds from `from` currency to `to`
// currency.
type Conversions struct {
	parent
}

// NewConversion will return a new conversion obejct to convert funds
func NewConversion(conn client.Connector) *CoinbaseAccounts {
	coinbaseAccounts := new(CoinbaseAccounts)
	construct(&coinbaseAccounts.parent, conn)
	return coinbaseAccounts
}
