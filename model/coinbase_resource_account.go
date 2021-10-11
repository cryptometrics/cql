package model

type CoinbaseResourceAccount struct {
	ID           string `json:"id"`
	Resource     string `json:"resource"`
	ResourcePath string `json:"resource_path"`
}
