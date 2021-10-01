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

type DecodeInput struct {
	Method       Method
	Endpoint     Endpoint
	EndpointArgs EndpointArgs
	Body         Body
}

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

// validateResponse is a switch condition that parses an error response
func validateResponse(res *http.Response) (err error) {
	switch res.StatusCode {
	case http.StatusBadRequest:
		err = parseErrorMessage(res, http.StatusBadRequest)
	case http.StatusUnauthorized:
		err = parseErrorMessage(res, http.StatusUnauthorized)
	case http.StatusInternalServerError:
		err = parseErrorMessage(res, http.StatusInternalServerError)
	}
	return
}

// fetch uses a client connector and an endpiont to fetch data from the client
func (conn Connector) fetch(input DecodeInput) (*http.Response, error) {
	c, err := conn()
	if err != nil {
		return nil, err
	}
	c.Connect()
	input.Method.LogInfo(c, input.Endpoint, input.EndpointArgs, input.Body)

	var res *http.Response
	switch input.Method {
	case GET:
		res, err = c.Get(input.Endpoint.Get(input.EndpointArgs))
	case POST:
		res, err = c.Post(input.Endpoint.Get(input.EndpointArgs), input.Body.Buf)
	}
	if err != nil {
		return nil, err
	}

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
func (conn Connector) Decode(input DecodeInput, v interface{}) error {
	res, err := conn.fetch(input)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// fmt.Println(res.StatusCode, res.Status)
	// printBody(res)

	if err := validateResponse(res); err != nil {
		return err
	}

	return json.NewDecoder(res.Body).Decode(v)
}
