package model

import "cql/serial"

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
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}

	data.UnmarshalString("id", &product.ID)
	data.UnmarshalString("display_name", &product.DisplayName)
	data.UnmarshalString("base_currency", &product.BaseCurrency)
	data.UnmarshalString("quote_currency", &product.QuoteCurrency)
	data.UnmarshalString("status", &product.Status)
	data.UnmarshalString("status_message", &product.StatusMessage)
	data.UnmarshalBool("cancel_only", &product.CancelOnly)
	data.UnmarshalBool("limit_only", &product.LimitOnly)
	data.UnmarshalBool("post_only", &product.PostOnly)
	data.UnmarshalBool("trading_disabled", &product.TradingDisabled)
	data.UnmarshalBool("fx_stablecoin", &product.FXStablecoin)

	err = data.UnmarshalFloatFromString("base_increment", &product.BaseIncrement)
	if err != nil {
		return err
	}

	err = data.UnmarshalFloatFromString("quote_increment", &product.QuoteIncrement)
	if err != nil {
		return err
	}

	err = data.UnmarshalFloatFromString("base_min_size", &product.BaseMinSize)
	if err != nil {
		return err
	}

	err = data.UnmarshalFloatFromString("base_max_size", &product.BaseMaxSize)
	if err != nil {
		return err
	}

	err = data.UnmarshalFloatFromString("min_market_funds",
		&product.MinMarketFunds)
	if err != nil {
		return err
	}

	err = data.UnmarshalFloatFromString("max_market_funds",
		&product.MaxMarketFunds)
	if err != nil {
		return err
	}
	return nil
}
