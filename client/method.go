package client

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

// Common HTTP methods.
//
// Unless otherwise noted, these are defined in RFC 7231 section 4.3.
type Method int

const (
	GET Method = iota
)

// String will return the string expectation of the method
func (m Method) String() string {
	return map[Method]string{
		GET: "GET",
	}[m]
}

// LogInfo will long that a request was made for a method
func (m Method) LogInfo(c C, endpoint Endpoint, args EndpointArgs) {
	logrus.Info(fmt.Sprintf("/%s (%s) %s", m.String(), c.Identifier(),
		endpoint.Get(args)))
}
