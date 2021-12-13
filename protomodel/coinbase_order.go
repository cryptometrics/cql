package protomodel

import (
	"time"

	"github.com/cryptometrics/cql/scalar"
	"github.com/cryptometrics/cql/serial"
)

// * This is a generated file, do not edit

// CoinbaseOrder is an open order.
type CoinbaseOrder struct {
	CreatedAt      time.Time          `json:"created_at"`
	DoneAt         time.Time          `json:"done_at"`
	DoneReason     string             `json:"done_reason"`
	ExecutedValue  float64            `json:"executed_value"`
	ExpireTime     time.Time          `json:"expire_time"`
	FillFees       float64            `json:"fill_fees"`
	FilledSize     float64            `json:"filled_size"`
	FundingAmount  float64            `json:"funding_amount"`
	Funds          float64            `json:"funds"`
	Id             string             `json:"id"`
	PostOnly       bool               `json:"post_only"`
	Price          float64            `json:"price"`
	ProductId      string             `json:"product_id"`
	RejectReason   string             `json:"reject_reason"`
	Settled        bool               `json:"settled"`
	Side           scalar.OrderSide   `json:"side"`
	Size           float64            `json:"size"`
	SpecifiedFunds float64            `json:"specified_funds"`
	Status         string             `json:"status"`
	Stop           string             `json:"stop"`
	StopPrice      float64            `json:"stop_price"`
	TimeInForce    scalar.TimeInForce `json:"time_in_force"`
	Type           scalar.OrderType   `json:"type"`
}

// UnmarshalJSON will deserialize bytes into a CoinbaseOrder model
func (coinbaseOrder *CoinbaseOrder) UnmarshalJSON(d []byte) error {
	const (
		createdAtJsonTag      = "created_at"
		doneAtJsonTag         = "done_at"
		doneReasonJsonTag     = "done_reason"
		executedValueJsonTag  = "executed_value"
		expireTimeJsonTag     = "expire_time"
		fillFeesJsonTag       = "fill_fees"
		filledSizeJsonTag     = "filled_size"
		fundingAmountJsonTag  = "funding_amount"
		fundsJsonTag          = "funds"
		idJsonTag             = "id"
		postOnlyJsonTag       = "post_only"
		priceJsonTag          = "price"
		productIdJsonTag      = "product_id"
		rejectReasonJsonTag   = "reject_reason"
		settledJsonTag        = "settled"
		sideJsonTag           = "side"
		sizeJsonTag           = "size"
		specifiedFundsJsonTag = "specified_funds"
		statusJsonTag         = "status"
		stopJsonTag           = "stop"
		stopPriceJsonTag      = "stop_price"
		timeInForceJsonTag    = "time_in_force"
		typeJsonTag           = "type"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	data.UnmarshalBool(postOnlyJsonTag, &coinbaseOrder.PostOnly)
	data.UnmarshalBool(settledJsonTag, &coinbaseOrder.Settled)
	data.UnmarshalFloatFromString(executedValueJsonTag, &coinbaseOrder.ExecutedValue)
	data.UnmarshalFloatFromString(fillFeesJsonTag, &coinbaseOrder.FillFees)
	data.UnmarshalFloatFromString(filledSizeJsonTag, &coinbaseOrder.FilledSize)
	data.UnmarshalFloatFromString(fundingAmountJsonTag, &coinbaseOrder.FundingAmount)
	data.UnmarshalFloatFromString(fundsJsonTag, &coinbaseOrder.Funds)
	data.UnmarshalFloatFromString(priceJsonTag, &coinbaseOrder.Price)
	data.UnmarshalFloatFromString(sizeJsonTag, &coinbaseOrder.Size)
	data.UnmarshalFloatFromString(specifiedFundsJsonTag, &coinbaseOrder.SpecifiedFunds)
	data.UnmarshalFloatFromString(stopPriceJsonTag, &coinbaseOrder.StopPrice)
	data.UnmarshalOrderSide(sideJsonTag, &coinbaseOrder.Side)
	data.UnmarshalOrderType(typeJsonTag, &coinbaseOrder.Type)
	data.UnmarshalString(doneReasonJsonTag, &coinbaseOrder.DoneReason)
	data.UnmarshalString(idJsonTag, &coinbaseOrder.Id)
	data.UnmarshalString(productIdJsonTag, &coinbaseOrder.ProductId)
	data.UnmarshalString(rejectReasonJsonTag, &coinbaseOrder.RejectReason)
	data.UnmarshalString(statusJsonTag, &coinbaseOrder.Status)
	data.UnmarshalString(stopJsonTag, &coinbaseOrder.Stop)
	data.UnmarshalTimeInForce(timeInForceJsonTag, &coinbaseOrder.TimeInForce)
	err = data.UnmarshalTime(time.RFC3339Nano, createdAtJsonTag, &coinbaseOrder.CreatedAt)
	if err != nil {
		return err
	}
	err = data.UnmarshalTime(time.RFC3339Nano, doneAtJsonTag, &coinbaseOrder.DoneAt)
	if err != nil {
		return err
	}
	err = data.UnmarshalTime(time.RFC3339Nano, expireTimeJsonTag, &coinbaseOrder.ExpireTime)
	if err != nil {
		return err
	}
	return nil
}
