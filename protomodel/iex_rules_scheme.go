package protomodel

// * This is a generated file, do not edit

// IEXRulesScheme is one of the latest schemes for data points, notification// types, and operators used to construct rules.
type IexRulesScheme struct {
	IsLookup  bool    `json:"isLookup"`
	Label     string  `json:"label"`
	Scope     string  `json:"scope"`
	Type      string  `json:"type"`
	Value     string  `json:"value"`
	Weight    float64 `json:"weight"`
	WeightKey string  `json:"weightKey"`
}
