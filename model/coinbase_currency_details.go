package model

// CoinbaseMarketDataCurrencyDetails are the details of the current CB market
// data currency
type CoinbaseCurrencyDetails struct {
	Type                  string   `json:"type"`
	Symbol                string   `json:"symbol"`
	NetworkConfirmations  int      `json:"network_confirmations"`
	SortOrder             int      `json:"sort_order"`
	CryptoAddressLink     string   `json:"crypto_address_link"`
	CryptoTransactionLink string   `json:"crypto_transaction_link"`
	PushPaymentMethods    []string `json:"push_payment_methods"`
	GroupTypes            []string `json:"group_types"`
	DisplayName           string   `json:"display_name"`
	ProcessingTimeSeconds float64  `json:"processing_time_seconds"`
	MinWithdrawalAmount   float64  `json:"min_withdrawal_amount"`
	MaxWithdrawalAmount   float64  `json:"max_withdrawal_amount"`
}
