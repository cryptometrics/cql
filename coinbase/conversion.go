package coinbase

import (
	"cql/client"
	"cql/model"
	"cql/null"
)

// Conversions is an object used to converts funds from `from` currency to `to`
// currency.
type Conversions struct {
	parent
}

// NewConversion will return a new conversion obejct to convert funds
func NewConversion(conn client.Connector) *Conversions {
	conversions := new(Conversions)
	construct(&conversions.parent, conn)
	return conversions
}

func (conversions *Conversions) fetchConversion(id string, opts *model.CoinbaseCurrencyConversionOpts) *client.FetchResponse {
	return conversions.conn.Fetch(&client.Request{
		Method:   client.GET,
		Endpoint: CoinbaseConversionEP,
		EndpointArgs: client.EndpointArgs{
			"id": &client.EndpointArg{PathParam: &id},
			"profile_id": &client.EndpointArg{QueryParam: func() (i *string) {
				if opts != nil {
					i = opts.ProfileID
				}
				return
			}()},
		},
	})
}

func (conversions *Conversions) fetchMake(from, to string, amount float64, opts *model.CoinbaseCurrencyConversionOpts) *client.FetchResponse {
	return conversions.conn.Fetch(&client.Request{
		Method:   client.POST,
		Endpoint: CoinbaseConversionsEP,
		Body: client.JSONBody(map[string]null.Interface{
			"from":   null.NewInterfaceString(&from),
			"to":     null.NewInterfaceString(&to),
			"amount": null.NewInterfaceFloat(&amount),
			"profile_id": null.NewInterfaceString(func() (s *string) {
				if opts != nil {
					s = opts.ProfileID
				}
				return
			}()),
			"nonce": null.NewInterfaceString(func() (s *string) {
				if opts != nil {
					s = opts.Nonce
				}
				return
			}()),
		}),
	})
}

// Get gets currency conversion by id (i.e. USD -> USDC).
func (conversions *Conversions) Get(id string, opts *model.CoinbaseCurrencyConversionOpts) (m *model.CoinbaseCurrencyConversion, err error) {
	return m, conversions.fetchConversion(id, opts).Assign(&m).Error
}

// Make converts funds from from currency to to currency. Funds are converted on the from account in
// the profile_id profile.
//
// This endpoint requires the "trade" permission.
//
// A successful conversion will be assigned a conversion id. The correspondin ledger entries for a
// conversion will reference this conversion id.
//
// source: https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_postconversion
func (conversions *Conversions) Make(from, to string, amount float64, opts *model.CoinbaseCurrencyConversionOpts) (m *model.CoinbaseCurrencyConversion, err error) {
	return m, conversions.fetchMake(from, to, amount, opts).Assign(&m).Error
}
