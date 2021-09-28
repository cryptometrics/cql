package coinbase

import (
	"cql/client"
	"cql/model"
)

type Private struct{}

// NewPrivate will return a new object to query the coinbase api for private
// data
func NewPrivate() *Private {
	return &Private{}
}

func (p *Private) accountHistory(gen client.Connector, id string, before, after *int, startDate, endDate *string, limit *int) (m []*model.CoinbaseAccountHistory, err error) {
	err = gen.Decode(&m, AccountHistoryEP, &id, IntString(before),
		IntString(after), startDate, endDate, IntString(limit))
	return
}

func (p *Private) accountHolds(gen client.Connector, id string, before, after, limit *int) (m []*model.CoinbaseAccountHold, err error) {
	err = gen.Decode(&m, AccountHoldsEP, &id, IntString(before), IntString(after),
		IntString(limit))
	return
}

func (p *Private) account(gen client.Connector, id string) (m *model.CoinbaseAccount, err error) {
	err = gen.Decode(&m, AccountEP, &id)
	return
}

func (p *Private) accounts(gen client.Connector) (m []*model.CoinbaseAccount, err error) {
	err = gen.Decode(&m, AccountsEP)
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
