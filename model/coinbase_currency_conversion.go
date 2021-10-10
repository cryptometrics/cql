package model

import "cql/serial"

// CoinbaseCurrencyConversion corresponds with a successful conversion, which will be assigned a
// conversion id. The corresponding ledger entries for a conversion will reference this conversion
// id.
type CoinbaseCurrencyConversion struct {
	ID            string  `json:"id"`
	Amount        float64 `json:"amount"`
	FromAccountID string  `json:"from_account_id"`
	ToAccountID   string  `json:"to_account_id"`
	From          string  `json:"from"`
	To            string  `json:"to"`
}

// UnmarshalJSON is an override required to parst strings from coinbases api
// into floats, specifically min_size and max_precision
func (conversion *CoinbaseCurrencyConversion) UnmarshalJSON(d []byte) error {
	_, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}

	// data.UnmarshalString("id", &account.ID)
	// data.UnmarshalString("currency", &account.Currency)
	// data.UnmarshalString("profile_id", &account.ProfileID)
	// data.UnmarshalBool("trading_enabled", &account.TradingEnabled)

	// err = data.UnmarshalFloatFromString("balance", &account.Balance)
	// if err != nil {
	// 	return err
	// }

	// err = data.UnmarshalFloatFromString("available", &account.Available)
	// if err != nil {
	// 	return err
	// }
	// err = data.UnmarshalFloatFromString("hold", &account.Hold)
	// if err != nil {
	// 	return err
	// }

	return nil
}
