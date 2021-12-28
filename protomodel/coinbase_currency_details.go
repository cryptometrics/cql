package protomodel

// * This is a generated file, do not edit

// CoinbaseCurrencyDetails are the details for a currency that coinbase knows// about
type CoinbaseCurrencyDetails struct {
	CryptoAddressLink     string   `json:"crypto_address_link"`
	CryptoTransactionLink string   `json:"crypto_transaction_link"`
	DisplayName           string   `json:"display_name"`
	GroupTypes            []string `json:"group_types"`
	MaxWithdrawalAmount   float64  `json:"max_withdrawal_amount"`
	MinWithdrawalAmount   float64  `json:"min_withdrawal_amount"`
	NetworkConfirmations  int      `json:"network_confirmations"`
	ProcessingTimeSeconds float64  `json:"processing_time_seconds"`
	PushPaymentMethods    []string `json:"push_payment_methods"`
	SortOrder             int      `json:"sort_order"`
	Symbol                string   `json:"symbol"`
	Type                  string   `json:"type"`
}
