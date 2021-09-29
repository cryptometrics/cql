package model

// OrderSide represents "Buy" or "Sell"
type OrderSide int

const (
	_ OrderSide = iota
	OrderSideBuy
	OrderSideSell
)

// String returns the side as a string literal
func (cot OrderSide) String() string {
	return map[OrderSide]string{
		OrderSideBuy:  "buy",
		OrderSideSell: "sell",
	}[cot]
}

// GetOrderSide takes a string literal and returns the enumerated side
func GetOrderSide(str string) OrderSide {
	return map[string]OrderSide{
		"buy":  OrderSideBuy,
		"sell": OrderSideSell,
	}[str]
}
