package coinbase

import (
	"strconv"
	"strings"

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

// All will list your current open orders. Only open or un-settled orders are
// returned by default. As soon as an order is no longer open and settled, it
// will no longer appear in the default request. Open orders may change state
// between the request and the response depending on market conditions.
//
// * source: https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_getorders
func (orders *Orders) All(opts *model.CoinbaseOrdersOptions) (m []*model.CoinbaseOrder, err error) {
	return m, orders.Get(FillsEndpoint).
		QueryParam("profile_id", func() (i *string) {
			if opts != nil {
				i = opts.ProfileID
			}
			return
		}()).
		QueryParam("product_id", func() (i *string) {
			if opts != nil && opts.ProductID != nil {
				i = opts.ProductID
			}
			return
		}()).
		QueryParam("sortedBy", func() (i *string) {
			if opts != nil && opts.SortedBy != nil {
				i = opts.SortedBy
			}
			return
		}()).
		QueryParam("sorting", func() (i *string) {
			if opts != nil && opts.Sorting != nil {
				i = opts.Sorting
			}
			return
		}()).
		QueryParam("start_date", func() (i *string) {
			if opts != nil && opts.StartDate != nil {
				tmp := opts.StartDate.String()
				i = &tmp
			}
			return
		}()).
		QueryParam("end_date", func() (i *string) {
			if opts != nil && opts.EndDate != nil {
				tmp := opts.EndDate.String()
				i = &tmp
			}
			return
		}()).
		QueryParam("before", func() (i *string) {
			if opts != nil && opts.Before != nil {
				i = opts.Before
			}
			return
		}()).
		QueryParam("after", func() (i *string) {
			if opts != nil && opts.After != nil {
				i = opts.After
			}
			return
		}()).
		QueryParam("limit", func() (i *string) {
			if opts != nil {
				tmp := strconv.Itoa(opts.Limit)
				i = &tmp
			}
			return
		}()).
		QueryParam("status", func() (i *string) {
			if opts != nil && opts.Status != nil {
				slice := []string{}
				for _, v := range opts.Status {
					slice = append(slice, *v)
				}
				tmp := strings.Join(slice, ", ")
				i = &tmp
			}
			return
		}()).
		Fetch().Assign(&m).JoinMessages()
}

// CancelAll will with best effort, cancel all open orders. This may require you
// to make the request multiple times until all of the open orders are deleted.
//
// * source: https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_deleteorders
func (orders *Orders) CancelAll(opts *model.CoinbaseOrdersOptions) (m []*string, err error) {
	return m, orders.Delete(OrdersEndpoint).
		QueryParam("profile_id", func() (i *string) {
			if opts != nil {
				i = opts.ProfileID
			}
			return
		}()).
		QueryParam("product_id", func() (i *string) {
			if opts != nil && opts.ProductID != nil {
				i = opts.ProductID
			}
			return
		}()).
		Fetch().Assign(&m).JoinMessages()
}

// Create will create an order. You can place two types of orders: limit and market.
// Orders can only be placed if your account has sufficient funds. Once an order
// is placed, your account funds will be put on hold for the duration of the
// order. How much and which funds are put on hold depends on the order type and
// parameters specified.
//
// * source: https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_postorders
func (order *Orders) Create(opts *model.CoinbaseNewOrderOptions) (m *model.CoinbaseNewOrder, err error) {
	return m, order.Post(NewOrderEndpoint).
		Body(client.NewBody(client.BodyTypeJSON).
			SetString("profile_id", opts.ProfileID).
			SetString("product_id", &opts.ProductID).
			SetFloat("stop_price", opts.StopPrice).
			SetFloat("size", opts.Size).
			SetFloat("price", opts.Price).
			SetFloat("funds", opts.Funds).
			SetBool("post_only", opts.PostOnly).
			SetString("client_oid", opts.ClientOid).
			SetString("type", func() *string { str := opts.Type.String(); return &str }()).
			SetString("side", func() *string { str := string(opts.Side); return &str }()).
			SetString("stp", func() *string { str := opts.Stp.String(); return &str }()).
			SetString("stop", func() *string { str := opts.Stop.String(); return &str }()).
			SetString("time_in_force", func() *string { str := opts.TimeInForce.String(); return &str }()).
			SetString("cancel_after", func() *string { str := opts.CancelAfter.String(); return &str }())).
		Fetch().Assign(&m).JoinMessages()
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
			if opts != nil && opts.Before != nil {
				tmp := strconv.Itoa(*opts.Before)
				i = &tmp
			}
			return
		}()).
		QueryParam("after", func() (i *string) {
			if opts != nil && opts.After != nil {
				tmp := strconv.Itoa(*opts.After)
				i = &tmp
			}
			return
		}()).
		Fetch().Assign(&m).JoinMessages()
}
