package coinbase

import (
	"cql/client"
	"cql/model"
	"cql/null"
	"io"
)

type Private struct{}

// NewPrivate will return a new object to query the coinbase api for private
// data
func NewPrivate() *Private {
	return &Private{}
}

func (p *Private) accountHistory(gen client.Connector, id string, before, after *int, startDate, endDate *string, limit *int) (m []*model.CoinbaseAccountHistory, err error) {
	err = gen.Decode(client.DecodeInput{
		Method:   client.GET,
		Endpoint: AccountHistoryEP,
		EndpointArgs: client.EndpointArgs{
			"id":         &client.EndpointArg{PathParam: &id},
			"before":     &client.EndpointArg{QueryParam: IntPtrStringPtr(before)},
			"after":      &client.EndpointArg{QueryParam: IntPtrStringPtr(after)},
			"start_date": &client.EndpointArg{QueryParam: startDate},
			"end_date":   &client.EndpointArg{QueryParam: endDate},
			"limit":      &client.EndpointArg{QueryParam: IntPtrStringPtr(limit)}},
	}, &m)
	return
}

func (p *Private) accountHolds(gen client.Connector, id string, before, after, limit *int) (m []*model.CoinbaseAccountHold, err error) {
	err = gen.Decode(client.DecodeInput{
		Method:   client.GET,
		Endpoint: AccountHoldsEP,
		EndpointArgs: client.EndpointArgs{
			"id":     &client.EndpointArg{PathParam: &id},
			"before": &client.EndpointArg{QueryParam: IntPtrStringPtr(before)},
			"after":  &client.EndpointArg{QueryParam: IntPtrStringPtr(after)},
			"limit":  &client.EndpointArg{QueryParam: IntPtrStringPtr(limit)}},
	}, &m)
	return
}

func (p *Private) account(gen client.Connector, id string) (m *model.CoinbaseAccount, err error) {
	err = gen.Decode(client.DecodeInput{
		Method:   client.GET,
		Endpoint: AccountEP,
		EndpointArgs: client.EndpointArgs{
			"id": &client.EndpointArg{PathParam: &id}},
	}, &m)
	return
}

func (p *Private) accounts(gen client.Connector) (m []*model.CoinbaseAccount, err error) {
	err = gen.Decode(client.DecodeInput{
		Method: client.GET, Endpoint: AccountsEP}, &m)
	return
}

func (p *Private) getLimitOrderBody(input *model.CoinbaseOrderInput) io.Reader {
	return nil
}

func (p *Private) createLimitOrder(gen client.Connector, input *model.CoinbaseOrderInput) (m *model.CoinbaseOrder, err error) {
	err = gen.Decode(client.DecodeInput{
		Method:   client.POST,
		Endpoint: CreateOrderEP,
		Body: client.JSONBody(map[string]null.Interface{
			"client_oid":    null.NewInterfaceString(input.ClientOid),
			"type":          null.NewInterfaceString(StringStringPtr("limit")),
			"side":          null.NewInterfaceString((*string)(&input.Side)),
			"product_id":    null.NewInterfaceString(&input.ProductID),
			"stp":           null.NewInterfaceString((*string)(input.Stp)),
			"stop":          null.NewInterfaceString((*string)(input.Stop)),
			"stop_price":    null.NewInterfaceString(input.StopPrice),
			"price":         null.NewInterfaceFloat(&input.Price),
			"size":          null.NewInterfaceFloat(&input.Size),
			"time_in_force": null.NewInterfaceString((*string)(input.TimeInForce)),
		})}, &m)
	return
}

// AccountHistory returns a list of account activity of the API key's profile
func (p *Private) AccountHistory(id string, before, after *int, startDate, endDate *string, limit *int) ([]*model.CoinbaseAccountHistory, error) {
	return p.accountHistory(newClient, id, before, after, startDate, endDate, limit)
}

// AccountHolds returns a list of holds of an account that belong to the same
// profile as the API key. Holds are placed on an account for any active orders
// or pending withdraw requests. As an order is filled, the hold amount is
// updated. If an order is canceled, any remaining hold is removed. For a
// withdraw, once it is completed, the hold is removed.
//
// before - Pagination parameter that requires a positive integer. If set, returns
// a list of holds before the specified integer.
//
// after - Pagination parameter that requires a positive integer. If set, returns
// a list of holds after the specified integer.
//
// limit - Number of results per request. Maximum 1000. (default 1000)
//
// source: https://docs.pro.coinbase.com/#get-holds
func (p *Private) AccountHolds(id string, before, after, limit *int) ([]*model.CoinbaseAccountHold, error) {
	return p.accountHolds(newClient, id, before, after, limit)
}

// Account the trading accounts from the profile of the API key, given the
// account id
func (p *Private) Account(id string) (*model.CoinbaseAccount, error) {
	return p.account(newClient, id)
}

// Accounts returns a list of trading accounts from the profile of the API key
func (p *Private) Accounts() ([]*model.CoinbaseAccount, error) {
	return p.accounts(newClient)
}

// CreateLimitOrder puts in a limit order, returning a CoinbaseOrder struct in
// response.  Limit orders are both the default and basic order type. A limit
// order requires specifying a price and size. The size is the number of base
// currency to buy or sell, and the price is the price per base currency. The
// limit order will be filled at the price specified or better. A sell order can
// be filled at the specified price per base currency or a higher price per base
// currency and a buy order can be filled at the specified price or a lower
// price depending on market conditions. If market conditions cannot fill the
// limit order immediately, then the limit order will become part of the open
// order book until filled by another incoming order or canceled by the user.
//
// For limit buy orders, we will hold price x size x (1 + fee-percent) USD. For
// sell orders, we will hold the number of base currency you wish to sell.
// Actual fees are assessed at time of trade. If you cancel a partially filled
// or unfilled order, any remaining funds will be released from hold.
func (p *Private) CreateLimitOrder(input *model.CoinbaseOrderInput) (*model.CoinbaseOrder, error) {
	return p.createLimitOrder(newClient, input)
}
