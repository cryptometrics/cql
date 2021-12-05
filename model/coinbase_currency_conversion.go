package model

import "cql/protomodel"

// * This file was initialized by schema/generate.py, but is open to extension

// CoinbaseCurrencyConversion is the response that converts funds from from
// currency to to currency. Funds are converted on the from account in the
// profile_id profile.
type CoinbaseCurrencyConversion struct {
	protomodel.CoinbaseCurrencyConversion
}
