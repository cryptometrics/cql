package protomodel

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
