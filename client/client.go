package client

import "net/http"

// C is the client interface
//go:generate mockgen -source client.go -destination ../client/client_mock.go -package client
type C interface {
	Connect() error
	Get(endpoint string) (*http.Response, error)
}

// New returns a new client interface, given a type
func New(t Kind) (C, error) {
	c := map[Kind]C{
		CoinbasePro: &coinbaseProC{},
	}[t]
	return c, nil
}
