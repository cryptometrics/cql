package model

import "cryptometrics/cql/protomodel"

// * This file was initialized by schema/generate.py, but is open to extension

// CoinbaseDeposit is the response for deposited funds from a www.coinbase.com
// wallet to the specified profile_id.
type CoinbaseDeposit struct{ protomodel.CoinbaseDeposit }
