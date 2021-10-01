package coinbase

import (
	"cql/client"
	"fmt"
)

type Endpoint int

const (
	_ Endpoint = iota
	AccountHoldsEP
	AccountEP
	AccountsEP
	AccountHistoryEP
	CreateOrderEP
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

// func setQueryParam(idx int, param string, values *url.Values, args ...*string) {
// 	if len(args) > idx {
// 		if v := args[idx]; v != nil {
// 			values.Add(param, *v)
// 		}
// 	}
// }

// func queryParams(params []string, args client.EndpointArgs, argsToExclude ...string) *url.URL {
// 	u, _ := url.Parse("")
// 	q, _ := url.ParseQuery(u.RawQuery)
// 	for idx, param := range params {
// 		setQueryParam(idx, param, &q, args...)
// 	}
// 	u.RawQuery = q.Encode()
// 	return u
// }

func (endpoint Endpoint) Get(args client.EndpointArgs) string {
	return map[Endpoint]func(args client.EndpointArgs) string{

		// List trading accounts from the profile of the API key
		AccountsEP: func(_ client.EndpointArgs) string { return "/accounts" },

		// get trading accouns from the profile of the API key using account id
		AccountEP: func(args client.EndpointArgs) string {
			return fmt.Sprintf("%s/%s", AccountsEP.Get(nil), *args["id"].PathParam)
		},

		AccountHistoryEP: func(args client.EndpointArgs) string {
			return fmt.Sprintf("%s/ledger%s", AccountEP.Get(args), args.QueryPath())
		},

		AccountHoldsEP: func(args client.EndpointArgs) string {
			return fmt.Sprintf("%s/holds%s", AccountEP.Get(args), args.QueryPath())
		},

		CreateOrderEP: func(args client.EndpointArgs) string {
			return fmt.Sprintf("/orders%s", args.QueryPath())

		},

		// List known currencies.
		CurrenciesEP: func(_ client.EndpointArgs) string { return "/currencies" },

		// List the currency for specified id.
		CurrencyEP: func(args client.EndpointArgs) string {
			return fmt.Sprintf("%s/%s", CurrenciesEP.Get(nil), *args["id"].PathParam)
		},

		// list the available currency pairs for trading
		ProductsEP: func(_ client.EndpointArgs) string { return "/products" },

		// market data for a specific currency pair
		ProductEP: func(args client.EndpointArgs) string {
			return fmt.Sprintf("%s/%s", ProductsEP.Get(nil), *args["id"].PathParam)
		},

		// get 24 hr stats for the product. volume is in base currency units. open,
		// high, low are in quote currency units.
		ProductDailyStatsEP: func(args client.EndpointArgs) string {
			return fmt.Sprintf("%s/stats", ProductEP.Get(args))
		},

		// List of historic rates for a product. Rates are returned in grouped
		// buckets based on requested granularity.
		ProductHistoricalRatesEP: func(args client.EndpointArgs) string {
			return fmt.Sprintf("%s/candles%s", ProductEP.Get(args), args.QueryPath())
		},

		// List of open orders for a product
		ProductOrderBookEP: func(args client.EndpointArgs) string {
			return fmt.Sprintf("%s/book%s", ProductEP.Get(args), args.QueryPath())
		},

		// Snapshot information about the last trade (tick), best bid/ask and 24h
		// volume
		ProductTickerEP: func(args client.EndpointArgs) string {
			return fmt.Sprintf("%s/ticker", ProductEP.Get(args))
		},

		// List the latest trades for a product
		ProductTradesEP: func(args client.EndpointArgs) string {
			return fmt.Sprintf("%s/trades", ProductEP.Get(args))
		},

		TimeEP: func(_ client.EndpointArgs) string { return "/time" },
	}[endpoint](args)
}
