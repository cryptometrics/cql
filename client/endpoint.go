package client

// Endpoint is an interface to build an endpiont from a list of string args
type Endpoint interface {
	// Get takes a list of strings as arguments to build an endpoint
	Get(...*string) string
}
