package model

type CoinbaseUKDepositInformation struct {
	SortCode      string `json:"sort_code"`
	AccountNumber string `json:"account_number"`
	BankName      string `json:"bank_name"`
	AccountName   string `json:"account_name"`
	Reference     string `json:"reference"`
}
