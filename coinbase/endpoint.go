package coinbase

import (
	"cql/client"
	"fmt"
)

type Endpoint int

const (
	_ Endpoint = iota
	ENDPOINT_ACCOUNT_HOLDS
	ENDPOINT_ACCOUNT
	ENDPOINT_ACCOUNTS
	ENDPOINT_ACCOUNT_LEDGER
	ENDPOINT_ACCOUNT_TRANSFERS
	ENDPOINT_COINBASE_ACCOUNTS
	ENDPOINT_COINBASE_ACCOUNT_DEPOSITS
	ENDPOINT_COINBASE_ACCOUNT_ADDRESSES
	ENDPOINT_COINBASE_CONVERSIONS
	ENDPOINT_COINBASE_CONVERSION
	ENDPOINT_CURRENCIES
	ENDPOINT_CURRENCY
	ENDPOINT_TRANSFERS_PAYMENT_METHODS
	ENDPOINT_TRANSFERS
	ENDPOINT_TRANSFER

	DepositsEP

	///////
	ClientOrderEP
	CreateOrderEP
	// CurrenciesEP
	// CurrencyEP
	OrderEP
	ProductsEP
	ProductEP
	ProductDailyStatsEP
	ProductHistoricalRatesEP
	ProductOrderBookEP
	ProductTickerEP
	ProductTradesEP
	TimeEP
)

func (endpoint Endpoint) Get(args client.EndpointArgs) string {
	return map[Endpoint]func(args client.EndpointArgs) string{

		ENDPOINT_ACCOUNTS: func(_ client.EndpointArgs) string { return "/accounts" },

		ENDPOINT_ACCOUNT: func(args client.EndpointArgs) string {
			return fmt.Sprintf("%s/%s", ENDPOINT_ACCOUNTS.Get(nil), *args["id"].PathParam)
		},

		ENDPOINT_ACCOUNT_LEDGER: func(args client.EndpointArgs) string {
			return fmt.Sprintf("%s/ledger%s", ENDPOINT_ACCOUNT.Get(args), args.QueryPath())
		},

		ENDPOINT_ACCOUNT_HOLDS: func(args client.EndpointArgs) string {
			return fmt.Sprintf("%s/holds%s", ENDPOINT_ACCOUNT.Get(args), args.QueryPath())
		},

		ENDPOINT_ACCOUNT_TRANSFERS: func(args client.EndpointArgs) string {
			return fmt.Sprintf("%s/transfers%s", ENDPOINT_ACCOUNT.Get(args), args.QueryPath())
		},

		ENDPOINT_COINBASE_ACCOUNTS: func(_ client.EndpointArgs) string {
			return "/coinbase-accounts"
		},

		ENDPOINT_COINBASE_ACCOUNT_DEPOSITS: func(_ client.EndpointArgs) string {
			return client.JoinEndpointParts(DepositsEP.Get(nil), "coinbase-account")
		},

		ENDPOINT_COINBASE_ACCOUNT_ADDRESSES: func(args client.EndpointArgs) string {
			return fmt.Sprintf("%s/%s/addresses", ENDPOINT_COINBASE_ACCOUNTS.Get(nil), *args["id"].PathParam)
		},

		ENDPOINT_COINBASE_CONVERSIONS: func(_ client.EndpointArgs) string {
			return "/conversions"
		},

		ENDPOINT_COINBASE_CONVERSION: func(_ client.EndpointArgs) string {
			return fmt.Sprintf("%s/%s", ENDPOINT_COINBASE_CONVERSIONS.Get(nil), *args["id"].PathParam)
		},

		ENDPOINT_TRANSFERS_PAYMENT_METHODS: func(_ client.EndpointArgs) string {
			return "/payment-methods"
		},

		ENDPOINT_TRANSFERS: func(_ client.EndpointArgs) string {
			return "/transfers"
		},

		ENDPOINT_TRANSFER: func(_ client.EndpointArgs) string {
			return client.JoinEndpointParts(ENDPOINT_TRANSFERS.Get(nil), *args["id"].PathParam)
		},

		////////

		ClientOrderEP: func(args client.EndpointArgs) string {
			return fmt.Sprintf("/orders/client:%s", *args["client_oid"].PathParam)
		},

		CreateOrderEP: func(args client.EndpointArgs) string {
			return fmt.Sprintf("/orders%s", args.QueryPath())
		},

		// List known currencies.
		ENDPOINT_CURRENCIES: func(_ client.EndpointArgs) string { return "/currencies" },

		// List the currency for specified id.
		ENDPOINT_CURRENCY: func(args client.EndpointArgs) string {
			return fmt.Sprintf("%s/%s", ENDPOINT_CURRENCIES.Get(nil), *args["id"].PathParam)
		},

		OrderEP: func(args client.EndpointArgs) string {
			return fmt.Sprintf("/orders/%s", *args["id"].PathParam)
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
