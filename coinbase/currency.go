package coinbase

import (
	"cql/client"
	"cql/model"
)

// Conversions is an object used to query on currency data
type Currency struct {
	parent
}

// NewCurrency will return a new conversion obejct to retrieve currency data
func NewCurrency(conn client.Connector) *Currency {
	currency := new(Currency)
	construct(&currency.parent, conn)
	return currency
}

func (currency *Currency) fetchCurrencies() *client.FetchResponse {
	return currency.conn.Fetch(&client.Request{Method: client.GET, Endpoint: CurrenciesEP})
}

func (currency *Currency) fetchCurrency(id string) *client.FetchResponse {
	return currency.conn.Fetch(&client.Request{
		Method:   client.GET,
		Endpoint: CurrencyEP,
		EndpointArgs: client.EndpointArgs{
			"id": &client.EndpointArg{PathParam: &id}},
	})
}

// All gets a list of all known currencies
// Note: Not all currencies may be currently in use for trading
func (currency *Currency) All() (m []*model.CoinbaseCurrency, err error) {
	return m, currency.fetchCurrencies().Assign(&m).Error
}

// Find gets a single currency by id.
//
// Currency codes will conform to the ISO 4217 standard where possible. Currencies which have or had
// no representation in ISO 4217 may use a custom code.
func (currency *Currency) Find(id string) (m *model.CoinbaseCurrency, err error) {
	return m, currency.fetchCurrency(id).Assign(&m).Error
}
