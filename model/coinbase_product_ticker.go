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
	data := make(umap)
	if err := json.Unmarshal(d, &data); err != nil {
		return err
	}

	var err error

	data.unmarshal("trade_id", func(v interface{}) error {
		ticker.TradeID = int(v.(float64))
		return nil
	})

	err = data.unmarshal("price", func(v interface{}) error {
		ticker.Price, err = strconv.ParseFloat(v.(string), 64)
		return err
	})
	if err != nil {
		return err
	}

	err = data.unmarshal("size", func(v interface{}) error {
		ticker.Size, err = strconv.ParseFloat(v.(string), 64)
		return err
	})
	if err != nil {
		return err
	}

	err = data.unmarshal("bid", func(v interface{}) error {
		ticker.Bid, err = strconv.ParseFloat(v.(string), 64)
		return err
	})
	if err != nil {
		return err
	}

	err = data.unmarshal("ask", func(v interface{}) error {
		ticker.Ask, err = strconv.ParseFloat(v.(string), 64)
		return err
	})
	if err != nil {
		return err
	}

	err = data.unmarshal("volume", func(v interface{}) error {
		ticker.Volume, err = strconv.ParseFloat(v.(string), 64)
		return err
	})
	if err != nil {
		return err
	}

	err = data.unmarshal("time", func(v interface{}) error {
		layOut := "2006-01-02T15:04:05.999999999Z07:00"
		ticker.Time, err = time.Parse(layOut, v.(string))
		ticker.Time = ticker.Time.UTC()
		return err
	})
	if err != nil {
		return err
	}

	return nil
}
