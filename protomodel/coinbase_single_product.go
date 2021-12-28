package protomodel

// * This is a generated file, do not edit

// CoinbaseSingleProduct is information on a single product
type CoinbaseSingleProduct struct {
	AuctionMode           bool    `json:"auction_mode"`
	BaseCurrency          string  `json:"base_currency"`
	BaseIncrement         float64 `json:"base_increment"`
	BaseMaxSize           float64 `json:"base_max_size"`
	BaseMinSize           float64 `json:"base_min_size"`
	CancelOnly            bool    `json:"cancel_only"`
	DisplayName           string  `json:"display_name"`
	FxStablecoin          bool    `json:"fx_stablecoin"`
	Id                    string  `json:"id"`
	LimitOnly             bool    `json:"limit_only"`
	MarginEnabled         bool    `json:"margin_enabled"`
	MaxMarketFunds        float64 `json:"max_market_funds"`
	MaxSlippagePercentage float64 `json:"max_slippage_percentage"`
	MinMarketFunds        float64 `json:"min_market_funds"`
	PostOnly              bool    `json:"post_only"`
	QuoteCurrency         string  `json:"quote_currency"`
	QuoteIncrement        float64 `json:"quote_increment"`
	Status                string  `json:"status"`
	StatusMessage         string  `json:"status_message"`
	TradingDisabled       bool    `json:"trading_disabled"`
}
