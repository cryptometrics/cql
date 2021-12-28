package protomodel

import "time"

// * This is a generated file, do not edit

// CoinbaseHold represents the hold on an account that belong to the same// profile as the API key. Holds are placed on an account for any active orders// or pending withdraw requests. As an order is filled, the hold amount is// updated. If an order is canceled, any remaining hold is removed. For// withdrawals, the hold is removed after it is completed.
type CoinbaseAccountHold struct {
	CreatedAt time.Time `json:"created_at"`
	Id        string    `json:"id"`
	Ref       string    `json:"ref"`
	Type      string    `json:"type"`
	UpdatedAt time.Time `json:"updated_at"`
}
