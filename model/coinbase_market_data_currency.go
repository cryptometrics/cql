package model

// CoinbaseMarketDataCurrency holds currency data from the coinbase API.  The
// gqlgen package does not support snake case json, which is why this is
// necessary.
type CoinbaseMarketDataCurrency struct {
	ID                    string                            `json:"id"`
	Name                  string                            `json:"name"`
	MinSize               string                            `json:"min_size"`
	Status                string                            `json:"status"`
	Message               string                            `json:"message"`
	MaxPrecision          string                            `json:"max_precision"`
	ConvertibleTo         []string                          `json:"convertible_to"`
	Details               CoinbaseMarketDataCurrencyDetails `json:"details"`
	DisplayName           string                            `json:"display_name"`
	ProcessingTimeSeconds int                               `json:"processing_time_seconds"`
	MinWithdrawalAmount   int                               `json:"min_withdrawal_amount"`
	MaxWithdrawalAmount   int                               `json:"max_withdrawal_amount"`
}
