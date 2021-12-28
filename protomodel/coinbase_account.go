package protomodel

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
