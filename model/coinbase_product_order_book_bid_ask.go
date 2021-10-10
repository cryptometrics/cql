package model

import (
	"encoding/json"
	"fmt"
	"strconv"

	"gopkg.in/guregu/null.v3"
)

// CoinbaseProductOrderBookBidAsk is the object encapsulation of the a list of
// values defined by the level passed to the client
type CoinbaseProductOrderBookBidAsk struct {
	Price     float64 `json:"price"`
	Size      float64 `json:"size"`
	NumOrders *int64  `json:"num_orders"`
	OrderID   *string `json:"order_id"`
}

type CoinbaseProductOrderBookBidAsks []CoinbaseProductOrderBookBidAsk

func (quotes *CoinbaseProductOrderBookBidAsks) Append(v interface{}) {
	*quotes = append(*quotes, *v.(*CoinbaseProductOrderBookBidAsk))
}

func (quotes *CoinbaseProductOrderBookBidAsks) Slice() (s []*CoinbaseProductOrderBookBidAsk) {
	for _, quote := range *quotes {
		s = append(s, &quote)
	}
	return
}

// UnmarshalJSON is an override required to parst strings from coinbases api
// into floats, specifically min_size and max_precision
func (quote *CoinbaseProductOrderBookBidAsk) UnmarshalJSON(d []byte) error {
	data := []interface{}{}
	if err := json.Unmarshal(d, &data); err != nil {
		return err
	}

	var err error

	if quote.Price, err = strconv.ParseFloat(data[0].(string), 64); err != nil {
		return err
	}

	if quote.Size, err = strconv.ParseFloat(data[1].(string), 64); err != nil {
		return err
	}

	switch data[2].(type) {
	case string:
		quote.OrderID = null.StringFrom(data[2].(string)).Ptr()
	case float64:
		quote.NumOrders = null.IntFrom(int64(data[2].(float64))).Ptr()
	default:
		msg := "unknown type unmarshalling CoinbaseProductOrderBookBidAsk"
		return fmt.Errorf(msg)
	}

	return nil
}

func (quotes *CoinbaseProductOrderBookBidAsks) UntypedSlice() interface{} {
	q := []*CoinbaseProductOrderBookBidAsk{}
	for _, quote := range *quotes {
		q = append(q, &quote)
	}
	return q
}
