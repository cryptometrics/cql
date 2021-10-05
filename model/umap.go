package model

import (
	"cql/scalar"
	"encoding/json"
	"strconv"
	"time"
)

// umap is mean to be used as a data type to hold data for unmarshaling
type umap map[string]interface{}

type umapStructSlice interface {
	append(v interface{})
	UntypedSlice() interface{}
}

func newUmap(d []byte) (umap, error) {
	data := make(umap)
	if err := json.Unmarshal(d, &data); err != nil {
		return nil, err
	}
	return data, nil
}

// unmarshal is a wrapper for setting data in model unmarhallers
func (m umap) unmarshal(key string, fn func(interface{}) error) error {
	if v := m[key]; v != nil {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (m umap) unmarshalBool(name string, v *bool) error {
	if val := m[name]; val != nil {
		*v = val.(bool)
	}
	return nil
}

func (m umap) unmarshalEntryType(name string, v *scalar.EntryType) error {
	if val := m[name]; val != nil {
		*v = scalar.EntryType(val.(string))
	}
	return nil
}

func (m umap) unmarshalOrderSide(name string, v *scalar.OrderSide) error {
	if val := m[name]; val != nil {
		*v = scalar.OrderSide(val.(string))
	}
	return nil
}

func (m umap) unmarshalOrderSTP(name string, v *scalar.OrderSTP) error {
	if val := m[name]; val != nil {
		*v = scalar.OrderSTP((val.(string)))
	}
	return nil
}

func (m umap) unmarshalOrderTimeInForce(name string, v *scalar.TimeInForce) error {
	if val := m[name]; val != nil {
		*v = scalar.TimeInForce(val.(string))
	}
	return nil
}

func (m umap) unmarshalOrderType(name string, v *scalar.OrderType) error {
	if val := m[name]; val != nil {
		*v = scalar.OrderType(val.(string))
	}
	return nil
}

func (m umap) unmarshalFloatFromString(name string, v *float64) (err error) {
	if val := m[name]; val != nil {
		*v, err = strconv.ParseFloat(val.(string), 64)
	}
	return err
}

func (m umap) unmarshalInt(name string, v *int) error {
	if val := m[name]; val != nil {
		*v = int(val.(float64))
	}
	return nil
}

func (m umap) unmarshalLocation(name string, v *time.Location) error {
	if val := m[name]; val != nil {
		*v = *time.FixedZone(val.(string), 0)
	}
	return nil
}

func (m umap) unmarshalStringSlice(name string, v *[]string) error {
	if val := m[name]; val != nil {
		for _, ct := range val.([]interface{}) {
			*v = append(*v, ct.(string))
		}
	}
	return nil
}

func (m umap) unmarshalString(name string, v *string) error {
	if val := m[name]; val != nil {
		*v = val.(string)
	}
	return nil
}

func (m umap) unmarshalStructSlice(name string, v umapStructSlice, template interface{}) error {
	if val := m[name]; val != nil {
		for _, ibid := range val.([]interface{}) {
			jsonString, _ := json.Marshal(ibid)
			if err := json.Unmarshal(jsonString, template); err != nil {
				return err
			}
			v.append(template)
		}
	}
	return nil
}

func (m umap) unmarshalStruct(name string, v interface{}) error {
	if val := m[name]; val != nil {
		jsonString, _ := json.Marshal(val)
		if err := json.Unmarshal(jsonString, v); err != nil {
			return err
		}
	}
	return nil
}

func (m umap) unmarshalTime(layout string, name string, v *time.Time) (err error) {
	if val := m[name]; val != nil {
		*v, err = time.Parse(layout, val.(string))
		if err == nil {
			*v = v.UTC()
		}
	}
	return err
}
