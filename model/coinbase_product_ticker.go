package model

import (
	"time"
)

// CoinbaseProductTicket encapsulates snapshot information about the last trade
// (tick), best bid/ask and 24h volume.
//
// Usage: REAL-TIME UPDATES
// Polling is discouraged in favor of connecting via the websocket stream and
// listening for match messages.
type CoinbaseProductTicker struct {
	TradeID int       `json:"trade_id"`
	Price   float64   `json:"price"`
	Size    float64   `json:"size"`
	Bid     float64   `json:"bid"`
	Ask     float64   `json:"ask"`
	Volume  float64   `json:"volume"`
	Time    time.Time `json:"time"`
}

// UnmarshalJSON is an override required to parst strings from coinbases api
// into floats, specifically min_size and max_precision
func (ticker *CoinbaseProductTicker) UnmarshalJSON(d []byte) error {
	data, err := newUmap(d)
	if err != nil {
		return err
	}

	if err := data.unmarshalInt("trade_id", &ticker.TradeID); err != nil {
		return err
	}

	if err := data.unmarshalFloatFromString("price", &ticker.Price); err != nil {
		return err
	}

	if err := data.unmarshalFloatFromString("size", &ticker.Size); err != nil {
		return err
	}

	if err := data.unmarshalFloatFromString("bid", &ticker.Bid); err != nil {
		return err
	}

	if err := data.unmarshalFloatFromString("ask", &ticker.Ask); err != nil {
		return err
	}

	err = data.unmarshalFloatFromString("volume", &ticker.Volume)
	if err != nil {
		return err
	}

	err = data.unmarshalTime(time.RFC3339Nano, "time", &ticker.Time)
	if err != nil {
		return err
	}

	return nil
}
