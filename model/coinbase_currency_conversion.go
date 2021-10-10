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
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}

	data.UnmarshalString("id", &conversion.ID)
	data.UnmarshalString("from_account_id", &conversion.FromAccountID)
	data.UnmarshalString("to_account_id", &conversion.ToAccountID)
	data.UnmarshalString("from", &conversion.From)
	data.UnmarshalString("to", &conversion.To)

	err = data.UnmarshalFloatFromString("amount", &conversion.Amount)
	if err != nil {
		return err
	}

	return nil
}
