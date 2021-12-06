package scalar

// OrderStop is either loss or entry
type OrderStop string

const (
	OrderStopLoss  = "loss"
	OrderStopEntry = "entry"
)
