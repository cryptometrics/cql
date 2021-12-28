package protomodel

import "time"

// * This is a generated file, do not edit

// CoinbaseAccountTransfer will lists past withdrawals and deposits for an// account.
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
