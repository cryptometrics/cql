package model

// CoinbaseBuyPrice encapsulates the total price to buy one bitcoin or ether.
type CoinbaseBuyPrice struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}
