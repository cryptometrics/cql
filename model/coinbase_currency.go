package model

// CoinbaseCurrency holds currency data from the coinbase API.  The
// gqlgen package does not support snake case json, which is why this is
// necessary.
type CoinbaseCurrency struct {
	ID                    string                  `json:"id"`
	Name                  string                  `json:"name"`
	MinSize               float64                 `json:"min_size"`
	Status                string                  `json:"status"`
	Message               string                  `json:"message"`
	MaxPrecision          float64                 `json:"max_precision"`
	ConvertibleTo         []string                `json:"convertible_to"`
	Details               CoinbaseCurrencyDetails `json:"details"`
	DisplayName           string                  `json:"display_name"`
	ProcessingTimeSeconds int                     `json:"processing_time_seconds"`
	MinWithdrawalAmount   int                     `json:"min_withdrawal_amount"`
	MaxWithdrawalAmount   int                     `json:"max_withdrawal_amount"`
}

// UnmarshalJSON is an override required to parst strings from coinbases api
// into floats, specifically min_size and max_precision
func (currency *CoinbaseCurrency) UnmarshalJSON(d []byte) error {
	data, err := newUmap(d)
	if err != nil {
		return err
	}

	data.unmarshalString("id", &currency.ID)
	data.unmarshalString("name", &currency.Name)
	data.unmarshalString("status", &currency.Status)
	data.unmarshalString("message", &currency.Message)
	data.unmarshalStringSlice("convertible_to", &currency.ConvertibleTo)
	data.unmarshalString("display_name", &currency.DisplayName)
	data.unmarshalInt("processing_time_seconds", &currency.ProcessingTimeSeconds)
	data.unmarshalInt("min_withdrawal_amount", &currency.MinWithdrawalAmount)
	data.unmarshalInt("max_withdrawal_amount", &currency.MaxWithdrawalAmount)

	currency.Details = CoinbaseCurrencyDetails{}
	if err := data.unmarshalStruct("details", &currency.Details); err != nil {
		return err
	}

	err = data.unmarshalFloatFromString("min_size", &currency.MinSize)
	if err != nil {
		return err
	}

	err = data.unmarshalFloatFromString("max_precision", &currency.MaxPrecision)
	if err != nil {
		return err
	}

	return nil
}
