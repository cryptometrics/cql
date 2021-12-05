package coinbase

import (
	"cql/client"
	"cql/model"
)

// CoinbaseAccounts is an object used to query coinbase account data.
type Conversion struct {
	client.Parent
}

// NewConversion will return an object to query currency conversions
func NewConversion(conn client.Connector) *Conversion {
	conversion := new(Conversion)
	client.ConstructParent(&conversion.Parent, conn)
	return conversion
}

// Find gets currency conversion by id (i.e. USD -> USDC).
//
// * source: https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_getconversion
func (conversion *Conversion) Find(
	conversionId string,
	opts *model.CoinbaseConversionOptions,
) (m *model.CoinbaseCurrencyConversion, err error) {
	return m, conversion.Get(ConversionEndpoint).
		PathParam("conversion_id", conversionId).
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
func (conversion *Conversion) Make(
	opts model.CoinbaseConversionsOptions,
) (m *model.CoinbaseCurrencyConversion, err error) {
	return m, conversion.Post(ConversionsEndpoint).
		Body(client.NewBody(client.BODY_TYPE_JSON).
			SetString("from", &opts.From).
			SetString("to", &opts.To).
			SetFloat("amount", &opts.Amount).
			SetString("profile_id", func() (s *string) {
				if opts.ProfileID != nil {
					s = opts.ProfileID
				}
				return
			}()).
			SetString("nonce", func() (s *string) {
				if opts.Nonce != nil {
					s = opts.Nonce
				}
				return
			}())).
		Fetch().Assign(&m).JoinMessages()
}
