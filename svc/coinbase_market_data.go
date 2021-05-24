package svc

import (
	"cql/client"
	"cql/model"
	"encoding/json"
	"net/http"
)

type CoinbaseMarketData struct{}

// Currencies returns a list of known currencies on coinbase and decodes them
// into a slice of model.CoinbaseMarketDataCurrency and, if it exists, returns
// an error.
func (md *CoinbaseMarketData) GetCurrencies() (m []*model.CoinbaseMarketDataCurrency, e error) {
	var cb client.C
	cb, e = client.New(client.COINBASE)
	if e != nil {
		return
	}
	cb.Connect()

	var res *http.Response
	res, e = cb.Get("/currencies")
	if e != nil {
		return
	}

	decoder := json.NewDecoder(res.Body)
	e = decoder.Decode(&m)

	return
}
