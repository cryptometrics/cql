package coinbase

import "fmt"

type Endpoint int

const (
	_ Endpoint = iota
	CurrenciesEP
	CurrencyEP
)

func (endpoint Endpoint) Get(args ...string) string {
	return map[Endpoint]func(args ...string) string{
		// List known currencies.
		CurrenciesEP: func(_ ...string) string { return "/currencies" },

		// List the currency for specified id.
		CurrencyEP: func(args ...string) string {
			return fmt.Sprintf("%s/%s", CurrenciesEP.Get(), args[0])
		},
	}[endpoint](args...)
}
