package model

// CoinbaseAccountLedgerDetails are the details for account history
type CoinbaseAccountLedgerDetails struct {
	OrderID   string `json:"order_id"`
	TradeID   string `json:"trade_id"`
	ProductID string `json:"product_id"`
}
