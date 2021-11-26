package model

import "cql/protomodel"

// * This file was initialized by schema/generate.py, but is open to extension

// CoinbaseAccountLedger lists ledger activity for an account. This includes
// anything that would affect the accounts balance - transfers, trades, fees,
// etc.
type CoinbaseAccountLedger struct {
	protomodel.CoinbaseAccountLedger
}
