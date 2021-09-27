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

// Account the trading accounts from the profile of the API key, given the
// account id
func (p *Private) Account(id string) (*model.CoinbaseAccount, error) {
	return p.account(newClient, id)
}

// Accounts returns a list of trading accounts from the profile of the API key
func (p *Private) Accounts() ([]*model.CoinbaseAccount, error) {
	return p.accounts(newClient)
}
