package model2

import "cql/serial"

// ! This is a generated file, do not edit

// CoinbaseAccount holds data for trading account from the profile of the API
// key
type CoinbaseAccount struct {
	Available      float64 `json:"available"`
	Balance        float64 `json:"balance"`
	Currency       string  `json:"currency"`
	Hold           float64 `json:"hold"`
	Id             string  `json:"id"`
	ProfileId      string  `json:"profile_id"`
	TradingEnabled bool    `json:"trading_enabled"`
}

// UnmarshalJSON will deserialize bytes into a CoinbaseAccount model
func (coinbaseAccount *CoinbaseAccount) UnmarshalJSON(d []byte) error {
	const (
		jsonTagAvailable      = "available"
		jsonTagBalance        = "balance"
		jsonTagCurrency       = "currency"
		jsonTagHold           = "hold"
		jsonTagId             = "id"
		jsonTagProfileId      = "profile_id"
		jsonTagTradingEnabled = "trading_enabled"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	data.UnmarshalBool(jsonTagTradingEnabled, &coinbaseAccount.TradingEnabled)
	data.UnmarshalFloatFromString(jsonTagAvailable, &coinbaseAccount.Available)
	data.UnmarshalFloatFromString(jsonTagBalance, &coinbaseAccount.Balance)
	data.UnmarshalFloatFromString(jsonTagHold, &coinbaseAccount.Hold)
	data.UnmarshalString(jsonTagCurrency, &coinbaseAccount.Currency)
	data.UnmarshalString(jsonTagId, &coinbaseAccount.Id)
	data.UnmarshalString(jsonTagProfileId, &coinbaseAccount.ProfileId)
	return nil
}
