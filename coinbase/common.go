package coinbase

import (
	"cql/client"
	"encoding/json"
	"net/http"
)

// fetch uses a client generator and an endpiont to fetch data from the client
func fetch(gen client.Generator, endpoint Endpoint, endpointArgs ...string) (*http.Response, error) {
	cb, err := gen(client.CoinbasePro)
	if err != nil {
		return nil, err
	}
	cb.Connect()
	res, err := cb.Get(endpoint.Get(endpointArgs...))
	if err != nil {
		return nil, err
	}
	return res, nil
}

// decode will fetch the data and then try to decode it into v, which should be
// the pointer to a struct
func decode(gen client.Generator, v interface{}, endpoint Endpoint, endpointArgs ...string) error {
	res, err := fetch(gen, endpoint, endpointArgs...)
	if err != nil {
		return err
	}
	// b, err := io.ReadAll(res.Body)
	// // b, err := ioutil.ReadAll(resp.Body)  Go.1.15 and earlier
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// fmt.Println(string(b))
	decoder := json.NewDecoder(res.Body)
	return decoder.Decode(v)
}
