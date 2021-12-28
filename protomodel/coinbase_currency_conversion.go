package protomodel

// * This is a generated file, do not edit

// CoinbaseCurrencyConversion is the response that converts funds from from// currency to to currency. Funds are converted on the from account in the// profile_id profile.
type CoinbaseCurrencyConversion struct {
	Amount        float64 `json:"amount"`
	From          string  `json:"from"`
	FromAccountId string  `json:"from_account_id"`
	Id            string  `json:"id"`
	To            string  `json:"to"`
	ToAccountId   string  `json:"to_account_id"`
}
