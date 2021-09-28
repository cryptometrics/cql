package model

import (
	"fmt"
)

// CoinbaseProductOrderBookBidAsk is the object encapsulation of the a list of
// values defined by the level passed to the client
type CoinbaseProductOrderBookBidAsk struct {
	Price     float64 `json:"price"`
	Size      float64 `json:"size"`
	NumOrders *int    `json:"num_orders"`
	OrderID   *string `json:"order_id"`
}

type CoinbaseProductOrderBookBidAsks []CoinbaseProductOrderBookBidAsk

func (quotes *CoinbaseProductOrderBookBidAsks) append(v interface{}) {
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
	data, err := newUslice(d)
	if err != nil {
		return err
	}

	if err := data.unmarshalFloatFromString(0, &quote.Price); err != nil {
		return err
	}

	if err := data.unmarshalFloatFromString(1, &quote.Size); err != nil {
		return err
	}

	switch data[2].(type) {
	case string:
		quote.OrderID = String(data[2].(string))
	case float64:
		quote.NumOrders = Int(int(data[2].(float64)))
	default:
		msg := "unknown type unmarshalling CoinbaseProductOrderBookBidAsk"
		return fmt.Errorf(msg)
	}

	return nil
}
