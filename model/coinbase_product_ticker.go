package model

import (
	"encoding/json"
	"strconv"
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
	m := map[string]interface{}{}
	if err := json.Unmarshal(d, &m); err != nil {
		return err
	}

	setter := func(name string, fn func(interface{}) error) error {
		if v := m[name]; v != nil {
			if err := fn(v); err != nil {
				return err
			}
		}
		return nil
	}

	var err error

	setter("trade_id", func(v interface{}) error {
		ticker.TradeID = int(v.(float64))
		return nil
	})

	err = setter("price", func(v interface{}) error {
		ticker.Price, err = strconv.ParseFloat(v.(string), 64)
		return err
	})
	if err != nil {
		return err
	}

	err = setter("size", func(v interface{}) error {
		ticker.Size, err = strconv.ParseFloat(v.(string), 64)
		return err
	})
	if err != nil {
		return err
	}

	err = setter("bid", func(v interface{}) error {
		ticker.Bid, err = strconv.ParseFloat(v.(string), 64)
		return err
	})
	if err != nil {
		return err
	}

	err = setter("ask", func(v interface{}) error {
		ticker.Ask, err = strconv.ParseFloat(v.(string), 64)
		return err
	})
	if err != nil {
		return err
	}

	err = setter("volume", func(v interface{}) error {
		ticker.Volume, err = strconv.ParseFloat(v.(string), 64)
		return err
	})
	if err != nil {
		return err
	}

	err = setter("time", func(v interface{}) error {
		layOut := "2006-01-02T15:04:05.000000Z"
		ticker.Time, err = time.Parse(layOut, v.(string))
		return err
	})
	if err != nil {
		return err
	}

	return nil
}
