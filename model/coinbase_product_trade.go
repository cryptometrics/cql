package model

import (
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
	data, err := newUmap(d)
	if err != nil {
		return err
	}

	data.unmarshalInt("trade_id", &trade.TradeID)
	data.unmarshalString("side", &trade.Side)

	if err := data.unmarshalFloatFromString("price", &trade.Price); err != nil {
		return err
	}

	if err := data.unmarshalFloatFromString("size", &trade.Size); err != nil {
		return err
	}

	err = data.unmarshalTime(time.RFC3339Nano, "time", &trade.Time)
	if err != nil {
		return err
	}

	return nil
}
