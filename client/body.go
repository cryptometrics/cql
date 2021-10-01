package client

import (
	"cql/null"
	"encoding/json"
)

type Body struct {
	Buf         []byte
	Description string
}

func JSONBody(_map map[string]null.Interface) (body Body) {
	template := map[string]interface{}{}
	for str, m := range _map {
		if m.Valid {
			template[str] = m.Value
		}
	}
	bs, _ := json.Marshal(template)
	body.Description = string(bs)
	buf, err := json.Marshal(template)
	if err != nil {
		panic(err)
	}
	body.Buf = buf
	return
}
