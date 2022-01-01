package model

import "github.com/cryptometrics/cql/protomodel"

// * This file was initialized by the meta-program, but is open to modification

// CoinbaseAccountLedger lists ledger activity for an account. This includes
// anything that would affect the accounts balance - transfers, trades, fees,
// etc.
type CoinbaseAccountLedger struct {
	protomodel.CoinbaseAccountLedger
}
