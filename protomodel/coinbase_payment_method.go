package protomodel

import (
	"cryptometrics/cql/serial"
	"encoding/json"
	"time"
)

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

// UnmarshalJSON will deserialize bytes into a CoinbasePaymentMethod model
func (coinbasePaymentMethod *CoinbasePaymentMethod) UnmarshalJSON(d []byte) error {
	const (
		allowBuyJsonTag           = "allow_buy"
		allowDepositJsonTag       = "allow_deposit"
		allowSellJsonTag          = "allow_sell"
		allowWithdrawJsonTag      = "allow_withdraw"
		availableBalanceJsonTag   = "available_balance"
		cdvStatusJsonTag          = "cdv_status"
		createAtJsonTag           = "create_at"
		cryptoAccountJsonTag      = "crypto_account"
		currencyJsonTag           = "currency"
		fiatAccountJsonTag        = "fiat_account"
		holdBusinessDaysJsonTag   = "hold_business_days"
		holdDaysJsonTag           = "hold_days"
		idJsonTag                 = "id"
		instantBuyJsonTag         = "instant_buy"
		instantSaleJsonTag        = "instant_sale"
		limitsJsonTag             = "limits"
		nameJsonTag               = "name"
		pickerDataJsonTag         = "picker_data"
		primaryBuyJsonTag         = "primary_buy"
		primarySellJsonTag        = "primary_sell"
		recurringOptionsJsonTag   = "recurring_options"
		resourceJsonTag           = "resource"
		resourcePathJsonTag       = "resource_path"
		typeJsonTag               = "type"
		updatedAtJsonTag          = "updated_at"
		verificationMethodJsonTag = "verification_method"
		verifiedJsonTag           = "verified"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	coinbasePaymentMethod.ProtoAvailableBalance = CoinbaseAvailableBalance{}
	if err := data.UnmarshalStruct(availableBalanceJsonTag, &coinbasePaymentMethod.ProtoAvailableBalance); err != nil {
		return err
	}
	coinbasePaymentMethod.ProtoCryptoAccount = CoinbaseCryptoAccount{}
	if err := data.UnmarshalStruct(cryptoAccountJsonTag, &coinbasePaymentMethod.ProtoCryptoAccount); err != nil {
		return err
	}
	coinbasePaymentMethod.ProtoFiatAccount = CoinbaseFiatAccount{}
	if err := data.UnmarshalStruct(fiatAccountJsonTag, &coinbasePaymentMethod.ProtoFiatAccount); err != nil {
		return err
	}
	coinbasePaymentMethod.ProtoLimits = CoinbaseLimits{}
	if err := data.UnmarshalStruct(limitsJsonTag, &coinbasePaymentMethod.ProtoLimits); err != nil {
		return err
	}
	coinbasePaymentMethod.ProtoPickerData = CoinbasePickerData{}
	if err := data.UnmarshalStruct(pickerDataJsonTag, &coinbasePaymentMethod.ProtoPickerData); err != nil {
		return err
	}
	data.UnmarshalBool(allowBuyJsonTag, &coinbasePaymentMethod.AllowBuy)
	data.UnmarshalBool(allowDepositJsonTag, &coinbasePaymentMethod.AllowDeposit)
	data.UnmarshalBool(allowSellJsonTag, &coinbasePaymentMethod.AllowSell)
	data.UnmarshalBool(allowWithdrawJsonTag, &coinbasePaymentMethod.AllowWithdraw)
	data.UnmarshalBool(instantBuyJsonTag, &coinbasePaymentMethod.InstantBuy)
	data.UnmarshalBool(instantSaleJsonTag, &coinbasePaymentMethod.InstantSale)
	data.UnmarshalBool(primaryBuyJsonTag, &coinbasePaymentMethod.PrimaryBuy)
	data.UnmarshalBool(primarySellJsonTag, &coinbasePaymentMethod.PrimarySell)
	data.UnmarshalBool(verifiedJsonTag, &coinbasePaymentMethod.Verified)
	data.UnmarshalInt(holdBusinessDaysJsonTag, &coinbasePaymentMethod.HoldBusinessDays)
	data.UnmarshalInt(holdDaysJsonTag, &coinbasePaymentMethod.HoldDays)
	data.UnmarshalString(cdvStatusJsonTag, &coinbasePaymentMethod.CdvStatus)
	data.UnmarshalString(currencyJsonTag, &coinbasePaymentMethod.Currency)
	data.UnmarshalString(idJsonTag, &coinbasePaymentMethod.Id)
	data.UnmarshalString(nameJsonTag, &coinbasePaymentMethod.Name)
	data.UnmarshalString(resourceJsonTag, &coinbasePaymentMethod.Resource)
	data.UnmarshalString(resourcePathJsonTag, &coinbasePaymentMethod.ResourcePath)
	data.UnmarshalString(typeJsonTag, &coinbasePaymentMethod.Type)
	data.UnmarshalString(verificationMethodJsonTag, &coinbasePaymentMethod.VerificationMethod)
	err = data.UnmarshalTime(time.RFC3339Nano, createAtJsonTag, &coinbasePaymentMethod.CreateAt)
	if err != nil {
		return err
	}
	err = data.UnmarshalTime(time.RFC3339Nano, updatedAtJsonTag, &coinbasePaymentMethod.UpdatedAt)
	if err != nil {
		return err
	}
	if v := data.Value(recurringOptionsJsonTag); v != nil {
		for _, item := range data.Value(recurringOptionsJsonTag).([]interface{}) {
			byt, _ := json.Marshal(item)
			obj := CoinbaseRecurringOptions{}
			if err := json.Unmarshal(byt, &obj); err != nil {
				return err
			}
			coinbasePaymentMethod.ProtoRecurringOptions = append(coinbasePaymentMethod.ProtoRecurringOptions, &obj)
		}
	}
	return nil
}
