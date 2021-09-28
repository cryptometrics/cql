package model

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
	data, err := newUmap(d)
	if err != nil {
		return err
	}

	if err := data.unmarshalFloatFromString("open", &stats.Open); err != nil {
		return err
	}

	if err := data.unmarshalFloatFromString("high", &stats.High); err != nil {
		return err
	}

	if err := data.unmarshalFloatFromString("low", &stats.Low); err != nil {
		return err
	}

	if err := data.unmarshalFloatFromString("volume", &stats.Volume); err != nil {
		return err
	}

	if err := data.unmarshalFloatFromString("last", &stats.Last); err != nil {
		return err
	}

	err = data.unmarshalFloatFromString("volume_30day", &stats.Volume30Day)
	if err != nil {
		return err
	}

	return nil
}
