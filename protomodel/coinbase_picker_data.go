package protomodel

// * This is a generated file, do not edit

// CoinbasePickerData ??
type CoinbasePickerData struct {
	AccountName           string          `json:"account_name"`
	AccountNumber         string          `json:"account_number"`
	AccountType           string          `json:"account_type"`
	BankName              string          `json:"bank_name"`
	BranchName            string          `json:"branch_name"`
	CustomerName          string          `json:"customer_name"`
	Iban                  string          `json:"iban"`
	IconUrl               string          `json:"icon_url"`
	InstitutionCode       string          `json:"institution_code"`
	InstitutionIdentifier string          `json:"institution_identifier"`
	InstitutionName       string          `json:"institution_name"`
	PaypalEmail           string          `json:"paypal_email"`
	PaypalOwner           string          `json:"paypal_owner"`
	ProtoBalance          CoinbaseBalance `json:"balance"`
	RoutingNumber         string          `json:"routing_number"`
	Swift                 string          `json:"swift"`
	Symbol                string          `json:"symbol"`
}
