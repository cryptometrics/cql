package model

import (
	"encoding/json"
	"time"
)

// CoinbaseProductHistoricalRate are historic rates for a product. Rates are
// returned in grouped buckets based on requested granularity.
//
// Historical rate data may be incomplete. No data is published for intervals
// where there are no ticks.
//
// Historical rates should not be polled frequently. If you need real-time
// information, use the trade and book endpoints along with the websocket feed.
type CoinbaseProductHistoricalRate struct {
	// Time is the bucket start time
	Time time.Time `json:"time"`

	// Low is the lowest price during the bucket interval
	Low float64 `json:"low"`

	// High is the highest price during the bucket interval
	High float64 `json:"high"`

	// Open is the opening price (first trade) in the bucket interval
	Open float64 `json:"open"`

	// Close is the closing price (last trade) in the bucket interval
	Close float64 `json:"close"`

	// Volume is the volume of trading activity during the bucket interval
	Volume float64 `json:"volume"`
}

// UnmarshalJSON is an override required to parst strings from coinbases api
// into floats, specifically min_size and max_precision
func (rate *CoinbaseProductHistoricalRate) UnmarshalJSON(d []byte) error {
	data := []interface{}{}
	if err := json.Unmarshal(d, &data); err != nil {
		return err
	}

	// var err error

	rate.Time = time.Unix(int64(data[0].(float64)), 0).UTC()
	rate.Low = data[1].(float64)
	rate.High = data[2].(float64)
	rate.Open = data[3].(float64)
	rate.Close = data[4].(float64)
	rate.Volume = data[5].(float64)

	return nil
}
