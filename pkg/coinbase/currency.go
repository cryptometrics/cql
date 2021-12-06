package coinbase

import (
	"cryptometrics/cql/client"
	"cryptometrics/cql/model"
)

// CoinbaseAccounts is an object used to query coinbase account data.
type Currency struct {
	client.Parent
}

// NewCurrency will return an object to query currency currencys
func NewCurrency(conn client.Connector) *Currency {
	currency := new(Currency)
	client.ConstructParent(&currency.Parent, conn)
	return currency
}

// All gets a list of all known currencies
// Note: Not all currencies may be currently in use for trading
func (currency *Currency) All() (m []*model.CoinbaseCurrency, err error) {
	return m, currency.Get(CurrenciesEndpoint).Fetch().Assign(&m).JoinMessages()
}

// Find gets a single currency by id.
//
// Currency codes will conform to the ISO 4217 standard where possible.
// Currencies which have or had no representation in ISO 4217 may use a custom
// code.
func (currency *Currency) Find(currencyId string) (m *model.CoinbaseCurrency, err error) {
	req := currency.Get(CurrencyEndpoint)
	return m, req.PathParam("currency_id", currencyId).Fetch().Assign(&m).JoinMessages()
}
