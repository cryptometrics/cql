package coinbase

import (
	"cql/client"
	"cql/model"
)

type Accounts struct{}

func NewAccounts() *Accounts { return &Accounts{} }

func (accounts *Accounts) all(conn client.Connector) (m []*model.CoinbaseAccount, err error) {
	err = conn.Decode(&client.Request{
		Method: client.GET, Endpoint: AccountsEP}, &m)
	return
}

func (accounts *Accounts) get(conn client.Connector, id string) (m *model.CoinbaseAccount, err error) {
	err = conn.Decode(&client.Request{
		Method:   client.GET,
		Endpoint: AccountEP,
		EndpointArgs: client.EndpointArgs{
			"id": &client.EndpointArg{PathParam: &id}},
	}, &m)
	return
}

func (accounts *Accounts) ledger(conn client.Connector, id string, opts *model.CoinbaseAccountLedgerOptions) (m []*model.CoinbaseAccountLedger, err error) {
	err = conn.Decode(&client.Request{
		Method:   client.GET,
		Endpoint: AccountLedgerEP,
		EndpointArgs: client.EndpointArgs{
			"id": &client.EndpointArg{PathParam: &id},
			"before": &client.EndpointArg{QueryParam: func() (i *string) {
				if opts != nil {
					i = opts.Before
				}
				return
			}()},
			"after": &client.EndpointArg{QueryParam: func() (i *string) {
				if opts != nil {
					i = opts.After
				}
				return
			}()},
			"start_date": &client.EndpointArg{QueryParam: func() (i *string) {
				if opts != nil {
					i = opts.StartDate
				}
				return
			}()},
			"end_date": &client.EndpointArg{QueryParam: func() (i *string) {
				if opts != nil {
					i = opts.EndDate
				}
				return
			}()},
			"limit": &client.EndpointArg{QueryParam: func() (i *string) {
				if opts != nil {
					i = IntPtrStringPtr(opts.Limit)
				}
				return
			}()},
			"profile_id": &client.EndpointArg{QueryParam: func() (i *string) {
				if opts != nil {
					i = opts.ProfileID
				}
				return
			}()},
		},
	}, &m)
	return
}

func (accounts *Accounts) holds(conn client.Connector, id string, opts *model.CoinbaseAccountHoldOptions) (m []*model.CoinbaseAccountHold, err error) {
	err = conn.Decode(&client.Request{
		Method:   client.GET,
		Endpoint: AccountHoldsEP,
		EndpointArgs: client.EndpointArgs{
			"id": &client.EndpointArg{PathParam: &id},
			"before": &client.EndpointArg{QueryParam: func() (i *string) {
				if opts != nil {
					i = opts.Before
				}
				return
			}()},
			"after": &client.EndpointArg{QueryParam: func() (i *string) {
				if opts != nil {
					i = opts.After
				}
				return
			}()},
			"limit": &client.EndpointArg{QueryParam: func() (i *string) {
				if opts != nil {
					i = IntPtrStringPtr(opts.Limit)
				}
				return
			}()},
		},
	}, &m)
	return
}

func (accounts *Accounts) transfers(conn client.Connector, id string, opts *model.CoinbaseAccountTransferOptions) (m []*model.CoinbaseAccountTransfer, err error) {
	err = conn.Decode(&client.Request{
		Method:   client.GET,
		Endpoint: AccountTransfersEP,
		EndpointArgs: client.EndpointArgs{
			"id": &client.EndpointArg{PathParam: &id},
			"before": &client.EndpointArg{QueryParam: func() (i *string) {
				if opts != nil {
					i = opts.Before
				}
				return
			}()},
			"after": &client.EndpointArg{QueryParam: func() (i *string) {
				if opts != nil {
					i = opts.After
				}
				return
			}()},
			"limit": &client.EndpointArg{QueryParam: func() (i *string) {
				if opts != nil {
					i = IntPtrStringPtr(opts.Limit)
				}
				return
			}()},
			"type": &client.EndpointArg{QueryParam: func() (i *string) {
				if opts != nil {
					i = opts.After
				}
				return
			}()},
		},
	}, &m)
	return
}

// All lists trading accounts from the profile of the API key.
//
// Your trading accounts are separate from your Coinbase accounts.  For
// documentation on how to deposit funds to begin trading, see deposits section:
// https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_postdepositcoinbaseaccount
//
// API Key Permissions:
// This endpoint requires either the "view" or "trade" permission.
//
// Rate Limits:
// This endpoint has a custom rate limit by profile ID: 25 requests per second,
// up to 50 requests per second in bursts
//
// Funds on Hold:
// When you place an order, the funds for the order are placed on hold. They
// cannot be used for other orders or withdrawn. Funds will remain on hold until
// the order is filled or canceled.
//
// source: https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_getaccounts
func (accounts *Accounts) All() ([]*model.CoinbaseAccount, error) {
	return accounts.all(newClient)
}

// Find returns information for a single account. Use this endpoint when you
// know the account_id. API key must belong to the same profile as the account.
//
// This endpoint requires either the "view" or "trade" permission.
//
// source: https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_getaccount
func (accounts *Accounts) Find(id string) (*model.CoinbaseAccount, error) {
	return accounts.get(newClient, id)
}

// Holds returns a list of holds of an account that belong to the same profile
// as the API key. Holds are placed on an account for any active orders or
// pending withdraw requests. As an order is filled, the hold amount is updated.
// If an order is canceled, any remaining hold is removed. For a withdraw, once
// it is completed, the hold is removed.
//
// source: https://docs.pro.coinbase.com/#get-hold
func (account *Accounts) Holds(id string, opts *model.CoinbaseAccountHoldOptions) ([]*model.CoinbaseAccountHold, error) {
	return account.holds(newClient, id, opts)
}

// Ledger lists ledger activity for an account. This includes anything that
// would affect the accounts balance - transfers, trades, fees, etc.
//
// This endpoint requires either the "view" or "trade" permission.
func (accounts *Accounts) Ledger(id string, opts *model.CoinbaseAccountLedgerOptions) ([]*model.CoinbaseAccountLedger, error) {
	return accounts.ledger(newClient, id, opts)
}

// Transfers lists past withdrawals and deposits for an account.
func (account *Accounts) Transfers(id string, opts *model.CoinbaseAccountTransferOptions) ([]*model.CoinbaseAccountTransfer, error) {
	return account.transfers(newClient, id, opts)
}
