package protomodel

import "cryptometrics/cql/serial"

// * This is a generated file, do not edit

// CoinbaseWireDepositInformation information regarding a wallet's deposits
type CoinbaseWireDepositInformation struct {
	AccountAddress   string              `json:"account_address"`
	AccountName      string              `json:"account_name"`
	AccountNumber    string              `json:"account_number"`
	BankAddress      string              `json:"bank_address"`
	BankName         string              `json:"bank_name"`
	ProtoBankCountry CoinbaseBankCountry `json:"bank_country"`
	Reference        string              `json:"reference"`
	RoutingNumber    string              `json:"routing_number"`
}

// UnmarshalJSON will deserialize bytes into a CoinbaseWireDepositInformation
// model
func (coinbaseWireDepositInformation *CoinbaseWireDepositInformation) UnmarshalJSON(d []byte) error {
	const (
		accountAddressJsonTag = "account_address"
		accountNameJsonTag    = "account_name"
		accountNumberJsonTag  = "account_number"
		bankAddressJsonTag    = "bank_address"
		bankCountryJsonTag    = "bank_country"
		bankNameJsonTag       = "bank_name"
		referenceJsonTag      = "reference"
		routingNumberJsonTag  = "routing_number"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	coinbaseWireDepositInformation.ProtoBankCountry = CoinbaseBankCountry{}
	if err := data.UnmarshalStruct(bankCountryJsonTag, &coinbaseWireDepositInformation.ProtoBankCountry); err != nil {
		return err
	}
	data.UnmarshalString(accountAddressJsonTag, &coinbaseWireDepositInformation.AccountAddress)
	data.UnmarshalString(accountNameJsonTag, &coinbaseWireDepositInformation.AccountName)
	data.UnmarshalString(accountNumberJsonTag, &coinbaseWireDepositInformation.AccountNumber)
	data.UnmarshalString(bankAddressJsonTag, &coinbaseWireDepositInformation.BankAddress)
	data.UnmarshalString(bankNameJsonTag, &coinbaseWireDepositInformation.BankName)
	data.UnmarshalString(referenceJsonTag, &coinbaseWireDepositInformation.Reference)
	data.UnmarshalString(routingNumberJsonTag, &coinbaseWireDepositInformation.RoutingNumber)
	return nil
}
