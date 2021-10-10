package model

import (
	"cql/serial"
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
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}

	data.UnmarshalInt("trade_id", &trade.TradeID)
	data.UnmarshalString("side", &trade.Side)

	if err := data.UnmarshalFloatFromString("price", &trade.Price); err != nil {
		return err
	}

	if err := data.UnmarshalFloatFromString("size", &trade.Size); err != nil {
		return err
	}

	err = data.UnmarshalTime(time.RFC3339Nano, "time", &trade.Time)
	if err != nil {
		return err
	}

	return nil
}
