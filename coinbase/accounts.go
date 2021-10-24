package coinbase

import (
	"cql/client2"
	"cql/model"
)

// Accounts is a structure used to maintain state while querying on trading
// accounts linked to coinbase.
type Accounts struct {
	parent
	conn client2.Connector
}

// NewAccounts will return a new accounts structure to query on trading accounts
func NewAccounts(conn client2.Connector) *Accounts {
	accounts := new(Accounts)
	construct(&accounts.parent, conn)
	return accounts
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
func (accounts *Accounts) All() (m []*model.CoinbaseAccount, err error) {
	req := accounts.get(ENDPOINT_ACCOUNTS)
	return m, req.Fetch().Assign(&m).JoinMessages()
}

// Find returns information for a single account. Use this endpoint when you
// know the account_id. API key must belong to the same profile as the account.
//
// This endpoint requires either the "view" or "trade" permission.
//
// * source: https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_getaccount
func (accounts *Accounts) Find(id string) (m *model.CoinbaseAccount, err error) {
	req := accounts.get(ENDPOINT_ACCOUNT)
	return m, req.PathParam("id", id).Fetch().Assign(&m).JoinMessages()
}

// Holds returns a list of holds of an account that belong to the same profile
// as the API key. Holds are placed on an account for any active orders or
// pending withdraw requests. As an order is filled, the hold amount is updated.
// If an order is canceled, any remaining hold is removed. For a withdraw, once
// it is completed, the hold is removed.
//
// * source: https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_getaccountholds
func (accounts *Accounts) Holds(id string, opts *model.CoinbaseAccountHoldOptions,
) (m []*model.CoinbaseAccountHold, err error) {
	return m, accounts.get(ENDPOINT_ACCOUNT).
		PathParam("id", id).
		QueryParam("before", func() (i *string) {
			if opts != nil {
				i = opts.Before
			}
			return
		}()).
		QueryParam("after", func() (i *string) {
			if opts != nil {
				i = opts.After
			}
			return
		}()).
		QueryParam("limit", func() (i *string) {
			if opts != nil {
				i = IntPtrStringPtr(opts.Limit)
			}
			return
		}()).
		Fetch().Assign(&m).JoinMessages()
}

// Ledger lists ledger activity for an account. This includes anything that
// would affect the accounts balance - transfers, trades, fees, etc.
//
// This endpoint requires either the "view" or "trade" permission.
//
// * source: https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_getaccountledger
func (accounts *Accounts) Ledger(id string, opts *model.CoinbaseAccountLedgerOptions,
) (m []*model.CoinbaseAccountLedger, err error) {
	return m, accounts.get(ENDPOINT_ACCOUNT_LEDGER).
		PathParam("id", id).
		QueryParam("before", func() (i *string) {
			if opts != nil {
				i = opts.Before
			}
			return
		}()).
		QueryParam("after", func() (i *string) {
			if opts != nil {
				i = opts.After
			}
			return
		}()).
		QueryParam("start_date", func() (i *string) {
			if opts != nil {
				i = opts.StartDate
			}
			return
		}()).
		QueryParam("end_date", func() (i *string) {
			if opts != nil {
				i = opts.EndDate
			}
			return
		}()).
		QueryParam("limit", func() (i *string) {
			if opts != nil {
				i = IntPtrStringPtr(opts.Limit)
			}
			return
		}()).
		QueryParam("profile_id", func() (i *string) {
			if opts != nil {
				i = opts.ProfileID
			}
			return
		}()).
		Fetch().Assign(&m).JoinMessages()
}

// Transfers lists past withdrawals and deposits for an account.
//
// * source: https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_getaccounttransfers
func (accounts *Accounts) Transfers(id string, opts *model.CoinbaseAccountTransferOptions,
) (m []*model.CoinbaseAccountTransfer, err error) {
	return m, accounts.get(ENDPOINT_ACCOUNT_TRANSFERS).
		PathParam("id", id).
		QueryParam("before", func() (i *string) {
			if opts != nil {
				i = opts.Before
			}
			return
		}()).
		QueryParam("after", func() (i *string) {
			if opts != nil {
				i = opts.After
			}
			return
		}()).
		QueryParam("limit", func() (i *string) {
			if opts != nil {
				i = IntPtrStringPtr(opts.Limit)
			}
			return
		}()).
		QueryParam("limit", func() (i *string) {
			if opts != nil {
				i = IntPtrStringPtr(opts.Limit)
			}
			return
		}()).
		QueryParam("type", func() (i *string) {
			if opts != nil {
				i = opts.After
			}
			return
		}()).
		Fetch().Assign(&m).JoinMessages()
}
