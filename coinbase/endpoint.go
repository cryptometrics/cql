package coinbase

import (
	"fmt"
	"net/url"
)

type Endpoint int

const (
	_ Endpoint = iota
	AccountEP
	AccountsEP
	AccountHistoryEP
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

func setQueryParam(idx int, param string, values *url.Values, args ...*string) {
	if len(args) > idx {
		if v := args[idx]; v != nil {
			values.Add(param, *v)
		}
	}
}

func queryParams(params []string, args []*string) *url.URL {
	u, _ := url.Parse("")
	q, _ := url.ParseQuery(u.RawQuery)
	for idx, param := range params {
		setQueryParam(idx, param, &q, args...)
	}
	u.RawQuery = q.Encode()
	return u
}

func (endpoint Endpoint) Get(args ...*string) string {
	return map[Endpoint]func(args ...*string) string{

		// List trading accounts from the profile of the API key
		AccountsEP: func(_ ...*string) string { return "/accounts" },

		// get trading accouns from the profile of the API key using account id
		AccountEP: func(args ...*string) string {
			return fmt.Sprintf("%s/%s", AccountsEP.Get(), *args[0])
		},

		AccountHistoryEP: func(args ...*string) string {
			params := []string{"before", "after", "start_date", "end_date", "limit"}
			return fmt.Sprintf("%s/ledger%s", AccountEP.Get(args[0]),
				queryParams(params, args[1:]).String())
		},

		// List known currencies.
		CurrenciesEP: func(_ ...*string) string { return "/currencies" },

		// List the currency for specified id.
		CurrencyEP: func(args ...*string) string {
			return fmt.Sprintf("%s/%s", CurrenciesEP.Get(), *args[0])
		},

		// list the available currency pairs for trading
		ProductsEP: func(_ ...*string) string { return "/products" },

		// market data for a specific currency pair
		ProductEP: func(args ...*string) string {
			return fmt.Sprintf("%s/%s", ProductsEP.Get(), *args[0])
		},

		// get 24 hr stats for the product. volume is in base currency units. open,
		// high, low are in quote currency units.
		ProductDailyStatsEP: func(args ...*string) string {
			return fmt.Sprintf("%s/stats", ProductEP.Get(args[0]))
		},

		// List of historic rates for a product. Rates are returned in grouped
		// buckets based on requested granularity.
		ProductHistoricalRatesEP: func(args ...*string) string {
			return fmt.Sprintf("%s/candles?start=%s&end=%s&granularity=%s",
				ProductEP.Get(args[0]), *args[1], *args[2], *args[3])
		},

		// List of open orders for a product
		ProductOrderBookEP: func(args ...*string) string {
			return fmt.Sprintf("%s/book?level=%s", ProductEP.Get(args[0]), *args[1])
		},

		// Snapshot information about the last trade (tick), best bid/ask and 24h
		// volume
		ProductTickerEP: func(args ...*string) string {
			return fmt.Sprintf("%s/ticker", ProductEP.Get(args[0]))
		},

		// List the latest trades for a product
		ProductTradesEP: func(args ...*string) string {
			return fmt.Sprintf("%s/trades", ProductEP.Get(args[0]))
		},

		TimeEP: func(_ ...*string) string { return "/time" },
	}[endpoint](args...)
}
