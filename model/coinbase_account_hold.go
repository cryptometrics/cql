package model

import "time"

// CoinbaseAccountHold represents a hold of an account that belong to the same
// profile as the API key. Holds are placed on an account for any active orders
// or pending withdraw requests. As an order is filled, the hold amount is
// updated. If an order is canceled, any remaining hold is removed. For a
// withdraw, once it is completed, the hold is removed
type CoinbaseAccountHold struct {
	ID        string    `json:"id"`
	AccountID string    `json:"account_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Amount    float64   `json:"amount"`
	Type      string    `json:"type"`
	Ref       string    `json:"ref"`
}

// UnmarshalJSON is an override required to parst strings from coinbases api
// into floats, specifically min_size and max_precision
func (hold *CoinbaseAccountHold) UnmarshalJSON(d []byte) error {
	data, err := newUmap(d)
	if err != nil {
		return err
	}

	data.unmarshalString("id", &hold.ID)
	data.unmarshalString("account_id", &hold.AccountID)
	data.unmarshalString("type", &hold.Type)
	data.unmarshalString("ref", &hold.Ref)

	if err := data.unmarshalFloatFromString("amount", &hold.Amount); err != nil {
		return err
	}

	err = data.unmarshalTime(time.RFC3339Nano, "created_at", &hold.CreatedAt)
	if err != nil {
		return err
	}

	err = data.unmarshalTime(time.RFC3339Nano, "updated_at", &hold.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}
