package coinbase

import (
	"strconv"

	"github.com/cryptometrics/cql/client"
	"github.com/cryptometrics/cql/model"
)

// Orders is a structure used to maintain state while querying on trading
// Orders linked to coinbase.
type Orders struct {
	client.Parent
	conn client.Connector
}

// NewOrders will return a new Orders structure to query on trading Orders
func NewOrders(conn client.Connector) *Orders {
	orders := new(Orders)
	client.ConstructParent(&orders.Parent, conn)
	return orders
}

// WithdrawalFeeEstimate gets the fee estimate for the crypto withdrawal to
// crypto address
//
// This endpoint requires the "transfer" permission. API key must belong to
// default profile.
//
// * source: https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_getwithdrawfeeestimate
func (orders *Orders) Fills(opts *model.CoinbaseFillsOptions) (m []*model.CoinbaseFill, err error) {
	return m, orders.Get(FillsEndpoint).
		QueryParam("order_id", func() (i *string) {
			if opts != nil {
				i = opts.OrderID
			}
			return
		}()).
		QueryParam("product_id", func() (i *string) {
			if opts != nil {
				i = opts.ProductID
			}
			return
		}()).
		QueryParam("profile_id", func() (i *string) {
			if opts != nil {
				i = opts.ProfileID
			}
			return
		}()).
		QueryParam("limit", func() (i *string) {
			if opts != nil && opts.Limit != nil {
				tmp := strconv.Itoa(*opts.Limit)
				i = &tmp
			}
			return
		}()).
		QueryParam("before", func() (i *string) {
			if opts != nil && opts.Limit != nil {
				tmp := strconv.Itoa(*opts.Limit)
				i = &tmp
			}
			return
		}()).
		QueryParam("after", func() (i *string) {
			if opts != nil && opts.Limit != nil {
				tmp := strconv.Itoa(*opts.Limit)
				i = &tmp
			}
			return
		}()).
		Fetch().Assign(&m).JoinMessages()
}
