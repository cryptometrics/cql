package model

import (
	"encoding/json"
	"strconv"
	"time"
)

// CoinbaseAccountHistory encapsulates data for account activity of the API
// key's profile. Account activity either increases or decreases your account
// balance. Items are paginated and sorted latest first. See the Pagination
// section for retrieving additional entries after the first page.
type CoinbaseAccountHistory struct {
	ID        string                        `json:"id"`
	CreatedAt time.Time                     `json:"created_at"`
	Amount    float64                       `json:"amount"`
	Balance   float64                       `json:"balance"`
	Type      string                        `json:"type"`
	Details   CoinbaseAccountHistoryDetails `json:"details"`
}

// UnmarshalJSON is an override required to parst strings from coinbases api
// into floats, specifically min_size and max_precision
func (history *CoinbaseAccountHistory) UnmarshalJSON(d []byte) error {
	data := make(umap)
	if err := json.Unmarshal(d, &data); err != nil {
		return err
	}

	var err error

	data.unmarshal("id", func(v interface{}) error {
		history.ID = v.(string)
		return nil
	})

	data.unmarshal("created_at", func(v interface{}) error {
		layOut := "2006-01-02T15:04:05.999999999Z07:00"
		history.CreatedAt, err = time.Parse(layOut, v.(string))
		history.CreatedAt = history.CreatedAt.UTC()
		return nil
	})

	data.unmarshal("amount", func(v interface{}) error {
		history.Amount, err = strconv.ParseFloat(v.(string), 64)
		return nil
	})

	data.unmarshal("balance", func(v interface{}) error {
		history.Balance, err = strconv.ParseFloat(v.(string), 64)
		return nil
	})

	data.unmarshal("type", func(v interface{}) error {
		history.Type = v.(string)
		return nil
	})

	err = data.unmarshal("details", func(v interface{}) error {
		history.Details = CoinbaseAccountHistoryDetails{}
		jsonString, _ := json.Marshal(data["details"])
		if err := json.Unmarshal(jsonString, &history.Details); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
