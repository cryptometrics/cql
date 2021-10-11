package model

import (
	"cql/scalar"
	"cql/serial"
	"time"
)

type CoinbasePaymentMethod struct {
	ID                 string                     `json:"id"`
	Type               scalar.PaymentMethod       `json:"type"`
	Name               string                     `json:"name"`
	Currency           string                     `json:"currency"`
	PrimaryBuy         bool                       `json:"primary_buy"`
	PrimarySell        bool                       `json:"primary_sell"`
	InstantBuy         bool                       `json:"instant_buy"`
	InstantSell        bool                       `json:"instant_sell"`
	Verified           bool                       `json:"verified"`
	CreatedAt          time.Time                  `json:"created_at"`
	UpdatedAt          time.Time                  `json:"updated_at"`
	Resource           string                     `json:"resource"`
	ResourcePath       string                     `json:"resource_path"`
	Limits             *CoinbaseLimits            `json:"limits"`
	AllowBuy           bool                       `json:"allow_buy"`
	AllowSell          bool                       `json:"allow_sell"`
	AllowDeposit       bool                       `json:"allow_deposit"`
	AllowWithdraw      bool                       `json:"allow_withdraw"`
	PickerData         *CoinbasePickerData        `json:"picker_data"`
	FiatAccount        *CoinbaseResourceAccount   `json:"fiat_account"`
	CryptoAccount      *CoinbaseResourceAccount   `json:"crypto_account"`
	RecurringOptions   []*CoinbaseRecurringOption `json:"recurring_options"`
	AvailableBalance   *CoinbaseAvailableBalance  `json:"available_balance"`
	HoldBusinessDays   int64                      `json:"hold_business_days"`
	HoldDays           int64                      `json:"hold_days"`
	VerificationMethod string                     `json:"verificationMethod"`
	CDVStatus          string                     `json:"cdvStatus"`
}

func (method *CoinbasePaymentMethod) UnmarshalJSON(d []byte) error {
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}

	data.UnmarshalBool("primary_buy", &method.PrimaryBuy)
	data.UnmarshalBool("primary_sell", &method.PrimarySell)
	data.UnmarshalBool("instant_buy", &method.InstantBuy)
	data.UnmarshalBool("instant_sell", &method.InstantSell)
	data.UnmarshalBool("allow_buy", &method.AllowBuy)
	data.UnmarshalBool("allow_sell", &method.AllowSell)
	data.UnmarshalBool("allow_deposit", &method.AllowDeposit)
	data.UnmarshalBool("allow_withdraw", &method.AllowWithdraw)
	data.UnmarshalBool("verified", &method.Verified)
	data.UnmarshalInt64("hold_business_days", &method.HoldBusinessDays)
	data.UnmarshalInt64("hold_days", &method.HoldDays)
	data.UnmarshalString("id", &method.ID)
	data.UnmarshalString("name", &method.Name)
	data.UnmarshalString("currency", &method.Currency)
	data.UnmarshalString("resource", &method.Resource)
	data.UnmarshalString("resource_path", &method.ResourcePath)
	data.UnmarshalString("verificationMethod", &method.VerificationMethod)
	data.UnmarshalString("cdvStatus", &method.CDVStatus)

	if err := data.UnmarshalPaymentMethod("type", &method.Type); err != nil {
		return err
	}

	if err := data.UnmarshalTime(time.RFC3339Nano, "created_at", &method.CreatedAt); err != nil {
		return err
	}

	if err := data.UnmarshalTime(time.RFC3339Nano, "updated_at", &method.UpdatedAt); err != nil {
		return err
	}

	if data.Value("limits") != nil {
		method.Limits = new(CoinbaseLimits)
		if err := data.UnmarshalStruct("limits", method.Limits); err != nil {
			return err
		}
	}

	if data.Value("picker_data") != nil {
		method.PickerData = new(CoinbasePickerData)
		if err := data.UnmarshalStruct("picker_data", method.PickerData); err != nil {
			return err
		}
	}

	if data.Value("fiat_account") != nil {
		method.FiatAccount = new(CoinbaseResourceAccount)
		if err := data.UnmarshalStruct("fiat_account", method.FiatAccount); err != nil {
			return err
		}
	}

	if data.Value("crypto_account") != nil {
		method.CryptoAccount = new(CoinbaseResourceAccount)
		if err := data.UnmarshalStruct("crypto_account", method.CryptoAccount); err != nil {
			return err
		}
	}

	if cros := data.Value("coinbase_recurring_options"); cros != nil {
		for _, cro := range cros.([]interface{}) {
			if cro != nil {
				method.RecurringOptions = append(method.RecurringOptions, cro.(*CoinbaseRecurringOption))
			}
		}

	}

	if data.Value("available_balance") != nil {
		method.AvailableBalance = new(CoinbaseAvailableBalance)
		if err := data.UnmarshalStruct("available_balance", method.AvailableBalance); err != nil {
			return err
		}
	}

	return nil
}
