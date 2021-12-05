package protomodel

import "cql/serial"

// * This is a generated file, do not edit

// CoinbaseCurrency is a currency that coinbase knows about.  Not al currencies
// may be currently in use for trading.
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

// UnmarshalJSON will deserialize bytes into a CoinbaseCurrency model
func (coinbaseCurrency *CoinbaseCurrency) UnmarshalJSON(d []byte) error {
	const (
		convertibleToJsonTag = "convertible_to"
		detailsJsonTag       = "details"
		idJsonTag            = "id"
		maxPrecisionJsonTag  = "max_precision"
		messageJsonTag       = "message"
		minSizeJsonTag       = "min_size"
		nameJsonTag          = "name"
		statusJsonTag        = "status"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	coinbaseCurrency.ProtoDetails = CoinbaseCurrencyDetails{}
	if err := data.UnmarshalStruct(detailsJsonTag, &coinbaseCurrency.ProtoDetails); err != nil {
		return err
	}
	data.UnmarshalFloatFromString(maxPrecisionJsonTag, &coinbaseCurrency.MaxPrecision)
	data.UnmarshalFloatFromString(minSizeJsonTag, &coinbaseCurrency.MinSize)
	data.UnmarshalString(idJsonTag, &coinbaseCurrency.Id)
	data.UnmarshalString(messageJsonTag, &coinbaseCurrency.Message)
	data.UnmarshalString(nameJsonTag, &coinbaseCurrency.Name)
	data.UnmarshalString(statusJsonTag, &coinbaseCurrency.Status)
	data.UnmarshalStringSlice(convertibleToJsonTag, &coinbaseCurrency.ConvertibleTo)
	return nil
}
