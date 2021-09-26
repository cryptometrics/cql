package model

import (
	"encoding/json"
	"strconv"
)

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
	m := make(map[string]interface{})
	if err := json.Unmarshal(d, &m); err != nil {
		return err
	}

	setter := func(name string, fn func(interface{}) error) error {
		if v := m[name]; v != nil {
			if err := fn(v); err != nil {
				return err
			}
		}
		return nil
	}

	var err error

	setter("id", func(v interface{}) error {
		currency.ID = v.(string)
		return nil
	})

	setter("name", func(v interface{}) error {
		currency.Name = v.(string)
		return nil
	})

	setter("status", func(v interface{}) error {
		currency.Status = v.(string)
		return nil
	})

	setter("message", func(v interface{}) error {
		currency.Message = v.(string)
		return nil
	})

	setter("convertible_to", func(v interface{}) error {
		for _, ct := range v.([]interface{}) {
			currency.ConvertibleTo = append(currency.ConvertibleTo, ct.(string))
		}
		return nil
	})

	setter("display_name", func(v interface{}) error {
		currency.DisplayName = v.(string)
		return nil
	})

	setter("processing_time_seconds", func(v interface{}) error {
		currency.ProcessingTimeSeconds = int(v.(float64))
		return nil
	})

	setter("min_withdrawal_amount", func(v interface{}) error {
		currency.MinWithdrawalAmount = int(v.(float64))
		return nil
	})

	setter("max_withdrawal_amount", func(v interface{}) error {
		currency.MaxWithdrawalAmount = int(v.(float64))
		return nil
	})

	err = setter("details", func(v interface{}) error {
		currency.Details = CoinbaseCurrencyDetails{}
		jsonString, _ := json.Marshal(m["details"])
		if err := json.Unmarshal(jsonString, &currency.Details); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	setter("min_size", func(v interface{}) error {
		strFloat := v.(string)
		currency.MinSize, err = strconv.ParseFloat(strFloat, 64)
		return err
	})
	if err != nil {
		return err
	}

	setter("max_precision", func(v interface{}) error {
		strFloat := v.(string)
		currency.MaxPrecision, err = strconv.ParseFloat(strFloat, 64)
		return err
	})
	if err != nil {
		return err
	}

	return nil
}
