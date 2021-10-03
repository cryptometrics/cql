package coinbase

import (
	"cql/client"
	"cql/model"
	"cql/null"

	"github.com/sirupsen/logrus"
)

type Private struct{}

// NewPrivate will return a new object to query the coinbase api for private
// data
func NewPrivate() *Private {
	return &Private{}
}

func (p *Private) clientOrder(conn client.Connector, clientOID string) (m *model.CoinbaseOrder, err error) {
	err = conn.Decode(&client.Request{
		Method:   client.GET,
		Endpoint: ClientOrderEP,
		EndpointArgs: client.EndpointArgs{
			"client_oid": &client.EndpointArg{PathParam: &clientOID}},
	}, &m)
	return
}

func (p *Private) createLimitOrder(conn client.Connector, input *model.CoinbaseOrderInput) (m *model.CoinbaseOrder, err error) {
	err = conn.Decode(&client.Request{
		Method:   client.POST,
		Endpoint: CreateOrderEP,
		Callback: func(i interface{}, r *client.Request) error {
			order := i.(**model.CoinbaseOrder)
			logrus.Debug(r.Logf(`{order_id:%s}`, (*order).ID))
			return nil
		},
		Body: client.JSONBody(map[string]null.Interface{
			"client_oid":    null.NewInterfaceString(input.ClientOid),
			"type":          null.NewInterfaceString(StringStringPtr("limit")),
			"side":          null.NewInterfaceString((*string)(&input.Side)),
			"product_id":    null.NewInterfaceString(&input.ProductID),
			"stp":           null.NewInterfaceString((*string)(input.Stp)),
			"stop":          null.NewInterfaceString((*string)(input.Stop)),
			"stop_price":    null.NewInterfaceString(input.StopPrice),
			"price":         null.NewInterfaceFloat(&input.Price),
			"size":          null.NewInterfaceFloat(&input.Size),
			"time_in_force": null.NewInterfaceString((*string)(input.TimeInForce)),
		})}, &m)
	return
}

func (p *Private) order(conn client.Connector, id string) (m *model.CoinbaseOrder, err error) {
	err = conn.Decode(&client.Request{
		Method:   client.GET,
		Endpoint: OrderEP,
		EndpointArgs: client.EndpointArgs{
			"id": &client.EndpointArg{PathParam: &id}},
	}, &m)
	return
}

func (p *Private) orders(conn client.Connector, params *model.CoinbaseOrderQueryParameters) (m *model.CoinbaseOrder, err error) {
	err = conn.Decode(&client.Request{
		Method:   client.GET,
		Endpoint: OrderEP,
		EndpointArgs: client.EndpointArgs{
			"status": &client.EndpointArg{QueryParam: SlicePtrStringPtr(params.Statuses)},
			// "product_id": &client.EndpointArg{QueryParam: params.ProductID},
			// "start_date": &client.EndpointArg{QueryParam: params.StartDate},
			// "end_date":   &client.EndpointArg{QueryParam: params.EndDate},
			// "before": &flient.EndpointArgs{Query
		},
	}, &m)
	return
}

// ClientOrder returns the order for a client's OID
func (p *Private) ClientOrder(clientOID string) (*model.CoinbaseOrder, error) {
	return p.clientOrder(newClient, clientOID)
}

// CreateLimitOrder puts in a limit order, returning a CoinbaseOrder struct in
// response.  Limit orders are both the default and basic order type. A limit
// order requires specifying a price and size. The size is the number of base
// currency to buy or sell, and the price is the price per base currency. The
// limit order will be filled at the price specified or better. A sell order can
// be filled at the specified price per base currency or a higher price per base
// currency and a buy order can be filled at the specified price or a lower
// price depending on market conditions. If market conditions cannot fill the
// limit order immediately, then the limit order will become part of the open
// order book until filled by another incoming order or canceled by the user.
//
// For limit buy orders, we will hold price x size x (1 + fee-percent) USD. For
// sell orders, we will hold the number of base currency you wish to sell.
// Actual fees are assessed at time of trade. If you cancel a partially filled
// or unfilled order, any remaining funds will be released from hold.
func (p *Private) CreateLimitOrder(input *model.CoinbaseOrderInput) (*model.CoinbaseOrder, error) {
	return p.createLimitOrder(newClient, input)
}

// Order will get a single order by order id from the profile that the API key
// belongs to
func (p *Private) Order(id string) (*model.CoinbaseOrder, error) {
	return p.order(newClient, id)
}
