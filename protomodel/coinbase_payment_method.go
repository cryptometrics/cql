package protomodel

import "time"

// * This is a generated file, do not edit

// CoinbasePaymentMethod is a payment method used on coinbase
type CoinbasePaymentMethod struct {
	AllowBuy              bool                        `json:"allow_buy"`
	AllowDeposit          bool                        `json:"allow_deposit"`
	AllowSell             bool                        `json:"allow_sell"`
	AllowWithdraw         bool                        `json:"allow_withdraw"`
	CdvStatus             string                      `json:"cdv_status"`
	CreateAt              time.Time                   `json:"create_at"`
	Currency              string                      `json:"currency"`
	HoldBusinessDays      int                         `json:"hold_business_days"`
	HoldDays              int                         `json:"hold_days"`
	Id                    string                      `json:"id"`
	InstantBuy            bool                        `json:"instant_buy"`
	InstantSale           bool                        `json:"instant_sale"`
	Name                  string                      `json:"name"`
	PrimaryBuy            bool                        `json:"primary_buy"`
	PrimarySell           bool                        `json:"primary_sell"`
	ProtoAvailableBalance CoinbaseAvailableBalance    `json:"available_balance"`
	ProtoCryptoAccount    CoinbaseCryptoAccount       `json:"crypto_account"`
	ProtoFiatAccount      CoinbaseFiatAccount         `json:"fiat_account"`
	ProtoLimits           CoinbaseLimits              `json:"limits"`
	ProtoPickerData       CoinbasePickerData          `json:"picker_data"`
	ProtoRecurringOptions []*CoinbaseRecurringOptions `json:"recurring_options"`
	Resource              string                      `json:"resource"`
	ResourcePath          string                      `json:"resource_path"`
	Type                  string                      `json:"type"`
	UpdatedAt             time.Time                   `json:"updated_at"`
	VerificationMethod    string                      `json:"verification_method"`
	Verified              bool                        `json:"verified"`
}
