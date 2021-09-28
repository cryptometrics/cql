package model

import (
	"encoding/json"
)

// CoinbaseProductOrderBook holds bid/ask data as a list of open orders for a
// product. The amount of detail shown can be customized with the level
// parameter
type CoinbaseProductOrderBook struct {
	Sequence int `json:"sequence"`

	// Bid prices refer to the highest price that traders are willing to pay for a
	// product
	Bids CoinbaseProductOrderBookBidAsks `json:"bids"`

	// The ask price refers to the lowest price that the owners of that product
	// are willing to sell it for
	Asks CoinbaseProductOrderBookBidAsks `json:"asks"`
}

// UnmarshalJSON is an override required to parst strings from coinbases api
// into floats, specifically min_size and max_precision
func (book *CoinbaseProductOrderBook) UnmarshalJSON(d []byte) error {
	data := make(umap)
	if err := json.Unmarshal(d, &data); err != nil {
		return err
	}

	var err error

	err = data.unmarshal("sequence", func(v interface{}) error {
		book.Sequence = int(v.(float64))
		return nil
	})

	err = data.unmarshalStructSlice("bids", &book.Bids,
		&CoinbaseProductOrderBookBidAsk{})
	if err != nil {
		return err
	}

	err = data.unmarshalStructSlice("asks", &book.Asks,
		&CoinbaseProductOrderBookBidAsk{})
	if err != nil {
		return err
	}

	return nil
}
