package protomodel

import "cryptometrics/cql/serial"

// * This is a generated file, do not edit

// CoinbaseCurrencyConversion is the response that converts funds from from
// currency to to currency. Funds are converted on the from account in the
// profile_id profile.
type CoinbaseCurrencyConversion struct {
	Amount        float64 `json:"amount"`
	From          string  `json:"from"`
	FromAccountId string  `json:"from_account_id"`
	Id            string  `json:"id"`
	To            string  `json:"to"`
	ToAccountId   string  `json:"to_account_id"`
}

// UnmarshalJSON will deserialize bytes into a CoinbaseCurrencyConversion model
func (coinbaseCurrencyConversion *CoinbaseCurrencyConversion) UnmarshalJSON(d []byte) error {
	const (
		amountJsonTag        = "amount"
		fromAccountIdJsonTag = "from_account_id"
		fromJsonTag          = "from"
		idJsonTag            = "id"
		toAccountIdJsonTag   = "to_account_id"
		toJsonTag            = "to"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	data.UnmarshalFloatFromString(amountJsonTag, &coinbaseCurrencyConversion.Amount)
	data.UnmarshalString(fromAccountIdJsonTag, &coinbaseCurrencyConversion.FromAccountId)
	data.UnmarshalString(fromJsonTag, &coinbaseCurrencyConversion.From)
	data.UnmarshalString(idJsonTag, &coinbaseCurrencyConversion.Id)
	data.UnmarshalString(toAccountIdJsonTag, &coinbaseCurrencyConversion.ToAccountId)
	data.UnmarshalString(toJsonTag, &coinbaseCurrencyConversion.To)
	return nil
}
