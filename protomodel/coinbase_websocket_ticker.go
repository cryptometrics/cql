package protomodel

import "time"

// * This is a generated file, do not edit

// CoinbaseWebsocketTicker is real-time price updates every time a match// happens. It batches updates in case of cascading matches, greatly reducing// bandwidth requirements.
type CoinbaseWebsocketTicker struct {
	BestAsk   float64   `json:"best_ask"`
	BestBid   float64   `json:"best_bid"`
	LastSize  float64   `json:"last_size"`
	Price     float64   `json:"price"`
	ProductId string    `json:"product_id"`
	Sequence  int       `json:"sequence"`
	Side      string    `json:"side"`
	Time      time.Time `json:"time"`
	TradeId   int       `json:"trade_id"`
	Type      string    `json:"type"`
}
