package protomodel

import (
	"cql/serial"
	"time"
)

// * This is a generated file, do not edit

// CoinbaseHold represents the hold on an account that belong to the same
// profile as the API key. Holds are placed on an account for any active orders
// or pending withdraw requests. As an order is filled, the hold amount is
// updated. If an order is canceled, any remaining hold is removed. For
// withdrawals, the hold is removed after it is completed.
type CoinbaseHold struct {
	CreatedAt time.Time `json:"created_at"`
	Id        string    `json:"id"`
	Ref       string    `json:"ref"`
	Type      string    `json:"type"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UnmarshalJSON will deserialize bytes into a CoinbaseHold model
func (coinbaseHold *CoinbaseHold) UnmarshalJSON(d []byte) error {
	const (
		createdAtJsonTag = "created_at"
		idJsonTag        = "id"
		refJsonTag       = "ref"
		typeJsonTag      = "type"
		updatedAtJsonTag = "updated_at"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	data.UnmarshalString(idJsonTag, &coinbaseHold.Id)
	data.UnmarshalString(refJsonTag, &coinbaseHold.Ref)
	data.UnmarshalString(typeJsonTag, &coinbaseHold.Type)
	err = data.UnmarshalTime(time.RFC3339Nano, createdAtJsonTag, &coinbaseHold.CreatedAt)
	if err != nil {
		return err
	}
	err = data.UnmarshalTime(time.RFC3339Nano, updatedAtJsonTag, &coinbaseHold.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}
