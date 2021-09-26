package model

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// CoinbaseProductOrderBookBidAsk is a slice of values defined by the level passed
// to the client
type CoinbaseProductOrderBookBidAsk struct {
	Price     float64 `json:"price"`
	Size      float64 `json:"size"`
	NumOrders *int    `json:"num_orders"`
	OrderID   *string `json:"order_id"`
}

// UnmarshalJSON is an override required to parst strings from coinbases api
// into floats, specifically min_size and max_precision
func (quote *CoinbaseProductOrderBookBidAsk) UnmarshalJSON(d []byte) error {
	m := []interface{}{}
	if err := json.Unmarshal(d, &m); err != nil {
		return err
	}

	var err error

	quote.Price, err = strconv.ParseFloat(m[0].(string), 64)
	if err != nil {
		return err
	}

	quote.Size, err = strconv.ParseFloat(m[1].(string), 64)
	if err != nil {
		return err
	}

	switch m[2].(type) {
	case string:
		quote.OrderID = String(m[2].(string))
	case float64:
		quote.NumOrders = Int(int(m[2].(float64)))
	default:
		msg := "unknown type unmarshalling CoinbaseProductOrderBookBidAsk"
		return fmt.Errorf(msg)
	}
	return nil
}
