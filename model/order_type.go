package model

type OrderType int

const (
	_ OrderType = iota
	OrderTypeMarket
	OrderTypeLimit
)

func (cot OrderType) String() string {
	return map[OrderType]string{
		OrderTypeMarket: "market",
		OrderTypeLimit:  "limit",
	}[cot]
}

func GetOrderType(str string) OrderType {
	return map[string]OrderType{
		"market": OrderTypeMarket,
		"limit":  OrderTypeLimit,
	}[str]
}
