package protomodel

import (
	"cql/serial"
	"time"
)

// * This is a generated file, do not edit

// CoinbaseAccountTransfer will lists past withdrawals and deposits for an
// account.
type CoinbaseAccountTransfer struct {
	Amount       float64                        `json:"amount"`
	CanceledAt   time.Time                      `json:"canceled_at"`
	CompletedAt  time.Time                      `json:"completed_at"`
	CreatedAt    time.Time                      `json:"created_at"`
	Id           string                         `json:"id"`
	ProcessedAt  time.Time                      `json:"processed_at"`
	ProtoDetails CoinbaseAccountTransferDetails `json:"details"`
	Type         string                         `json:"type"`
	UserNonce    string                         `json:"user_nonce"`
}

// UnmarshalJSON will deserialize bytes into a CoinbaseAccountTransfer model
func (coinbaseAccountTransfer *CoinbaseAccountTransfer) UnmarshalJSON(d []byte) error {
	const (
		amountJsonTag      = "amount"
		canceledAtJsonTag  = "canceled_at"
		completedAtJsonTag = "completed_at"
		createdAtJsonTag   = "created_at"
		detailsJsonTag     = "details"
		idJsonTag          = "id"
		processedAtJsonTag = "processed_at"
		typeJsonTag        = "type"
		userNonceJsonTag   = "user_nonce"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	coinbaseAccountTransfer.ProtoDetails = CoinbaseAccountTransferDetails{}
	if err := data.UnmarshalStruct(detailsJsonTag, &coinbaseAccountTransfer.ProtoDetails); err != nil {
		return err
	}
	data.UnmarshalFloatFromString(amountJsonTag, &coinbaseAccountTransfer.Amount)
	data.UnmarshalString(idJsonTag, &coinbaseAccountTransfer.Id)
	data.UnmarshalString(typeJsonTag, &coinbaseAccountTransfer.Type)
	data.UnmarshalString(userNonceJsonTag, &coinbaseAccountTransfer.UserNonce)
	err = data.UnmarshalTime(CoinbaseTimeLayout1, canceledAtJsonTag, &coinbaseAccountTransfer.CanceledAt)
	if err != nil {
		return err
	}
	err = data.UnmarshalTime(CoinbaseTimeLayout1, completedAtJsonTag, &coinbaseAccountTransfer.CompletedAt)
	if err != nil {
		return err
	}
	err = data.UnmarshalTime(CoinbaseTimeLayout1, createdAtJsonTag, &coinbaseAccountTransfer.CreatedAt)
	if err != nil {
		return err
	}
	err = data.UnmarshalTime(CoinbaseTimeLayout1, processedAtJsonTag, &coinbaseAccountTransfer.ProcessedAt)
	if err != nil {
		return err
	}
	return nil
}
