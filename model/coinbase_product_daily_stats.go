package model

import (
	"encoding/json"
	"strconv"
)

// CoinbaseProductDailyStats encapsulates 24 hr stats for the product
type CoinbaseProductDailyStats struct {
	// Open is in quote currency units; product={base}-{quote}
	Open float64 `json:"open"`

	// High is in quote currency units; product={base}-{quote}
	High float64 `json:"high"`
	Low  float64 `json:"low"`

	//Volume is in base currency units; product={base}-{quote}
	Volume      float64 `json:"volume"`
	Last        float64 `json:"last"`
	Volume30Day float64 `json:"volume_30day"`
}

// UnmarshalJSON is an override required to parst strings from coinbases api
// into floats, specifically min_size and max_precision
func (stats *CoinbaseProductDailyStats) UnmarshalJSON(d []byte) error {
	data := make(umap)
	if err := json.Unmarshal(d, &data); err != nil {
		return err
	}

	var err error

	err = data.unmarshal("open", func(v interface{}) error {
		stats.Open, err = strconv.ParseFloat(v.(string), 64)
		return err
	})
	if err != nil {
		return err
	}

	err = data.unmarshal("high", func(v interface{}) error {
		stats.High, err = strconv.ParseFloat(v.(string), 64)
		return err
	})
	if err != nil {
		return err
	}

	err = data.unmarshal("low", func(v interface{}) error {
		stats.Low, err = strconv.ParseFloat(v.(string), 64)
		return err
	})
	if err != nil {
		return err
	}

	err = data.unmarshal("volume", func(v interface{}) error {
		stats.Volume, err = strconv.ParseFloat(v.(string), 64)
		return err
	})
	if err != nil {
		return err
	}

	err = data.unmarshal("last", func(v interface{}) error {
		stats.Last, err = strconv.ParseFloat(v.(string), 64)
		return err
	})
	if err != nil {
		return err
	}

	err = data.unmarshal("volume_30day", func(v interface{}) error {
		stats.Volume30Day, err = strconv.ParseFloat(v.(string), 64)
		return err
	})
	if err != nil {
		return err
	}

	return nil
}
