package model

type CoinbaseSEPADepositInformation struct {
	IBAN           string              `json:"iban"`
	SWIFT          string              `json:"swift"`
	BankName       string              `json:"bank_name"`
	BankAddress    string              `json:"bank_address"`
	BankCountry    CoinbaseBankCountry `json:"bank_country"`
	AccountName    string              `json:"account_name"`
	AccountAddress string              `json:"account_address"`
	Reference      string              `json:"reference"`
}
