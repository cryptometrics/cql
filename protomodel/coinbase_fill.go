package protomodel

// * This is a generated file, do not edit

// CoinbasePaymentMethod is a partial or complete match on a specific order.
type CoinbaseFill struct {
	Fee       float64 `json:"fee"`
	Liquidity string  `json:"liquidity"`
	OrderId   string  `json:"order_id"`
	Price     float64 `json:"price"`
	ProductId string  `json:"product_id"`
	ProfileId string  `json:"profile_id"`
	Settled   bool    `json:"settled"`
	Side      string  `json:"side"`
	Size      float64 `json:"size"`
	TradeId   int     `json:"trade_id"`
	UsdVolume float64 `json:"usd_volume"`
	UserId    string  `json:"user_id"`
}
