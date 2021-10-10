package model

import (
	"cql/serial"
	"fmt"
	"time"
)

var CoinbaseAccountTransferTimeLayout = "2006-01-02 15:04:05.999999999+00"

// CoinbaseAccountTransfer represents withdrawals and deposits for an account
type CoinbaseAccountTransfer struct {
	ID          string                         `json:"id"`
	Type        string                         `json:"type"`
	CreatedAt   time.Time                      `json:"created_at"`
	CompletedAt time.Time                      `json:"completed_at"`
	CanceledAt  time.Time                      `json:"canceled_at"`
	ProcessedAt time.Time                      `json:"processed_at"`
	Amount      float64                        `json:"amount"`
	UserNonce   string                         `json:"user_nonce"`
	Details     CoinbaseAccountTransferDetails `json:"details"`
}

// UnmarshalJSON is an override required to parst strings from coinbases api
// into floats, specifically min_size and max_precision
func (transfer *CoinbaseAccountTransfer) UnmarshalJSON(d []byte) error {
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}

	data.UnmarshalString("id", &transfer.ID)
	data.UnmarshalString("type", &transfer.Type)
	data.UnmarshalString("user_nonce", &transfer.UserNonce)
	data.UnmarshalFloatFromString("amount", &transfer.Amount)

	err = data.UnmarshalTime(CoinbaseAccountTransferTimeLayout,
		"created_at", &transfer.CreatedAt)
	if err != nil {
		return err
	}

	err = data.UnmarshalTime(CoinbaseAccountTransferTimeLayout,
		"completed_at", &transfer.CompletedAt)
	if err != nil {
		return err
	}

	err = data.UnmarshalTime(CoinbaseAccountTransferTimeLayout,
		"canceled_at", &transfer.CanceledAt)
	if err != nil {
		return err
	}

	err = data.UnmarshalTime(CoinbaseAccountTransferTimeLayout,
		"processed_at", &transfer.ProcessedAt)
	if err != nil {
		return err
	}

	transfer.Details = CoinbaseAccountTransferDetails{}
	if err := data.UnmarshalStruct("details", &transfer.Details); err != nil {
		return err
	}
	fmt.Println(transfer.Details)

	return nil
}
