package protomodel

import (
	"cryptometrics/cql/serial"
	"encoding/json"
)

// * This is a generated file, do not edit

// IexRulesSchema is the latest schema for data points, notification types, and
// operators used to construct rules.
type IexRulesSchema struct {
	ProtoSchema []*IexRulesScheme `json:"schema"`
}

// UnmarshalJSON will deserialize bytes into a IexRulesSchema model
func (iexRulesSchema *IexRulesSchema) UnmarshalJSON(d []byte) error {
	const (
		schemaJsonTag = "schema"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	if v := data.Value(schemaJsonTag); v != nil {
		for _, item := range data.Value(schemaJsonTag).([]interface{}) {
			byt, _ := json.Marshal(item)
			obj := IexRulesScheme{}
			if err := json.Unmarshal(byt, &obj); err != nil {
				return err
			}
			iexRulesSchema.ProtoSchema = append(iexRulesSchema.ProtoSchema, &obj)
		}
	}
	return nil
}
