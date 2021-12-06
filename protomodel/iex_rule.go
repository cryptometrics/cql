package protomodel

import "cryptometrics/cql/serial"

// * This is a generated file, do not edit

// Rule to evaluate thousands of data points per second and build event-driven,
// automated alerts using Rules Engine. You can access Rules Engine through the
// IEX Console or through our API using the guidelines below.
type IexRule struct {
	// Label of the lookup
	Label string `json:"label"`

	// The actual lookup value that should be used in the right condition
	Value   string `json:"value"`
	Formula string `json:"formula"`
	Scope   string `json:"scope"`
	Type    string `json:"type"`
}

// UnmarshalJSON will deserialize bytes into a IexRule model
func (iexRule *IexRule) UnmarshalJSON(d []byte) error {
	const (
		formulaJsonTag = "formula"
		labelJsonTag   = "label"
		scopeJsonTag   = "scope"
		typeJsonTag    = "type"
		valueJsonTag   = "value"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	data.UnmarshalString(formulaJsonTag, &iexRule.Formula)
	data.UnmarshalString(labelJsonTag, &iexRule.Label)
	data.UnmarshalString(scopeJsonTag, &iexRule.Scope)
	data.UnmarshalString(typeJsonTag, &iexRule.Type)
	data.UnmarshalString(valueJsonTag, &iexRule.Value)
	return nil
}
