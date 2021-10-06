package client

import (
	"encoding/json"
	"net/http"
)

// connector is a client connector, such as the New function.  It's broken into
// it's own type for decoding purposes.
type Connector func() (C, error)

// fetch uses a client connector and an endpiont to fetch data from the client
func (conn Connector) fetch(req *Request) (*http.Response, error) {
	c, err := conn()
	if err != nil {
		return nil, err
	}
	c.Connect()
	req.setClientName(c)
	req.logHTTPRequest()
	req.logHTTPRequestBody()

	var res *http.Response
	res, err = c.Do(*req)
	if err != nil {
		return nil, err
	}

	req.logResponseStatus(res)
	return res, nil
}

// decode will fetch the data and then try to decode it into v, which should be
// the pointer to a struct
// ! deprecated
func (conn Connector) Decode(req *Request, v interface{}) error {
	req.generateSlug()
	res, err := conn.fetch(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if err := req.validateResponse(res); err != nil {
		return err
	}
	if err := json.NewDecoder(res.Body).Decode(v); err != nil {
		return err
	}
	if cb := req.Callback; cb != nil {
		return cb(v, req)
	}
	return nil
}

// decode will fetch the data and then try to decode it into v, which should be
// the pointer to a struct
func (conn Connector) Fetch(req *Request) *FetchResponse {
	fetchResponse := newFetchResponse()

	// Generate the slug for identifying requests in async logging (if that ever
	// happens)
	req.generateSlug()

	// Then fetch the request from the API
	res, err := conn.fetch(req)
	if err != nil {
		return fetchResponse.setError(err)
	}

	// Validate the response, ensuring that there are no error codes or suspicious
	// messages
	if err := req.validateResponse(res); err != nil {
		return fetchResponse.setError(err)
	}

	// If everything works out, set the body and the request on the fetch response
	// to be used as needed.
	return fetchResponse.setBody(res.Body).setReq(req)
}
