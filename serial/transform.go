package serial

import (
	"cql/scalar"
	"time"
)

// Transform interfaces data that can be deserialized/unmarshalled by type.
type Transform interface {
	// unmarshal is a wrapper for setting data in model unmarhallers
	Unmarshal(key string, fn func(interface{}) error) error
	UnmarshalBool(name string, v *bool) error
	UnmarshalEntryType(name string, v *scalar.EntryType) error
	UnmarshalFloatFromString(name string, v *float64) error
	UnmarshalInt(name string, v *int) error
	UnmarshalLocation(name string, v *time.Location) error
	UnmarshalOrderSide(name string, v *scalar.OrderSide) error
	UnmarshalOrderSTP(name string, v *scalar.OrderSTP) error
	UnmarshalOrderTimeInForce(name string, v *scalar.TimeInForce) error
	UnmarshalOrderType(name string, v *scalar.OrderType) error
	UnmarshalStringSlice(name string, v *[]string) error
	UnmarshalString(name string, v *string) error
	UnmarshalStructSlice(name string, v _jsonStructSlice, template interface{}) error
	UnmarshalStruct(name string, v interface{}) error
	UnmarshalTime(layout string, name string, v *time.Time) error
}

// NewJSONTransform will return a new json transfrom that will create a map of objects from a byte
// stream of JSON to be deserialized.
func NewJSONTransform(d []byte) (Transform, error) {
	return make_jsonTransform(d)
}
