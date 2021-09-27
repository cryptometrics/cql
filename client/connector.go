package client

import (
	"cql/model"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/sirupsen/logrus"
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

// validateResponse is a switch condition that parses an error response
func validateResponse(res *http.Response) (err error) {
	switch res.StatusCode {
	case http.StatusBadRequest:
		err = parseErrorMessage(res, http.StatusBadRequest)
	case http.StatusUnauthorized:
		err = parseErrorMessage(res, http.StatusUnauthorized)
	}
	return nil
}

// fetch uses a client connector and an endpiont to fetch data from the client
func (conn Connector) fetch(endpoint Endpoint, endpointArgs ...*string) (*http.Response, error) {
	c, err := conn()
	if err != nil {
		return nil, err
	}
	c.Connect()
	parsedEndpoint := endpoint.Get(endpointArgs...)
	logrus.Info("/GET ", parsedEndpoint)
	res, err := c.Get(endpoint.Get(endpointArgs...))
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
func (conn Connector) Decode(v interface{}, endpoint Endpoint, endpointArgs ...*string) error {
	res, err := conn.fetch(endpoint, endpointArgs...)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if err := validateResponse(res); err != nil {
		return err
	}

	return json.NewDecoder(res.Body).Decode(v)
}
