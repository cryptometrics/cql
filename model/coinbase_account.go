package model

import (
	"encoding/json"
	"strconv"
)

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
	data := make(umap)
	if err := json.Unmarshal(d, &data); err != nil {
		return err
	}

	var err error

	err = data.unmarshal("id", func(v interface{}) error {
		account.ID = v.(string)
		return nil
	})

	err = data.unmarshal("currency", func(v interface{}) error {
		account.Currency = v.(string)
		return nil
	})

	err = data.unmarshal("balance", func(v interface{}) error {
		account.Balance, err = strconv.ParseFloat(v.(string), 64)
		return err
	})
	if err != nil {
		return err
	}

	err = data.unmarshal("available", func(v interface{}) error {
		account.Available, err = strconv.ParseFloat(v.(string), 64)
		return err
	})
	if err != nil {
		return err
	}

	err = data.unmarshal("hold", func(v interface{}) error {
		account.Hold, err = strconv.ParseFloat(v.(string), 64)
		return err
	})
	if err != nil {
		return err
	}

	err = data.unmarshal("profile_id", func(v interface{}) error {
		account.ProfileID = v.(string)
		return nil
	})

	err = data.unmarshal("trading_enabled", func(v interface{}) error {
		account.TradingEnabled = v.(bool)
		return nil
	})

	return nil
}
