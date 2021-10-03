package client

import (
	"cql/null"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
)

type Body struct {
	Buf         []byte
	Description string

	// Slug is an 8 character randomly generated identifiery for the body, to be
	// used to identify request info in logging.
	Slug string
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
	(&body).generateSlug()
	return
}

func (body *Body) generateSlug() {
	b := make([]byte, 4)
	rand.Read(b)
	body.Slug = hex.EncodeToString(b)
}
