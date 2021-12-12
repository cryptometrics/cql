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

// ## API Key Permissions
// This endpoint requires either the "view" or "trade" permission.
//
// ## Settlement and Fees
// Fees are recorded in two stages. Immediately after the matching engine
// completes a match, the fill is inserted into our datastore. Once the fill is
// recorded, a settlement process will settle the fill and credit both trading
// counterparties.
//
// The fee field indicates the fees charged for this individual fill.
//
// ### Liquidity
// The liquidity field indicates if the fill was the result of a liquidity
// provider or liquidity taker. M indicates Maker and T indicates Taker.
//
// ### Pagination
// Fills are returned sorted by descending trade_id from the largest trade_id to
// the smallest trade_id. The CB-BEFORE header will have this first trade id so
// that future requests using the cb-before parameter will fetch fills with a
// greater trade id (newer fills).
// * source: https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_getfills
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