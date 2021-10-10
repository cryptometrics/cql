package model

import "cql/serial"

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
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}

	if err := data.UnmarshalFloatFromString("open", &stats.Open); err != nil {
		return err
	}

	if err := data.UnmarshalFloatFromString("high", &stats.High); err != nil {
		return err
	}

	if err := data.UnmarshalFloatFromString("low", &stats.Low); err != nil {
		return err
	}

	if err := data.UnmarshalFloatFromString("volume", &stats.Volume); err != nil {
		return err
	}

	if err := data.UnmarshalFloatFromString("last", &stats.Last); err != nil {
		return err
	}

	err = data.UnmarshalFloatFromString("volume_30day", &stats.Volume30Day)
	if err != nil {
		return err
	}

	return nil
}
