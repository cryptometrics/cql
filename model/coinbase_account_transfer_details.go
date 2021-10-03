package model

type CoinbaseAccountTransferDetails struct {
	CoinbaseAccountID       string `json:"coinbase_account_id"`
	CoinbaseTransactionID   string `json:"coinbase_transaction_id"`
	CoinbasePaymentMethodID string `json:"coinbase_payment_method_id"`
}
