package model

import (
	"cql/scalar"
	"cql/serial"
)

// CoinbaseMarketDataCurrencyDetails are the details of the current CB market
// data currency
type CoinbaseCurrencyDetails struct {
	Type                  string                 `json:"type"`
	Symbol                string                 `json:"symbol"`
	NetworkConfirmations  int                    `json:"network_confirmations"`
	SortOrder             int                    `json:"sort_order"`
	CryptoAddressLink     string                 `json:"crypto_address_link"`
	CryptoTransactionLink string                 `json:"crypto_transaction_link"`
	PushPaymentMethods    []scalar.PaymentMethod `json:"push_payment_methods"`
	GroupTypes            []string               `json:"group_types"`
	DisplayName           string                 `json:"display_name"`
	ProcessingTimeSeconds float64                `json:"processing_time_seconds"`
	MinWithdrawalAmount   float64                `json:"min_withdrawal_amount"`
	MaxWithdrawalAmount   float64                `json:"max_withdrawal_amount"`
}

func (detail *CoinbaseCurrencyDetails) UnmarshalJSON(d []byte) error {
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}

	data.UnmarshalInt("network_confirmations", &detail.NetworkConfirmations)
	data.UnmarshalInt("sort_order", &detail.SortOrder)
	data.UnmarshalString("type", &detail.Type)
	data.UnmarshalString("symbol", &detail.Symbol)
	data.UnmarshalString("crypto_address_link", &detail.CryptoAddressLink)
	data.UnmarshalString("crypto_transaction_link", &detail.CryptoTransactionLink)
	data.UnmarshalString("display_name", &detail.DisplayName)
	data.UnmarshalFloat("min_withdrawal_amount", &detail.MinWithdrawalAmount)
	data.UnmarshalFloat("max_withdrawal_amount", &detail.MinWithdrawalAmount)
	data.UnmarshalFloat("processing_time_seconds", &detail.ProcessingTimeSeconds)

	if err := data.UnmarshalStringSlice("group_types", &detail.GroupTypes); err != nil {
		return err
	}

	if pushPaymentMethods := data.Value("push_payment_methods"); pushPaymentMethods != nil {
		for _, pm := range data.Value("push_payment_methods").([]interface{}) {
			detail.PushPaymentMethods = append(detail.PushPaymentMethods, scalar.PaymentMethod(pm.(string)))
		}
	}

	return nil
}
