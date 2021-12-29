package protomodel

import "github.com/cryptometrics/cql/serial"

// * This is a generated file, do not edit

// CoinbaseAccount holds data for trading account from the profile of the API// key
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
		idJsonTag             = "id"
		currencyJsonTag       = "currency"
		balanceJsonTag        = "balance"
		availableJsonTag      = "available"
		holdJsonTag           = "hold"
		profileIdJsonTag      = "profile_id"
		tradingEnabledJsonTag = "trading_enabled"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	data.UnmarshalBool(tradingEnabledJsonTag, &coinbaseAccount.TradingEnabled)
	data.UnmarshalFloat(availableJsonTag, &coinbaseAccount.Available)
	data.UnmarshalFloat(balanceJsonTag, &coinbaseAccount.Balance)
	data.UnmarshalFloat(holdJsonTag, &coinbaseAccount.Hold)
	data.UnmarshalString(currencyJsonTag, &coinbaseAccount.Currency)
	data.UnmarshalString(idJsonTag, &coinbaseAccount.Id)
	data.UnmarshalString(profileIdJsonTag, &coinbaseAccount.ProfileId)
	return nil
}
