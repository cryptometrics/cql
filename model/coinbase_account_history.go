package model

import (
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
	data, err := newUmap(d)
	if err != nil {
		return err
	}

	data.unmarshalString("id", &history.ID)
	data.unmarshalString("type", &history.Type)

	if err := data.unmarshalTime("created_at", &history.CreatedAt); err != nil {
		return err
	}

	err = data.unmarshalFloatFromString("amount", &history.Amount)
	if err != nil {
		return err
	}

	err = data.unmarshalFloatFromString("balance", &history.Balance)
	if err != nil {
		return err
	}

	history.Details = CoinbaseAccountHistoryDetails{}
	if err := data.unmarshalStruct("details", &history.Details); err != nil {
		return err
	}

	return nil
}
