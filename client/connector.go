package client

import (
	"cql/model"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// connector is a client connector, such as the New function.  It's broken into
// it's own type for decoding purposes.
type Connector func() (C, error)

// parseErrorMessage takes a response and a status and builder an error message
// to send to the server
func parseErrorMessage(res *http.Response, status int) error {
	msg := model.CoinbaseMessage{}
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&msg); err != nil {
		return err
	}
	msg.Status = res.Status
	msg.StatusCode = http.StatusText(status)
	return fmt.Errorf("%+v", msg)
}

// fetch uses a client connector and an endpiont to fetch data from the client
func (conn Connector) fetch(req *Request) (*http.Response, error) {
	c, err := conn()
	if err != nil {
		return nil, err
	}
	c.Connect()
	req.setClientName(c)
	req.logHTTPRequest()
	req.logHTTPRequestBody()

	var res *http.Response
	res, err = c.Do(*req)
	if err != nil {
		return nil, err
	}

	req.logResponseStatus(res)
	return res, nil
}

func printBody(res *http.Response) {
	b, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(b))
}

// decode will fetch the data and then try to decode it into v, which should be
// the pointer to a struct
func (conn Connector) Decode(req *Request, v interface{}) error {
	req.generateSlug()
	res, err := conn.fetch(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if err := req.validateResponse(res); err != nil {
		return err
	}
	if err := json.NewDecoder(res.Body).Decode(v); err != nil {
		return err
	}
	if cb := req.Callback; cb != nil {
		return cb(v, req)
	}
	return nil
}
