package protomodel

import "cql/serial"

// * This is a generated file, do not edit

// CoinbaseCurrencyDetails are the details for a currency that coinbase knows
// about
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

// UnmarshalJSON will deserialize bytes into a CoinbaseCurrencyDetails model
func (coinbaseCurrencyDetails *CoinbaseCurrencyDetails) UnmarshalJSON(d []byte) error {
	const (
		cryptoAddressLinkJsonTag     = "crypto_address_link"
		cryptoTransactionLinkJsonTag = "crypto_transaction_link"
		displayNameJsonTag           = "display_name"
		groupTypesJsonTag            = "group_types"
		maxWithdrawalAmountJsonTag   = "max_withdrawal_amount"
		minWithdrawalAmountJsonTag   = "min_withdrawal_amount"
		networkConfirmationsJsonTag  = "network_confirmations"
		processingTimeSecondsJsonTag = "processing_time_seconds"
		pushPaymentMethodsJsonTag    = "push_payment_methods"
		sortOrderJsonTag             = "sort_order"
		symbolJsonTag                = "symbol"
		typeJsonTag                  = "type"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	data.UnmarshalFloat(maxWithdrawalAmountJsonTag, &coinbaseCurrencyDetails.MaxWithdrawalAmount)
	data.UnmarshalFloat(minWithdrawalAmountJsonTag, &coinbaseCurrencyDetails.MinWithdrawalAmount)
	data.UnmarshalFloat(processingTimeSecondsJsonTag, &coinbaseCurrencyDetails.ProcessingTimeSeconds)
	data.UnmarshalInt(networkConfirmationsJsonTag, &coinbaseCurrencyDetails.NetworkConfirmations)
	data.UnmarshalInt(sortOrderJsonTag, &coinbaseCurrencyDetails.SortOrder)
	data.UnmarshalString(cryptoAddressLinkJsonTag, &coinbaseCurrencyDetails.CryptoAddressLink)
	data.UnmarshalString(cryptoTransactionLinkJsonTag, &coinbaseCurrencyDetails.CryptoTransactionLink)
	data.UnmarshalString(displayNameJsonTag, &coinbaseCurrencyDetails.DisplayName)
	data.UnmarshalString(symbolJsonTag, &coinbaseCurrencyDetails.Symbol)
	data.UnmarshalString(typeJsonTag, &coinbaseCurrencyDetails.Type)
	data.UnmarshalStringSlice(groupTypesJsonTag, &coinbaseCurrencyDetails.GroupTypes)
	data.UnmarshalStringSlice(pushPaymentMethodsJsonTag, &coinbaseCurrencyDetails.PushPaymentMethods)
	return nil
}
