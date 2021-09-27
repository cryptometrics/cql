package model

import (
	"encoding/json"
	"strconv"
	"time"
)

// CoinbaseProductTrade lists the latest trades for a product
type CoinbaseProductTrade struct {
	Time    time.Time `json:"time"`
	TradeID int       `json:"trade_id"`
	Price   float64   `json:"price"`
	Size    float64   `json:"size"`

	// side indicates the maker order side. The maker order is the order that was
	// open on the order book. buy side indicates a down-tick because the maker
	// was a buy order and their order was removed. Conversely, sell side
	// indicates an up-tick.
	Side string `json:"side"`
}

// UnmarshalJSON is an override required to parst strings from coinbases api
// into floats, specifically min_size and max_precision
func (trade *CoinbaseProductTrade) UnmarshalJSON(d []byte) error {
	data := make(umap)
	if err := json.Unmarshal(d, &data); err != nil {
		return err
	}

	var err error

	data.unmarshal("trade_id", func(v interface{}) error {
		trade.TradeID = int(v.(float64))
		return nil
	})

	err = data.unmarshal("price", func(v interface{}) error {
		trade.Price, err = strconv.ParseFloat(v.(string), 64)
		return err
	})
	if err != nil {
		return err
	}

	err = data.unmarshal("size", func(v interface{}) error {
		trade.Size, err = strconv.ParseFloat(v.(string), 64)
		return err
	})
	if err != nil {
		return err
	}

	err = data.unmarshal("side", func(v interface{}) error {
		trade.Side = v.(string)
		return nil
	})
	if err != nil {
		return err
	}

	err = data.unmarshal("time", func(v interface{}) error {
		layOut := "2006-01-02T15:04:05.999999999Z07:00"
		trade.Time, err = time.Parse(layOut, v.(string))
		trade.Time = trade.Time.UTC()
		return err
	})
	if err != nil {
		return err
	}

	return nil
}
