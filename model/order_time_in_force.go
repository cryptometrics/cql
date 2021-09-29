package model

// OrderTimeInForce (Time in force) policies provide guarantees about the
// lifetime of an order
type OrderTimeInForce int

const (
	_ OrderTimeInForce = iota

	// OrderTimeInForceGTC = good till canceled
	// The default is GTC (CB)
	OrderTimeInForceGTC

	// OrderTimeInForceIOC = immediate or cancel
	// IOC and FOK require post_only be False (CB)
	OrderTimeInForceIOC

	// OrderTimeInForceFOK = fill or kill
	// IOC and FOK require post_only be False (CB)
	OrderTimeInForceFOK

	// OrderTimeInForceGTT = good till time
	// requires that cancel_after be set (CB)
	OrderTimeInForceGTT
)

// String returns the TIF as a string literal
func (cot OrderTimeInForce) String() string {
	return map[OrderTimeInForce]string{
		OrderTimeInForceGTC: "GTC",
		OrderTimeInForceIOC: "IOC",
		OrderTimeInForceFOK: "FOK",
		OrderTimeInForceGTT: "GTT",
	}[cot]
}

// GetOrderTimeInForce takes a string literal and returns the enumerated TIF
func GetOrderTimeInForce(str string) OrderTimeInForce {
	return map[string]OrderTimeInForce{
		"GTC": OrderTimeInForceGTC,
		"IOC": OrderTimeInForceIOC,
		"FOK": OrderTimeInForceFOK,
		"GTT": OrderTimeInForceGTT,
	}[str]
}
