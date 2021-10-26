package model

import (
	"cql/serial"
)

type CoinbaseAccountWithdrawal struct {
	ProfileID         string  `json:"profile_id"`
	Amount            float64 `json:"amount"`
	CoinbaseAccountID string  `json:"coinbase_account_id"`
	Currency          string  `json:"currency"`
}

// UnmarshalJSON is an override required to parst strings from coinbases api
// into floats, specifically min_size and max_precision
func (withdrawal *CoinbaseAccountWithdrawal) UnmarshalJSON(d []byte) error {
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}

	data.UnmarshalString("profile_id", &withdrawal.ProfileID)
	data.UnmarshalString("coinbase_account_id", &withdrawal.CoinbaseAccountID)
	data.UnmarshalString("currency", &withdrawal.Currency)
	data.UnmarshalFloatFromString("amount", &withdrawal.Amount)

	return nil
}
