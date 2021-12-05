package coinbase

import (
	"cql/client"
	"path"
	"strings"
)

// * This is a generated file, do not edit

type Endpoint int

const (
	_ Endpoint = iota
	AccountEndpoint
	AccountHoldsEndpoint
	AccountLedgerEndpoint
	AccountTransfersEndpoint
	AccountsEndpoint
	AddressesEndpoint
	ConversionEndpoint
	ConversionsEndpoint
	CurrenciesEndpoint
	CurrencyEndpoint
	WalletsEndpoint
)

// Information for a single account. Use this endpoint when you know the
// account_id. API key must belong to the same profile as the account.
func AccountPath(args client.EndpointArgs) string {
	return path.Join("/accounts", *args["account_id"].PathParam)
}

// List the holds of an account that belong to the same profile as the API key.
// Holds are placed on an account for any active orders or pending withdraw
// requests. As an order is filled, the hold amount is updated. If an order is
// canceled, any remaining hold is removed. For withdrawals, the hold is removed
// after it is completed.
func AccountHoldsPath(args client.EndpointArgs) (p string) {
	p = path.Join("/accounts", *args["account_id"].PathParam, "holds")
	var sb strings.Builder
	sb.WriteString(p)
	sb.WriteString(args.QueryPath().String())
	return sb.String()
}

// Lists ledger activity for an account. This includes anything that would
// affect the accounts balance - transfers, trades, fees, etc. This endpoint
// requires either the "view" or "trade" permission.
func AccountLedgerPath(args client.EndpointArgs) (p string) {
	p = path.Join("/accounts", *args["account_id"].PathParam, "ledger")
	var sb strings.Builder
	sb.WriteString(p)
	sb.WriteString(args.QueryPath().String())
	return sb.String()
}

// Lists past withdrawals and deposits for an account.
func AccountTransfersPath(args client.EndpointArgs) (p string) {
	p = path.Join("/accounts", *args["account_id"].PathParam, "transfers")
	var sb strings.Builder
	sb.WriteString(p)
	sb.WriteString(args.QueryPath().String())
	return sb.String()
}

// Get a list of trading accounts from the profile of the API key.
func AccountsPath(_ client.EndpointArgs) string {
	return path.Join("/accounts")
}

// Generates a one-time crypto address for depositing crypto, using a wallet
// account id.  This endpoint requires the "transfer" permission. API key must
// belong to default profile.
func AddressesPath(args client.EndpointArgs) string {
	return path.Join("/coinbase-accounts", *args["account_id"].PathParam, "addresses")
}

// Gets a currency conversion by id (i.e. USD -> USDC).
func ConversionPath(args client.EndpointArgs) (p string) {
	p = path.Join("/conversions", *args["conversion_id"].PathParam)
	var sb strings.Builder
	sb.WriteString(p)
	sb.WriteString(args.QueryPath().String())
	return sb.String()
}

// Converts funds from from currency to to currency. Funds are converted on the
// from account in the profile_id profile.  This endpoint requires the "trade"
// permission.  A successful conversion will be assigned a conversion id. The
// corresponding ledger entries for a conversion will reference this conversion
// id
func ConversionsPath(args client.EndpointArgs) (p string) {
	p = path.Join("/conversions")
	var sb strings.Builder
	sb.WriteString(p)
	sb.WriteString(args.QueryPath().String())
	return sb.String()
}

// Gets a list of all known currencies.  Note: Not all currencies may be
// currently in use for trading.
func CurrenciesPath(_ client.EndpointArgs) string {
	return path.Join("/currencies")
}

// Gets a single currency by id.
func CurrencyPath(args client.EndpointArgs) string {
	return path.Join("/currencies", *args["currency_id"].PathParam)
}

// Gets all the user's available Coinbase wallets (These are the
// wallets/accounts that are used for buying and selling on www.coinbase.com)
func WalletsPath(_ client.EndpointArgs) string {
	return path.Join("/coinbase-accounts")
}

// Get takes an endpoint const and endpoint arguments to parse the URL endpoint
// path.
func (endpoint Endpoint) Path(args client.EndpointArgs) string {
	return map[Endpoint]func(args client.EndpointArgs) string{
		AccountEndpoint:          AccountPath,
		AccountHoldsEndpoint:     AccountHoldsPath,
		AccountLedgerEndpoint:    AccountLedgerPath,
		AccountTransfersEndpoint: AccountTransfersPath,
		AccountsEndpoint:         AccountsPath,
		AddressesEndpoint:        AddressesPath,
		ConversionEndpoint:       ConversionPath,
		ConversionsEndpoint:      ConversionsPath,
		CurrenciesEndpoint:       CurrenciesPath,
		CurrencyEndpoint:         CurrencyPath,
		WalletsEndpoint:          WalletsPath,
	}[endpoint](args)
}
