package client

import (
	"encoding/json"
	"io"
)

// FetchResponse encapsulates a reponse from fetching data from an API
type FetchResponse struct {
	Error error
	Body  io.ReadCloser

	req *Request
}

func newFetchResponse() *FetchResponse {
	return new(FetchResponse)
}

func (fetchResponse *FetchResponse) setBody(body io.ReadCloser) *FetchResponse {
	fetchResponse.Body = body
	return fetchResponse
}

func (fetchResponse *FetchResponse) setReq(req *Request) *FetchResponse {
	fetchResponse.req = req
	return fetchResponse
}

func (fetchResponse *FetchResponse) setError(err error) *FetchResponse {
	fetchResponse.Error = err
	return fetchResponse
}

// Assign will assign the body of the response to some pointer-value, probably
// a struct.
func (fetchResponse *FetchResponse) Assign(v interface{}) *FetchResponse {
	// If no error, decode the body, defer closure
	if fetchResponse.Error == nil {
		defer fetchResponse.Body.Close()
		if err := json.NewDecoder(fetchResponse.Body).Decode(v); err != nil {
			fetchResponse.setError(err)
		}
	}

	// If no error and callback exists, run callback
	if fetchResponse.Error == nil {
		if cb := fetchResponse.req.Callback; cb != nil {
			if err := cb(v, fetchResponse.req); err != nil {
				fetchResponse.setError(err)
			}
		}
	}

	return fetchResponse
}
