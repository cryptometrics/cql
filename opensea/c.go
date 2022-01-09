package opensea

import (
	"bytes"
	"net/http"

	"github.com/cryptometrics/cql/client"

	"github.com/sirupsen/logrus"
)

// C is the opensea client
type C struct {
	client http.Client
}

func (opensea *C) setHeaders(hreq *http.Request, creq client.Request) (e error) {
	// TODO depricate getting key/passphrase/secret with secret keeper
	hreq.Header.Add("accept", "*/*")
	// hreq.Header.Add("content-type", "application/json")
	// hreq.Header.Add("User-Agent", "Go coinbase Pro Client 1.0")
	return
}

func (opensea *C) Do(creq client.Request) (*http.Response, error) {
	uri := "https://api.opensea.io" + creq.EndpointPath()
	logrus.Debug(client.Logf(&creq, `{Client:{URI:%s}}`, uri))
	hreq, err := http.NewRequest(creq.MethodStr(), uri, bytes.NewReader(creq.GetBody().Bytes()))
	if err != nil {
		return nil, err
	}
	if err := opensea.setHeaders(hreq, creq); err != nil {
		return nil, err
	}
	return opensea.client.Do(hreq)
}

// Connect creats a new client instance
func (opensea *C) Connect() error {
	opensea.client = http.Client{}
	return nil
}

// Identifier identifies requests
func (opensea *C) Identifier() string {
	return "Opensea"
}

// newClient returns a new client interface.  This method is what we call a
// "connector"
func newClient() (client.C, error) {
	return &C{}, nil
}

func DefaultClient() (client.C, error) {
	return &C{}, nil
}
