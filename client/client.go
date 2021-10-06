package client

import (
	"cql/model"
	"encoding/json"
	"fmt"
	http "net/http"
)

// C is the client interface
//go:generate mockgen -source client.go -destination ../client/client_mock.go -package client
type C interface {
	Identifier() string
	Connect() error
	Do(Request) (*http.Response, error)
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
