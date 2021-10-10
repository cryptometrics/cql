package model

import (
	"cql/serial"
)

type CoinbaseDeposit struct {
	ID       string  `json:"id"`
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
	PayoutAt string  `json:"payout_at"`
	Fee      float64 `json:"fee"`
	Subtotal float64 `json:"subtotal"`
}

// UnmarshalJSON is an override required to parst strings from coinbases api
// into floats, specifically min_size and max_precision
func (deposit *CoinbaseDeposit) UnmarshalJSON(d []byte) error {
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}

	data.UnmarshalString("id", &deposit.ID)
	data.UnmarshalString("currency", &deposit.Currency)
	data.UnmarshalString("payout_at", &deposit.PayoutAt)

	err = data.UnmarshalFloatFromString("amount", &deposit.Amount)
	if err != nil {
		return err
	}

	err = data.UnmarshalFloatFromString("fee", &deposit.Fee)
	if err != nil {
		return err
	}

	err = data.UnmarshalFloatFromString("subtotal", &deposit.Subtotal)
	if err != nil {
		return err
	}

	return nil
}
