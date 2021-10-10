package model

import (
	"cql/scalar"
	"cql/serial"
	"time"
)

// CoinbaseOrder is the response to an order request through the cb api.
type CoinbaseOrder struct {
	ID             string             `json:"id"`
	Price          float64            `json:"price"`
	Size           float64            `json:"size"`
	ProductID      string             `json:"product_id"`
	Side           scalar.OrderSide   `json:"side"`
	STP            scalar.OrderSTP    `json:"stp"`
	Type           scalar.OrderType   `json:"type"`
	TimeInForce    scalar.TimeInForce `json:"time_in_force"`
	PostOnly       bool               `json:"post_only"`
	CreatedAt      time.Time          `json:"created_at"`
	FillFees       float64            `json:"fill_fees"`
	FilledSize     float64            `json:"filled_size"`
	ExecutedValue  float64            `json:"executed_value"`
	Status         string             `json:"status"`
	Settled        bool               `json:"settled"`
	Funds          float64            `json:"funds"`
	SpecifiedFunds float64            `json:"specified_funds"`
}

// UnmarshalJSON is an override required to parst strings from coinbases api
// into floats, specifically min_size and max_precision
func (order *CoinbaseOrder) UnmarshalJSON(d []byte) error {
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}

	data.UnmarshalOrderType("type", &order.Type)
	data.UnmarshalOrderSide("side", &order.Side)
	data.UnmarshalOrderTimeInForce("time_in_force", &order.TimeInForce)
	data.UnmarshalString("id", &order.ID)
	data.UnmarshalString("product_id", &order.ProductID)
	data.UnmarshalOrderSTP("stp", &order.STP)
	data.UnmarshalBool("post_only", &order.PostOnly)
	data.UnmarshalString("status", &order.Status)
	data.UnmarshalBool("settled", &order.Settled)

	if err := data.UnmarshalFloatFromString("price", &order.Price); err != nil {
		return err
	}

	if err := data.UnmarshalFloatFromString("size", &order.Size); err != nil {
		return err
	}

	err = data.UnmarshalTime(time.RFC3339Nano, "created_at", &order.CreatedAt)
	if err != nil {
		return err
	}

	err = data.UnmarshalFloatFromString("fill_fees", &order.FillFees)
	if err != nil {
		return err
	}

	err = data.UnmarshalFloatFromString("filled_size", &order.FilledSize)
	if err != nil {
		return err
	}

	err = data.UnmarshalFloatFromString("executed_value", &order.ExecutedValue)
	if err != nil {
		return err
	}

	err = data.UnmarshalFloatFromString("funds", &order.Funds)
	if err != nil {
		return err
	}

	err = data.UnmarshalFloatFromString("specified_funds", &order.SpecifiedFunds)
	if err != nil {
		return err
	}

	return nil
}
