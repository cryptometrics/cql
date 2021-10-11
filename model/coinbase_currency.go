package model

import "cql/serial"

// CoinbaseCurrency holds currency data from the coinbase API.  The
// gqlgen package does not support snake case json, which is why this is
// necessary.
type CoinbaseCurrency struct {
	ID            string                  `json:"id"`
	Name          string                  `json:"name"`
	MinSize       float64                 `json:"min_size"`
	Status        string                  `json:"status"`
	Message       string                  `json:"message"`
	MaxPrecision  float64                 `json:"max_precision"`
	ConvertibleTo []string                `json:"convertible_to"`
	Details       CoinbaseCurrencyDetails `json:"details"`
}

// UnmarshalJSON is an override required to parst strings from coinbases api
// into floats, specifically min_size and max_precision
func (currency *CoinbaseCurrency) UnmarshalJSON(d []byte) error {
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}

	data.UnmarshalString("id", &currency.ID)
	data.UnmarshalString("name", &currency.Name)
	data.UnmarshalString("status", &currency.Status)
	data.UnmarshalString("message", &currency.Message)
	data.UnmarshalStringSlice("convertible_to", &currency.ConvertibleTo)

	err = data.UnmarshalFloatFromString("min_size", &currency.MinSize)
	if err != nil {
		return err
	}

	err = data.UnmarshalFloatFromString("max_precision", &currency.MaxPrecision)
	if err != nil {
		return err
	}

	if err := data.UnmarshalStruct("details", &currency.Details); err != nil {
		return err
	}

	return nil
}
