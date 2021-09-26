package model

// umap is mean to be used as a data type to hold data for unmarshaling
type umap map[string]interface{}

// unmarshal is a wrapper for setting data in model unmarhallers
func (m umap) unmarshal(key string, fn func(interface{}) error) error {
	if v := m[key]; v != nil {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}
