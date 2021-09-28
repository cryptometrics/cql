package model

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
	data, err := newUmap(d)
	if err != nil {
		return err
	}

	data.unmarshalString("id", &account.ID)
	data.unmarshalString("currency", &account.Currency)
	data.unmarshalString("profile_id", &account.ProfileID)
	data.unmarshalBool("trading_enabled", &account.TradingEnabled)

	err = data.unmarshalFloatFromString("balance", &account.Balance)
	if err != nil {
		return err
	}

	err = data.unmarshalFloatFromString("available", &account.Available)
	if err != nil {
		return err
	}
	err = data.unmarshalFloatFromString("hold", &account.Hold)
	if err != nil {
		return err
	}

	return nil
}
