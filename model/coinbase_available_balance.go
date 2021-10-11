package model

import "cql/serial"

type CoinbaseAvailableBalance struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
	Scale    float64 `json:"scale"`
}

func (balance *CoinbaseAvailableBalance) UnmarshalJSON(d []byte) error {
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}

	data.UnmarshalString("currency", &balance.Currency)

	err = data.UnmarshalFloatFromString("amount", &balance.Amount)
	if err != nil {
		return err
	}

	err = data.UnmarshalFloatFromString("scale", &balance.Scale)
	if err != nil {
		return err
	}

	return nil
}
