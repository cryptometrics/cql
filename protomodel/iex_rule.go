package protomodel

// * This is a generated file, do not edit

// Rule to evaluate thousands of data points per second and build event-driven,// automated alerts using Rules Engine. You can access Rules Engine through the// IEX Console or through our API using the guidelines below.
type IexRule struct {
	// Label of the lookup
	Label string `json:"label"`

	// The actual lookup value that should be used in the right condition
	Value   string `json:"value"`
	Formula string `json:"formula"`
	Scope   string `json:"scope"`
	Type    string `json:"type"`
}
