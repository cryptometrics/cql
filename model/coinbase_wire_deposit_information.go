package model

type CoinbaseWireDepositInformation struct {
	AccountNumber  string              `json:"account_number"`
	RoutingNumber  string              `json:"routing_number"`
	BankName       string              `json:"bank_name"`
	BankAddress    string              `json:"bank_address"`
	BankCountry    CoinbaseBankCountry `json:"bank_country"`
	AccountName    string              `json:"account_name"`
	AccountAddress string              `json:"account_address"`
	Reference      string              `json:"reference"`
}
