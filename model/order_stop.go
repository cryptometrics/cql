package model

// OrderStop is either loss or entry
type OrderStop int

const (
	_ OrderStop = iota
	OrderStopLoss
	OrderStopEntry
)

// String returns the stop as a string literal
func (cot OrderStop) String() string {
	return map[OrderStop]string{
		OrderStopLoss:  "loss",
		OrderStopEntry: "entry",
	}[cot]
}

// GetOrderStop takes a string literal and returns the enumerated stop
func GetOrderStop(str string) OrderStop {
	return map[string]OrderStop{
		"loss":  OrderStopLoss,
		"entry": OrderStopEntry,
	}[str]
}
