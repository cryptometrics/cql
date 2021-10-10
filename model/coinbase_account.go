package model

import "cql/serial"

// CoinbaseAccount encapsulates information for a coinbase account
type CoinbaseAccount struct {
	ID             string  `json:"id"`
	Currency       string  `json:"currency"`
	Balance        float64 `json:"balance"`
	Available      float64 `json:"available"`
	Hold           float64 `json:"hold"`
	ProfileID      string  `json:"profile_id"`
	TradingEnabled bool    `json:"trading_enabled"`
}

// UnmarshalJSON is an override required to parst strings from coinbases api
// into floats, specifically min_size and max_precision
func (account *CoinbaseAccount) UnmarshalJSON(d []byte) error {
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}

	data.UnmarshalString("id", &account.ID)
	data.UnmarshalString("currency", &account.Currency)
	data.UnmarshalString("profile_id", &account.ProfileID)
	data.UnmarshalBool("trading_enabled", &account.TradingEnabled)

	err = data.UnmarshalFloatFromString("balance", &account.Balance)
	if err != nil {
		return err
	}

	err = data.UnmarshalFloatFromString("available", &account.Available)
	if err != nil {
		return err
	}
	err = data.UnmarshalFloatFromString("hold", &account.Hold)
	if err != nil {
		return err
	}

	return nil
}
