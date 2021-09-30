package scalar

// OrderTimeInForce (Time in force) policies provide guarantees about the
// lifetime of an order
type TimeInForce string

const (
	// TimeInForceGTC = good till canceled
	// The default is GTC (CB)
	TimeInForceGTC = "GTC"

	// TimeInForceIOC = immediate or cancel
	// IOC and FOK require post_only be False (CB)
	TimeInForceIOC = "IOC"

	// TimeInForceFOK = fill or kill
	// IOC and FOK require post_only be False (CB)
	TimeInForceFOK = "FOK"

	// TimeInForceGTT = good till time
	// requires that cancel_after be set (CB)
	TimeInForceGTT = "GTT"
)
