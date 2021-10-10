package model

type CoinbaseDepositAddressWarnings []CoinbaseDepositAddressWarning

type CoinbaseDepositAddressWarning struct {
	Title    string `json:"title"`
	Details  string `json:"details"`
	ImageURL string `json:"image_url"`
}

func (warnings *CoinbaseDepositAddressWarnings) Append(v interface{}) {
	*warnings = append(*warnings, *v.(*CoinbaseDepositAddressWarning))
}

func (warnings *CoinbaseDepositAddressWarnings) UntypedSlice() interface{} {
	w := []*CoinbaseDepositAddressWarning{}
	for _, warning := range *warnings {
		w = append(w, &warning)
	}
	return w
}
