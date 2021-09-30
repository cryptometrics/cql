package model

import (
	"cql/scalar"
	"time"
)

// CoinbaseOrder is the response to an order request through the cb api.
type CoinbaseOrder struct {
	ID            string             `json:"id"`
	Price         float64            `json:"price"`
	Size          float64            `json:"size"`
	ProductID     string             `json:"product_id"`
	Side          scalar.OrderSide   `json:"side"`
	STP           scalar.OrderSTP    `json:"stp"`
	Type          scalar.OrderType   `json:"type"`
	TimeInForce   scalar.TimeInForce `json:"time_in_force"`
	PostOnly      bool               `json:"post_only"`
	CreatedAt     time.Time          `json:"created_at"`
	FillFees      float64            `json:"fill_fees"`
	FilledSize    float64            `json:"filled_size"`
	ExecutedValue float64            `json:"executed_value"`
	Status        string             `json:"status"`
	Settled       bool               `json:"settled"`
}

// UnmarshalJSON is an override required to parst strings from coinbases api
// into floats, specifically min_size and max_precision
func (order *CoinbaseOrder) UnmarshalJSON(d []byte) error {
	data, err := newUmap(d)
	if err != nil {
		return err
	}

	data.unmarshalOrderType("type", &order.Type)
	data.unmarshalOrderSide("side", &order.Side)
	data.unmarshalOrderTimeInForce("time_in_force", &order.TimeInForce)
	data.unmarshalString("id", &order.ID)
	data.unmarshalString("product_id", &order.ProductID)
	data.unmarshalOrderSTP("stp", &order.STP)
	data.unmarshalBool("post_only", &order.PostOnly)
	data.unmarshalString("status", &order.Status)
	data.unmarshalBool("settled", &order.Settled)

	if err := data.unmarshalFloatFromString("price", &order.Price); err != nil {
		return err
	}

	if err := data.unmarshalFloatFromString("size", &order.Size); err != nil {
		return err
	}

	if err := data.unmarshalTime("created_at", &order.CreatedAt); err != nil {
		return err
	}

	err = data.unmarshalFloatFromString("fill_fees", &order.FillFees)
	if err != nil {
		return err
	}

	err = data.unmarshalFloatFromString("filled_size", &order.FilledSize)
	if err != nil {
		return err
	}

	err = data.unmarshalFloatFromString("executed_value", &order.ExecutedValue)
	if err != nil {
		return err
	}

	return nil
}
