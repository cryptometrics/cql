package model

import (
	"encoding/json"
	"strconv"
	"time"
)

// umap is mean to be used as a data type to hold data for unmarshaling
type umap map[string]interface{}

type umapStructSlice interface {
	append(v interface{})
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
	*v = m[name].(bool)
	return nil
}

func (m umap) unmarshalFloatFromString(name string, v *float64) (err error) {
	*v, err = strconv.ParseFloat(m[name].(string), 64)
	return err
}

func (m umap) unmarshalInt(name string, v *int) error {
	*v = int(m[name].(float64))
	return nil
}

func (m umap) unmarshalStringSlice(name string, v *[]string) error {
	for _, ct := range m[name].([]interface{}) {
		*v = append(*v, ct.(string))
	}
	return nil
}

func (m umap) unmarshalString(name string, v *string) error {
	*v = m[name].(string)
	return nil
}

func (m umap) unmarshalStructSlice(name string, v umapStructSlice, template interface{}) error {
	for _, ibid := range m[name].([]interface{}) {
		jsonString, _ := json.Marshal(ibid)
		if err := json.Unmarshal(jsonString, template); err != nil {
			return err
		}
		v.append(template)
	}
	return nil
}

func (m umap) unmarshalStruct(name string, v interface{}) error {
	jsonString, _ := json.Marshal(m[name])
	if err := json.Unmarshal(jsonString, v); err != nil {
		return err
	}
	return nil
}

func (m umap) unmarshalTime(name string, v *time.Time) (err error) {
	*v, err = time.Parse(timeLayout, m[name].(string))
	if err == nil {
		*v = v.UTC()
	}
	return err
}
