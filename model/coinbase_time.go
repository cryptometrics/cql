package model

// CoinbaseTime encapsulates the coinbase API server time
type CoinbaseTime struct {
	ISO string `json:"iso"`

	// Epoch represents decimal seconds since Unix Epoch
	Epoch float64 `json:"epoch"`
}
