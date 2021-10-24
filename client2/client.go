package client2

import "net/http"

// C is the client interface
//go:generate mockgen -source client.go -destination ../client2/client_mock.go -package client2
type C interface {
	Identifier() string
	Connect() error
	Do(Request) (*http.Response, error)
}
