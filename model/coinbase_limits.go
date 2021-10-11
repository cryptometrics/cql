package model

import "cql/scalar"

type CoinbaseLimits struct {
	Type scalar.PaymentMethod `json:"type"`
	Name string               `json:"name"`
}
