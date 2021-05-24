package client

import "net/http"

// C is the client interface
type C interface {
	Connect() error
	Get(endpoint string) (*http.Response, error)
}

type Type int

const (
	COINBASE Type = iota
)

// New returns a new client interface, given a type
func New(t Type) (C, error) {
	c := map[Type]C{
		COINBASE: &coinbase{},
	}[t]
	return c, nil
}
