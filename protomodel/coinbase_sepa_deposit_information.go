package protomodel

import "cql/serial"

// * This is a generated file, do not edit

// CoinbaseSepaDepositInformation information regarding a wallet's deposits.  A
// SEPA credit transfer is a single transfer of Euros from one person or
// organisation to another. For example, this could be to pay the deposit for a
// holiday rental or to settle an invoice. A SEPA direct debit is a recurring
// payment, for example to pay monthly rent or for a service like a mobile phone
// contract.
type CoinbaseSepaDepositInformation struct {
	AccountAddress   string              `json:"account_address"`
	AccountName      string              `json:"account_name"`
	BankAddress      string              `json:"bank_address"`
	BankName         string              `json:"bank_name"`
	Iban             string              `json:"iban"`
	ProtoBankCountry CoinbaseBankCountry `json:"bank_country"`
	Reference        string              `json:"reference"`
	Swift            string              `json:"swift"`
}

// UnmarshalJSON will deserialize bytes into a CoinbaseSepaDepositInformation
// model
func (coinbaseSepaDepositInformation *CoinbaseSepaDepositInformation) UnmarshalJSON(d []byte) error {
	const (
		accountAddressJsonTag = "account_address"
		accountNameJsonTag    = "account_name"
		bankAddressJsonTag    = "bank_address"
		bankCountryJsonTag    = "bank_country"
		bankNameJsonTag       = "bank_name"
		ibanJsonTag           = "iban"
		referenceJsonTag      = "reference"
		swiftJsonTag          = "swift"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	coinbaseSepaDepositInformation.ProtoBankCountry = CoinbaseBankCountry{}
	if err := data.UnmarshalStruct(bankCountryJsonTag, &coinbaseSepaDepositInformation.ProtoBankCountry); err != nil {
		return err
	}
	data.UnmarshalString(accountAddressJsonTag, &coinbaseSepaDepositInformation.AccountAddress)
	data.UnmarshalString(accountNameJsonTag, &coinbaseSepaDepositInformation.AccountName)
	data.UnmarshalString(bankAddressJsonTag, &coinbaseSepaDepositInformation.BankAddress)
	data.UnmarshalString(bankNameJsonTag, &coinbaseSepaDepositInformation.BankName)
	data.UnmarshalString(ibanJsonTag, &coinbaseSepaDepositInformation.Iban)
	data.UnmarshalString(referenceJsonTag, &coinbaseSepaDepositInformation.Reference)
	data.UnmarshalString(swiftJsonTag, &coinbaseSepaDepositInformation.Swift)
	return nil
}
