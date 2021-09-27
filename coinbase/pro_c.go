package coinbase

import (
	"cql/client"
	"cql/env"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

// proC is the coinbase pro client
type proC struct {
	client http.Client
}

// generateSig generates the coinbase base64-encoded signature required to make
// requests.  In particular, the CB-ACCESS-SIGN header is generated by creating
// a sha256 HMAC using the base64-decoded secret key on the prehash string
// timestamp + method + requestPath + body (where + represents string
// concatenation) and base64-encode the output. The timestamp value is the same
// as the CB-ACCESS-TIMESTAMP header.
func (cb *proC) generateSig(secret, message string) (string, error) {
	key, err := base64.StdEncoding.DecodeString(secret)
	if err != nil {
		return "", err
	}

	signature := hmac.New(sha256.New, key)
	_, err = signature.Write([]byte(message))
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(signature.Sum(nil)), nil
}

// generageMsg makes the message to be signed
func (cb *proC) generageMsg(m client.Method, endpoint, data, timestamp string) string {
	return fmt.Sprintf("%s%s%s%s", timestamp, m.String(), endpoint, data)
}

// setHeaders sets the headers for a coinbase api request, in particular:
//
// - CB-ACCESS-KEY The api key as a string.
// - CB-ACCESS-SIGN The base64-encoded signature (see Signing a Message).
// - CB-ACCESS-TIMESTAMP A timestamp for your request.
// - CB-ACCESS-PASSPHRASE The passphrase you specified when creating the API key.
func (cb *proC) setHeaders(req *http.Request, m client.Method, endpoint, data string) (e error) {
	// TODO depricate getting key/passphrase/secret with secret keeper
	var (
		key        = env.CoinbaseProAccessKey.Get()
		passphrase = env.CoinbaseProAccessPassphrase.Get()
		secret     = env.CoinbaseProSecret.Get()
		timestamp  = strconv.FormatInt(time.Now().Unix(), 10)
		msg        = cb.generageMsg(m, endpoint, data, timestamp)
	)
	var sig string
	sig, e = cb.generateSig(secret, msg)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "Go coinbase Pro Client 1.0")
	req.Header.Add("CB-ACCESS-KEY", key)
	req.Header.Add("CB-ACCESS-PASSPHRASE", passphrase)
	req.Header.Add("CB-ACCESS-TIMESTAMP", timestamp)
	req.Header.Add("CB-ACCESS-SIGN", sig)
	return
}

// request makes an http request to the coinbase api, given a method and an
// endpoint.
//
// TODO make data-compatible for non-get requests
func (cb *proC) request(m client.Method, endpoint string) (*http.Response, error) {
	fullURL := env.CoinbaseProURL.Get() + endpoint
	req, err := http.NewRequest(m.String(), fullURL, nil)
	if err != nil {
		return nil, err
	}
	if err := cb.setHeaders(req, m, endpoint, ""); err != nil {
		return nil, err
	}
	return cb.client.Do(req)
}

// Connect creats a new client instance
func (cb *proC) Connect() error {
	cb.client = http.Client{}
	return nil
}

// Identifier identifies requests
func (cb *proC) Identifier() string {
	return "Coinbase Pro"
}

// Get makes and http GET request, given a an endpoint
func (cb *proC) Get(endpoint string) (*http.Response, error) {
	return cb.request(client.GET, endpoint)
}

// newClient returns a new client interface.  This method is what we call a
// "connector"
func newClient() (client.C, error) {
	return &proC{}, nil
}
