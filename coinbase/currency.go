package coinbase

import (
	"cql/client2"
	"cql/model"
)

// Conversions is an object used to query on currency data
type Currency struct {
	parent
}

// NewCurrency will return a new conversion obejct to retrieve currency data
func NewCurrency(conn client2.Connector) *Currency {
	currency := new(Currency)
	construct(&currency.parent, conn)
	return currency
}

// All gets a list of all known currencies
// Note: Not all currencies may be currently in use for trading
func (currency *Currency) All() (m []*model.CoinbaseCurrency, err error) {
	return m, currency.get(ENDPOINT_CURRENCIES).Fetch().Assign(&m).JoinMessages()
}

// Find gets a single currency by id.
//
// Currency codes will conform to the ISO 4217 standard where possible.
// Currencies which have or had no representation in ISO 4217 may use a custom
// code.
func (currency *Currency) Find(id string) (m *model.CoinbaseCurrency, err error) {
	req := currency.get(ENDPOINT_CURRENCY)
	return m, req.PathParam("id", id).Fetch().Assign(&m).JoinMessages()

}
