package coinbase

import "fmt"

type Endpoint int

const (
	_ Endpoint = iota
	CurrenciesEP
	CurrencyEP
	ProductsEP
	ProductEP
	ProductDailyStatsEP
	ProductHistoricalRatesEP
	ProductOrderBookEP
	ProductTickerEP
	ProductTradesEP
	TimeEP
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

		// get 24 hr stats for the product. volume is in base currency units. open,
		// high, low are in quote currency units.
		ProductDailyStatsEP: func(args ...string) string {
			return fmt.Sprintf("%s/stats", ProductEP.Get(args[0]))
		},

		// List of historic rates for a product. Rates are returned in grouped
		// buckets based on requested granularity.
		ProductHistoricalRatesEP: func(args ...string) string {
			return fmt.Sprintf("%s/candles?start=%s&end=%s&granularity=%s",
				ProductEP.Get(args[0]), args[1], args[2], args[3])
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

		// List the latest trades for a product
		ProductTradesEP: func(args ...string) string {
			return fmt.Sprintf("%s/trades", ProductEP.Get(args[0]))
		},

		TimeEP: func(_ ...string) string { return "/time" },
	}[endpoint](args...)
}
