package protomodel

// * This is a generated file, do not edit

// CoinbaseWithdrawal is data concerning withdrawing funds from the specified// profile_id to a www.coinbase.com wallet.
type CoinbaseWithdrawal struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
	Fee      float64 `json:"fee"`
	Id       string  `json:"id"`
	PayoutAt string  `json:"payout_at"`
	Subtotal float64 `json:"subtotal"`
}
