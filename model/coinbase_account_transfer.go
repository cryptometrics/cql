package model

import "cql/protomodel"

// * This file was initialized by schema/generate.py, but is open to extension

// CoinbaseAccountTransfer will lists past withdrawals and deposits for an
// account.
type CoinbaseAccountTransfer struct {
	protomodel.CoinbaseAccountTransfer
}
