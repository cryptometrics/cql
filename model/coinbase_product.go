package model

import (
	"encoding/json"
	"strconv"
)

// CoinbaseProduct returns the market data for a specific currency pair
//
// Only a maximum of one of trading_disabled, cancel_only, post_only,
// limit_only can be true at once. If none are true, the product is
// trading normally.
type CoinbaseProduct struct {
	ID            string `json:"id"`
	DisplayName   string `json:"display_name"`
	BaseCurrency  string `json:"base_currency"`
	QuoteCurrency string `json:"quote_currency"`

	// BaseIncrement field specifies the minimum increment for the base_currency
	BaseIncrement float64 `json:"base_increment"`

	// QuoteIncrement field specifies the min order price as well as the
	// price increment.  The order price must be a multiple of this increment
	// (i.e. if the increment is 0.01, order prices of 0.001 or 0.021 would be
	// rejected)
	QuoteIncrement float64 `json:"quote_increment"`

	// BaseMinSize defines the min order size
	BaseMinSize float64 `json:"base_min_size"`

	// BaseMaxSize defines the max order size
	BaseMaxSize float64 `json:"base_max_size"`

	// MinMarketFunds defines the min funds allowed in a market order
	MinMarketFunds float64 `json:"min_market_funds"`

	// MaxMarketFunds defines the max funds allowed in a market order
	MaxMarketFunds float64 `json:"max_market_funds"`

	Status string `json:"status"`

	// StatusMessage provides any extra information regarding the status if
	// available
	StatusMessage string `json:"status_message"`

	// CancelOnly indicates whether this product only accepts cancel requests for
	// orders.
	CancelOnly bool `json:"cancel_only"`

	// LimitOnly indicates whether this product only accepts limit orders
	LimitOnly bool `json:"limit_only"`

	// PostOnly indicates whether only maker orders can be placed. No orders will
	// be matched when post_only mode is active
	PostOnly bool `json:"post_only"`

	// TradingDisabled indicates whether trading is currently restricted on this
	// product, this includes whether both new orders and order cancelations are
	// restricted.
	TradingDisabled bool `json:"trading_disabled"`

	// FXStablecoin indicates whether the currency pair is a Stable Pair
	FXStablecoin bool `json:"fx_stablecoin"`
}

// UnmarshalJSON is an override required to parst strings from coinbases api
// into floats
func (product *CoinbaseProduct) UnmarshalJSON(d []byte) error {
	m := make(map[string]interface{})
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

	setter("id", func(v interface{}) error {
		product.ID = v.(string)
		return nil
	})

	setter("display_name", func(v interface{}) error {
		product.DisplayName = v.(string)
		return nil
	})

	setter("base_currency", func(v interface{}) error {
		product.BaseCurrency = v.(string)
		return nil
	})

	setter("quote_currency", func(v interface{}) error {
		product.QuoteCurrency = v.(string)
		return nil
	})

	err = setter("base_increment", func(v interface{}) error {
		strFloat := v.(string)
		product.BaseIncrement, err = strconv.ParseFloat(strFloat, 64)
		return err
	})
	if err != nil {
		return err
	}

	err = setter("quote_increment", func(v interface{}) error {
		strFloat := v.(string)
		product.QuoteIncrement, err = strconv.ParseFloat(strFloat, 64)
		return err
	})
	if err != nil {
		return err
	}

	err = setter("base_min_size", func(v interface{}) error {
		strFloat := v.(string)
		product.BaseMinSize, err = strconv.ParseFloat(strFloat, 64)
		return err
	})
	if err != nil {
		return err
	}

	err = setter("base_max_size", func(v interface{}) error {
		product.BaseMaxSize, err = strconv.ParseFloat(v.(string), 64)
		return err
	})
	if err != nil {
		return err
	}

	setter("min_market_funds", func(v interface{}) error {
		product.MinMarketFunds, err = strconv.ParseFloat(v.(string), 64)
		return nil
	})
	if err != nil {
		return err
	}

	err = setter("max_market_funds", func(v interface{}) error {
		product.MaxMarketFunds, err = strconv.ParseFloat(v.(string), 64)
		return nil
	})
	if err != nil {
		return err
	}

	setter("status", func(v interface{}) error {
		product.Status = v.(string)
		return nil
	})

	setter("status_message", func(v interface{}) error {
		product.StatusMessage = v.(string)
		return nil
	})

	setter("cancel_only", func(v interface{}) error {
		product.CancelOnly = v.(bool)
		return nil
	})

	setter("limit_only", func(v interface{}) error {
		product.LimitOnly = v.(bool)
		return nil
	})

	setter("post_only", func(v interface{}) error {
		product.PostOnly = v.(bool)
		return nil
	})

	setter("trading_disabled", func(v interface{}) error {
		product.TradingDisabled = v.(bool)
		return nil
	})

	setter("fx_stablecoin", func(v interface{}) error {
		product.FXStablecoin = v.(bool)
		return nil
	})

	return nil
}
