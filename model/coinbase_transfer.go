package model

import (
	"cql/scalar"
	"cql/serial"
	"time"
)

type CoinbaseTransfer struct {
	ID          string                   `json:"id"`
	Type        scalar.TransferMethod    `json:"type"`
	CreatedAt   time.Time                `json:"created_at"`
	CompletedAt time.Time                `json:"completed_at"`
	CanceledAt  time.Time                `json:"canceledAt"`
	ProcessedAt time.Time                `json:"processed_at"`
	Amount      float64                  `json:"amount"`
	Details     *CoinbaseTransferDetails `json:"details"`
	UserNonce   string                   `json:"user_nonce"`
}

func (transfer *CoinbaseTransfer) UnmarshalJSON(d []byte) error {
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	data.UnmarshalString("id", &transfer.ID)
	data.UnmarshalString("user_nonce", &transfer.UserNonce)
	data.UnmarshalFloatFromString("amount", &transfer.Amount)

	if err := data.UnmarshalTransferMethod("type", &transfer.Type); err != nil {
		return err
	}

	layout := CoinbaseAccountTransferTimeLayout

	err = data.UnmarshalTime(layout, "created_at", &transfer.CreatedAt)
	if err != nil {
		return err
	}

	err = data.UnmarshalTime(layout, "completed_at", &transfer.CompletedAt)
	if err != nil {
		return err
	}

	err = data.UnmarshalTime(layout, "canceled_at", &transfer.CanceledAt)
	if err != nil {
		return err
	}

	err = data.UnmarshalTime(layout, "processed_at", &transfer.ProcessedAt)
	if err != nil {
		return err
	}

	if data.Value("details") != nil {
		transfer.Details = new(CoinbaseTransferDetails)
		if err := data.UnmarshalStruct("details", transfer.Details); err != nil {
			return err
		}
	}

	return nil
}
