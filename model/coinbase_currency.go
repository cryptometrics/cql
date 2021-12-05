package model

import "cql/protomodel"

// * This file was initialized by schema/generate.py, but is open to extension

// CoinbaseCurrency is a currency that coinbase knows about.  Not al currencies
// may be currently in use for trading.
type CoinbaseCurrency struct{ protomodel.CoinbaseCurrency }
