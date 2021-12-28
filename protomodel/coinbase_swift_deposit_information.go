package protomodel

// * This is a generated file, do not edit

// CoinbaseSwiftDepositInformation information regarding a wallet's deposits.// SWIFT stands for Society for Worldwide Interbank Financial// Telecommunications. Basically, it's a computer network that connects over 900// banks around the world – and enables them to transfer money. ING is part of// this network. There is no fee for accepting deposits into your account with// ING.
type CoinbaseSwiftDepositInformation struct {
	AccountAddress   string              `json:"account_address"`
	AccountName      string              `json:"account_name"`
	AccountNumber    string              `json:"account_number"`
	BankAddress      string              `json:"bank_address"`
	BankName         string              `json:"bank_name"`
	ProtoBankCountry CoinbaseBankCountry `json:"bank_country"`
	Reference        string              `json:"reference"`
}
