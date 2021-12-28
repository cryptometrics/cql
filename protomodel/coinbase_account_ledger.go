package protomodel

import (
	"time"

	"github.com/cryptometrics/cql/scalar"
)

// * This is a generated file, do not edit

// CoinbaseAccountLedger lists ledger activity for an account. This includes// anything that would affect the accounts balance - transfers, trades, fees,// etc.
type CoinbaseAccountLedger struct {
	Amount       float64                      `json:"amount"`
	Balance      float64                      `json:"balance"`
	CreatedAt    time.Time                    `json:"created_at"`
	Id           string                       `json:"id"`
	ProtoDetails CoinbaseAccountLedgerDetails `json:"details"`
	Type         scalar.EntryType             `json:"type"`
}
