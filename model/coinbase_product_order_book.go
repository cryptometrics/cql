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
	Bids []CoinbaseProductOrderBookBidAsk `json:"bids"`

	// The ask price refers to the lowest price that the owners of that product
	// are willing to sell it for
	Asks []CoinbaseProductOrderBookBidAsk `json:"asks"`
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

	err = data.unmarshal("bids", func(v interface{}) error {
		for _, ibid := range v.([]interface{}) {
			jsonString, _ := json.Marshal(ibid)
			bid := CoinbaseProductOrderBookBidAsk{}
			if err := json.Unmarshal(jsonString, &bid); err != nil {
				return err
			}
			book.Bids = append(book.Bids, bid)
		}
		return nil
	})
	if err != nil {
		return err
	}

	err = data.unmarshal("asks", func(v interface{}) error {
		for _, iask := range v.([]interface{}) {
			jsonString, _ := json.Marshal(iask)
			ask := CoinbaseProductOrderBookBidAsk{}
			if err := json.Unmarshal(jsonString, &ask); err != nil {
				return err
			}
			book.Asks = append(book.Asks, ask)
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
