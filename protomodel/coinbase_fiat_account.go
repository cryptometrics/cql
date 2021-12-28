package protomodel

// * This is a generated file, do not edit

// CoinbaseLimits references a FIAT account thata CoinbasePaymentMethod belongs// to
type CoinbaseFiatAccount struct {
	Id           string `json:"id"`
	Resource     string `json:"resource"`
	ResourcePath string `json:"resource_path"`
}
