package protomodel

import "github.com/cryptometrics/cql/serial"

// * This is a generated file, do not edit

// CoinbaseLimits references a FIAT account thata CoinbasePaymentMethod belongs
// to
type CoinbaseFiatAccount struct {
	Id           string `json:"id"`
	Resource     string `json:"resource"`
	ResourcePath string `json:"resource_path"`
}

// UnmarshalJSON will deserialize bytes into a CoinbaseFiatAccount model
func (coinbaseFiatAccount *CoinbaseFiatAccount) UnmarshalJSON(d []byte) error {
	const (
		idJsonTag           = "id"
		resourceJsonTag     = "resource"
		resourcePathJsonTag = "resource_path"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	data.UnmarshalString(idJsonTag, &coinbaseFiatAccount.Id)
	data.UnmarshalString(resourceJsonTag, &coinbaseFiatAccount.Resource)
	data.UnmarshalString(resourcePathJsonTag, &coinbaseFiatAccount.ResourcePath)
	return nil
}
