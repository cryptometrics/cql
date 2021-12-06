package model

import "github.com/cryptometrics/cql/protomodel"

// * This file was initialized by schema/generate.py, but is open to extension

// CoinbaseWebsocketTicker is real-time price updates every time a match
// happens.  It batches updates in case of cascading matches, greatly reducing
// bandwidth requirements.
type CoinbaseWebsocketTicker struct {
	protomodel.CoinbaseWebsocketTicker
}
