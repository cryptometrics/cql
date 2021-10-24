package coinbase

import (
	"cql/client2"
	"cql/model"
)

// Conversions is an object used to converts funds from `from` currency to `to`
// currency.
type Conversions struct {
	parent
}

// NewConversion will return a new conversion obejct to convert funds
func NewConversion(conn client2.Connector) *Conversions {
	conversions := new(Conversions)
	construct(&conversions.parent, conn)
	return conversions
}

// Find gets currency conversion by id (i.e. USD -> USDC).
//
// * source: https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_getconversion
func (conversions *Conversions) Find(id string, opts *model.CoinbaseCurrencyConversionOpts,
) (m *model.CoinbaseCurrencyConversion, err error) {
	return m, conversions.get(ENDPOINT_COINBASE_CONVERSION).
		PathParam("id", id).
		QueryParam("profile_id", func() (i *string) {
			if opts != nil {
				i = opts.ProfileID
			}
			return
		}()).
		Fetch().Assign(&m).JoinMessages()
}

// Make converts funds from from currency to to currency. Funds are converted on
// the from account in the profile_id profile.
//
// This endpoint requires the "trade" permission.
//
// A successful conversion will be assigned a conversion id. The correspondin
// ledger entries for a conversion will reference this conversion id.
//
// * source: https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_postconversion
func (conversions *Conversions) Make(from, to string, amount float64, opts *model.CoinbaseCurrencyConversionOpts,
) (m *model.CoinbaseCurrencyConversion, err error) {
	return m, conversions.post(ENDPOINT_COINBASE_CONVERSIONS).
		Body(client2.NewBody(client2.BODY_TYPE_JSON).
			SetString("from", &from).
			SetString("to", &to).
			SetFloat("amount", &amount).
			SetString("profile_id", func() (s *string) {
				if opts != nil {
					s = opts.ProfileID
				}
				return
			}()).
			SetString("nonce", func() (s *string) {
				if opts != nil {
					s = opts.Nonce
				}
				return
			}())).
		Fetch().Assign(&m).JoinMessages()
}
