package model

import (
	"cql/scalar"
	"cql/serial"
	"time"
)

// CoinbaseAccountHistory encapsulates data for account activity of the API
// key's profile. Account activity either increases or decreases your account
// balance. Items are paginated and sorted latest first. See the Pagination
// section for retrieving additional entries after the first page.
//
// If an entry is the result of a trade (match, fee), the details field will
// contain additional information about the trade.
type CoinbaseAccountLedger struct {
	ID        string                       `json:"id"`
	CreatedAt time.Time                    `json:"created_at"`
	Amount    float64                      `json:"amount"`
	Balance   float64                      `json:"balance"`
	Type      scalar.EntryType             `json:"type"`
	Details   CoinbaseAccountLedgerDetails `json:"details"`
}

// UnmarshalJSON is an override required to parst strings from coinbases api
// into floats, specifically min_size and max_precision
func (ledger *CoinbaseAccountLedger) UnmarshalJSON(d []byte) error {
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}

	data.UnmarshalString("id", &ledger.ID)
	data.UnmarshalEntryType("type", &ledger.Type)

	err = data.UnmarshalTime(time.RFC3339Nano, "created_at", &ledger.CreatedAt)
	if err != nil {
		return err
	}

	err = data.UnmarshalFloatFromString("amount", &ledger.Amount)
	if err != nil {
		return err
	}

	err = data.UnmarshalFloatFromString("balance", &ledger.Balance)
	if err != nil {
		return err
	}

	ledger.Details = CoinbaseAccountLedgerDetails{}
	if err := data.UnmarshalStruct("details", &ledger.Details); err != nil {
		return err
	}

	return nil
}
