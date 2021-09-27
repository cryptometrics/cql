package model

// CoinbaseAccountHistoryDetails are the details for account history
type CoinbaseAccountHistoryDetails struct {
	OrderID   string `json:"order_id"`
	TradeID   string `json:"trade_id"`
	ProductID string `json:"product_id"`
}
