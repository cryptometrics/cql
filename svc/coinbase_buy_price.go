package svc

import (
	"cql/client"
	"cql/model"
	"encoding/json"
	"fmt"
	"net/http"
)

type CoinbaseBuyPrice struct{}

// Price gets the total price to buy one bitcoin or ether.
func (bp *CoinbaseBuyPrice) Price(currencyPair string) (p *model.CoinbaseBuyPrice, e error) {
	var cb client.C
	cb, e = client.New(client.CoinbasePro)
	if e != nil {
		return
	}
	cb.Connect()

	var res *http.Response
	// TODO need to fix the way we set path params.  This is smoothy
	res, e = cb.Get(fmt.Sprintf("/prices/%s/buy", currencyPair))
	if e != nil {
		return
	}
	return nil, json.NewDecoder(res.Body).Decode(&p)
}
