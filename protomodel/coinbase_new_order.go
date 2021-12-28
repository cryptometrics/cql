package protomodel

import (
	"time"

	"github.com/cryptometrics/cql/scalar"
)

// * This is a generated file, do not edit

// CoinbaseNewOrder is the server's response for placing a new order.
type CoinbaseNewOrder struct {
	CreatedAt     time.Time          `json:"created_at"`
	DoneAt        time.Time          `json:"done_at"`
	DoneReason    string             `json:"done_reason"`
	ExpireTime    time.Time          `json:"expire_time"`
	FillFees      float64            `json:"fill_fees"`
	FilledSize    float64            `json:"filled_size"`
	FundingAmount float64            `json:"funding_amount"`
	Funds         float64            `json:"funds"`
	Id            string             `json:"id"`
	PostOnly      bool               `json:"post_only"`
	Price         float64            `json:"price"`
	ProductId     string             `json:"product_id"`
	ProfileId     string             `json:"profile_id"`
	RejectReason  string             `json:"reject_reason"`
	Settled       bool               `json:"settled"`
	Side          scalar.OrderSide   `json:"side"`
	Size          float64            `json:"size"`
	SpecificFunds float64            `json:"specific_funds"`
	Status        string             `json:"status"`
	Stop          scalar.OrderStop   `json:"stop"`
	StopPrice     float64            `json:"stop_price"`
	TimeInForce   scalar.TimeInForce `json:"time_in_force"`
	Type          scalar.OrderType   `json:"type"`
}
