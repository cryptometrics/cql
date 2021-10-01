package client

import (
	"net/http"
)

// C is the client interface
//go:generate mockgen -source client.go -destination ../client/client_mock.go -package client
type C interface {
	Identifier() string
	Connect() error
	Get(endpoint string) (*http.Response, error)
	Post(endpoint string, buf []byte) (*http.Response, error)
}
