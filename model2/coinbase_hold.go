package model2

import (
	"cql/serial"
	"time"
)

// ! This is a generated file, do not edit

// CoinbaseHolds list the holds of an account that belong to the same profile as
// the API key. Holds are placed on an account for any active orders or pending
// withdraw requests. As an order is filled, the hold amount is updated. If an
// order is canceled, any remaining hold is removed. For withdrawals, the hold
// is removed after it is complet
type CoinbaseHold struct {
	CreatedAt time.Time `json:"created_at"`
	ID        string    `json:"id"`
	Ref       string    `json:"ref"`
	Type      string    `json:"type"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UnmarshalJSON will deserialize bytes into a CoinbaseHold model
func (coinbaseHold *CoinbaseHold) UnmarshalJSON(d []byte) error {
	const (
		jsonTagCreatedAt = "created_at"
		jsonTagId        = "id"
		jsonTagRef       = "ref"
		jsonTagType      = "type"
		jsonTagUpdatedAt = "updated_at"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	data.UnmarshalString(jsonTagId, &coinbaseHold.ID)
	data.UnmarshalString(jsonTagRef, &coinbaseHold.Ref)
	data.UnmarshalString(jsonTagType, &coinbaseHold.Type)
	err = data.UnmarshalTime(time.RFC3339Nano, "created_at", &coinbaseHold.CreatedAt)
	if err != nil {
		return err
	}
	err = data.UnmarshalTime(time.RFC3339Nano, "updated_at", &coinbaseHold.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}
