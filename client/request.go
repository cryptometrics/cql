package client

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

type RequestCallback func(interface{}, *Request) error

type Request struct {
	Method       Method
	Endpoint     Endpoint
	EndpointArgs EndpointArgs
	Body         Body
	Callback     RequestCallback

	// slug is an 8 character randomly generated identifiery for the body, to be
	// used to identify request info in logging.
	slug string

	clientName string
}

func (req *Request) generateSlug() {
	b := make([]byte, 4)
	rand.Read(b)
	req.slug = hex.EncodeToString(b)
}

// logInfo will long that a request was made for a method
func (req Request) logHTTPRequest() {
	logrus.Info(req.Logf("/%s %v", req.Method.String(),
		req.Endpoint.Get(req.EndpointArgs)))
}

func (req Request) logHTTPRequestBody() {
	if body := req.Body.Description; len(body) > 0 {
		logrus.Debug(req.Logf(`{Body:%s}`, req.Body.Description))
	}
}

func (req Request) logResponseStatus(res *http.Response) {
	logrus.Debug(req.Logf(`{Response:{StatusCode:%v,Status:%s}}`,
		res.StatusCode, res.Status))
}

func (req *Request) setClientName(c C) {
	req.clientName = c.Identifier()
}

// validateResponse is a switch condition that parses an error response
func (req *Request) validateResponse(res *http.Response) (err error) {
	switch res.StatusCode {
	case http.StatusBadRequest:
		err = parseErrorMessage(res, http.StatusBadRequest)
	case http.StatusUnauthorized:
		err = parseErrorMessage(res, http.StatusUnauthorized)
	case http.StatusInternalServerError:
		err = parseErrorMessage(res, http.StatusInternalServerError)
	case http.StatusNotFound:
		err = parseErrorMessage(res, http.StatusNotFound)
	}
	if err != nil {
		logrus.Warn(req.Logf(err.Error()))
	}
	return
}

func (req Request) Logf(msg string, args ...interface{}) string {
	return fmt.Sprintf(`%s : %v : %s`, req.clientName,
		req.slug, fmt.Sprintf(msg, args...))
}
