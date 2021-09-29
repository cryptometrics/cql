package model

// OrderSTP is the order's Self trade preservation flag
type OrderSTP int

const (
	_ OrderSTP = iota

	// OrderSTP_DC = decrease and cancel, default
	OrderSTP_DC

	// OrderSTP_CO = cancel oldest
	OrderSTP_CO

	// OrderSTP_CN = cancel newest
	OrderSTP_CN

	// OrderSTP_CB = cancel boath
	OrderSTP_CB
)

// String returns the stp as a string literal
func (cot OrderSTP) String() string {
	return map[OrderSTP]string{
		OrderSTP_DC: "dc",
		OrderSTP_CO: "co",
		OrderSTP_CN: "cn",
		OrderSTP_CB: "cb",
	}[cot]
}

// GetOrderSTP takes a string literal and returns the enumerated stp
func GetOrderSTP(str string) OrderSTP {
	return map[string]OrderSTP{
		"dc": OrderSTP_DC,
		"co": OrderSTP_CO,
		"cn": OrderSTP_CN,
		"cb": OrderSTP_CB,
	}[str]
}
