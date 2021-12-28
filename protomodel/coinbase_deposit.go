package protomodel

// * This is a generated file, do not edit

// CoinbaseDeposit is the response for deposited funds from a www.coinbase.com// wallet to the specified profile_id.
type CoinbaseDeposit struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
	Fee      float64 `json:"fee"`
	Id       string  `json:"id"`
	PayoutAt string  `json:"payout_at"`
	Subtotal float64 `json:"subtotal"`
}
