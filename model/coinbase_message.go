package model

type CoinbaseMessage struct {
	Message    string `json:"message"`
	Status     string `json:"status"`
	StatusCode string `json:"status_code"`
}
