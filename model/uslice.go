package model

import (
	"encoding/json"
	"strconv"
)

type uslice []interface{}

func newUslice(d []byte) (uslice, error) {
	data := []interface{}{}
	if err := json.Unmarshal(d, &data); err != nil {
		return nil, err
	}
	return data, nil
}

func (m uslice) unmarshalFloatFromString(idx int, v *float64) (err error) {
	*v, err = strconv.ParseFloat(m[idx].(string), 64)
	return err
}
