package protomodel

// * This is a generated file, do not edit

// CoinbaseCurrency is a currency that coinbase knows about. Not al currencies// may be currently in use for trading.
type CoinbaseCurrency struct {
	ConvertibleTo []string                `json:"convertible_to"`
	Id            string                  `json:"id"`
	MaxPrecision  float64                 `json:"max_precision"`
	Message       string                  `json:"message"`
	MinSize       float64                 `json:"min_size"`
	Name          string                  `json:"name"`
	ProtoDetails  CoinbaseCurrencyDetails `json:"details"`
	Status        string                  `json:"status"`
}
