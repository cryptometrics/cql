// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

type CoinbaseAccountDepositOptions struct {
	ProfileID         *string `json:"profileId"`
	Amount            float64 `json:"amount"`
	CoinbaseAccountID string  `json:"coinbaseAccountId"`
	Currency          string  `json:"currency"`
}

type CoinbaseAccountHoldsOptions struct {
	Before *string `json:"before"`
	After  *string `json:"after"`
	Limit  *int    `json:"limit"`
}

type CoinbaseAccountLedgerOptions struct {
	StartDate *string `json:"startDate"`
	EndDate   *string `json:"endDate"`
	Before    *string `json:"before"`
	After     *string `json:"after"`
	ProfileID *string `json:"profileId"`
	Limit     *int    `json:"limit"`
}

type CoinbaseAccountTransferOptions struct {
	Before *string `json:"before"`
	After  *string `json:"after"`
	Limit  *int    `json:"limit"`
	Type   *string `json:"type"`
}

type CoinbaseAccountTransfersOptions struct {
	Before *string `json:"before"`
	After  *string `json:"after"`
	Limit  *int    `json:"limit"`
	Type   *string `json:"type"`
}

type CoinbaseAccountWithdrawalOptions struct {
	ProfileID         *string `json:"profileId"`
	Amount            float64 `json:"amount"`
	CoinbaseAccountID string  `json:"coinbaseAccountId"`
	Currency          string  `json:"currency"`
}

type CoinbaseCoinbaseAccountDepositOptions struct {
	ProfileID         *string  `json:"profileId"`
	Amount            *float64 `json:"amount"`
	CoinbaseAccountID *string  `json:"coinbaseAccountId"`
	Currency          *string  `json:"currency"`
}

type CoinbaseConversionOptions struct {
	ProfileID *string `json:"profileId"`
}

type CoinbaseConversionsOptions struct {
	ProfileID *string `json:"profileId"`
	From      string  `json:"from"`
	To        string  `json:"to"`
	Amount    float64 `json:"amount"`
	Nonce     *string `json:"nonce"`
}

type CoinbaseCryptoWithdrawalOptions struct {
	ProfileID        *string  `json:"profileId"`
	Amount           float64  `json:"amount"`
	CryptoAddress    string   `json:"cryptoAddress"`
	Currency         string   `json:"currency"`
	DestinationTag   *string  `json:"destinationTag"`
	NoDestinationTag *bool    `json:"noDestinationTag"`
	TwoFactorCode    *string  `json:"twoFactorCode"`
	Nonce            *int     `json:"nonce"`
	Fee              *float64 `json:"fee"`
}

type CoinbaseFillsOptions struct {
	OrderID   *string `json:"orderId"`
	ProductID *string `json:"productId"`
	ProfileID *string `json:"profileId"`
	Limit     *int    `json:"limit"`
	Before    *int    `json:"before"`
	After     *int    `json:"after"`
}

type CoinbaseOrdersOptions struct {
	ProfileID *string    `json:"profileId"`
	ProductID *string    `json:"productId"`
	SortedBy  *string    `json:"sortedBy"`
	Sorting   *string    `json:"sorting"`
	StartDate *time.Time `json:"startDate"`
	EndDate   *time.Time `json:"endDate"`
	Before    *string    `json:"before"`
	After     *string    `json:"after"`
	Limit     int        `json:"limit"`
	Status    []*string  `json:"status"`
}

type CoinbasePaymentMethodDepositOptions struct {
	ProfileID       *string `json:"profileId"`
	Amount          float64 `json:"amount"`
	PaymentMethodID string  `json:"paymentMethodId"`
	Currency        string  `json:"currency"`
}

type CoinbasePaymentMethodWithdrawalOptions struct {
	ProfileID       *string `json:"profileId"`
	Amount          float64 `json:"amount"`
	PaymentMethodID string  `json:"paymentMethodId"`
	Currency        string  `json:"currency"`
}

type CoinbaseWithdrawalFeeEstimateOptions struct {
	Currency      *string `json:"currency"`
	CryptoAddress *string `json:"cryptoAddress"`
}
