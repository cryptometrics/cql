package coinbase

import "fmt"

type Endpoint int

const (
	_ Endpoint = iota
	CurrenciesEP
	CurrencyEP
	ProductsEP
	ProductEP
	ProductOrderBookEP
	ProductTickerEP
)

func (endpoint Endpoint) Get(args ...string) string {
	return map[Endpoint]func(args ...string) string{
		// List known currencies.
		CurrenciesEP: func(_ ...string) string { return "/currencies" },

		// List the currency for specified id.
		CurrencyEP: func(args ...string) string {
			return fmt.Sprintf("%s/%s", CurrenciesEP.Get(), args[0])
		},

		// list the available currency pairs for trading
		ProductsEP: func(_ ...string) string { return "/products" },

		// market data for a specific currency pair
		ProductEP: func(args ...string) string {
			return fmt.Sprintf("%s/%s", ProductsEP.Get(), args[0])
		},

		// List of open orders for a product
		ProductOrderBookEP: func(args ...string) string {
			return fmt.Sprintf("%s/book?level=%s", ProductEP.Get(args[0]), args[1])
		},

		// Snapshot information about the last trade (tick), best bid/ask and 24h
		// volume
		ProductTickerEP: func(args ...string) string {
			return fmt.Sprintf("%s/ticker", ProductEP.Get(args[0]))
		},
	}[endpoint](args...)
}
