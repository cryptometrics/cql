package coinbase

import (
	"cql/client"
	"cql/model"
	"encoding/json"
	"net/http"
)

type MarketData struct{}

// NewMarketData will return a new object to query the coinbase api for market
// data.
func NewMarketData() *MarketData {
	return &MarketData{}
}

func (md *MarketData) currencies(gen client.Generator) (m []*model.CoinbaseMarketDataCurrency, e error) {
	var cb client.C
	cb, e = gen(client.CoinbasePro)
	if e != nil {
		return
	}
	cb.Connect()

	var res *http.Response
	res, e = cb.Get(CurrenciesEP.Get())
	if e != nil {
		return
	}

	decoder := json.NewDecoder(res.Body)
	e = decoder.Decode(&m)
	return
}

func (md *MarketData) currency(gen client.Generator, id string) (m *model.CoinbaseMarketDataCurrency, e error) {
	var cb client.C
	cb, e = gen(client.CoinbasePro)
	if e != nil {
		return
	}
	cb.Connect()

	var res *http.Response
	res, e = cb.Get(CurrencyEP.Get(id))
	if e != nil {
		return
	}

	decoder := json.NewDecoder(res.Body)
	e = decoder.Decode(&m)
	return
}

// Currencies returns a list of all known currencies on coinbase
func (md *MarketData) Currencies() ([]*model.CoinbaseMarketDataCurrency, error) {
	return md.currencies(client.New)
}

// Currencies returns a single MD currency object for a specified id
func (md *MarketData) Currency(id string) (*model.CoinbaseMarketDataCurrency, error) {
	return md.currency(client.New, id)
}
