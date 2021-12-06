package model

import "github.com/cryptometrics/cql/protomodel"

// * This file was initialized by schema/generate.py, but is open to extension

// CoinbaseWallet represents a user's available Coinbase wallet (These are the
// wallets/accounts that are used for buying and selling on www.coinbase.com)
type CoinbaseWallet struct{ protomodel.CoinbaseWallet }
