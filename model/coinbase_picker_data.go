package model

type CoinbasePickerData struct {
	Symbol                string                    `json:"symbol"`
	CustomerName          string                    `json:"customer_name"`
	AccountName           string                    `json:"account_name"`
	AccountNumber         string                    `json:"account_number"`
	AccountType           string                    `json:"account_type"`
	InstitutionCode       string                    `json:"institution_code"`
	InstitutionName       string                    `json:"institution_name"`
	IBAN                  string                    `json:"iban"`
	SWIFT                 string                    `json:"swift"`
	PaypalEmail           string                    `json:"paypal_email"`
	PaypalOwner           string                    `json:"paypal_owner"`
	RoutingNumber         string                    `json:"routing_number"`
	InstitutionIdentifier string                    `json:"institution_identifier"`
	BankName              string                    `json:"bank_name"`
	BranchName            string                    `json:"branch_name"`
	IconURL               string                    `json:"icon_url"`
	Balance               *CoinbaseAvailableBalance `json:"balance"`
}
