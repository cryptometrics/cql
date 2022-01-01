package protomodel

import "github.com/cryptometrics/cql/serial"

// * This is a generated file, do not edit

// CoinbaseDeposit is the response for deposited funds from a www.coinbase.com
// wallet to the specified profile_id.
type CoinbaseDeposit struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
	Fee      float64 `json:"fee"`
	Id       string  `json:"id"`
	PayoutAt string  `json:"payout_at"`
	Subtotal float64 `json:"subtotal"`
}

// UnmarshalJSON will deserialize bytes into a CoinbaseDeposit model
func (coinbaseDeposit *CoinbaseDeposit) UnmarshalJSON(d []byte) error {
	const (
		idJsonTag       = "id"
		amountJsonTag   = "amount"
		currencyJsonTag = "currency"
		payoutAtJsonTag = "payout_at"
		feeJsonTag      = "fee"
		subtotalJsonTag = "subtotal"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	data.UnmarshalFloatFromString(amountJsonTag, &coinbaseDeposit.Amount)
	data.UnmarshalFloatFromString(feeJsonTag, &coinbaseDeposit.Fee)
	data.UnmarshalFloatFromString(subtotalJsonTag, &coinbaseDeposit.Subtotal)
	data.UnmarshalString(currencyJsonTag, &coinbaseDeposit.Currency)
	data.UnmarshalString(idJsonTag, &coinbaseDeposit.Id)
	data.UnmarshalString(payoutAtJsonTag, &coinbaseDeposit.PayoutAt)
	return nil
}
