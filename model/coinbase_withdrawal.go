package model

import "cql/protomodel"

// * This file was initialized by schema/generate.py, but is open to extension

// CoinbaseWithdrawal is data concerning withdrawing funds from the specified
// profile_id to a www.coinbase.com wallet.
type CoinbaseWithdrawal struct{ protomodel.CoinbaseWithdrawal }
