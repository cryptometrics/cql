package client

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"io"
)

type assignmentCallback func(interface{}, *Client) error

type Client struct {
	assignmentCallback
	body      io.ReadCloser
	connector Connector
	endpoint  Endpoint
	errors    *Errors
	method    Method

	// slug is an 8 character randomly generated identifiery for the body, to be
	// used to identify request info in logging.
	slug string
}

type Assigner struct {
	client *Client
}

func newAssigner(client *Client) *Assigner {
	assigner := new(Assigner)
	assigner.client = client
	return assigner
}

func New(conn Connector, method Method, endpoint Endpoint) *Client {
	client := new(Client)
	client.connector = conn
	client.endpoint = endpoint
	client.method = method
	return client
}

func (assigner *Assigner) decode(v interface{}) {
	if !assigner.client.errors.Any() {
		if err := json.NewDecoder(assigner.client.body).Decode(v); err != nil {
			assigner.client.errors.add(err)
		}
	}
}

// generateSlub will make an 8 character randomly generated identifier for the body, which can be
// used to identify request info in logging.
func (client *Client) generateSlug() {
	b := make([]byte, 4)
	rand.Read(b)
	client.slug = hex.EncodeToString(b)
}

func (assigner *Assigner) runAssignmentCallback(v interface{}) {
	if !assigner.client.errors.Any() {
		if err := assigner.client.assignmentCallback(v, assigner.client); err != nil {
			assigner.client.errors.add(err)
		}
	}
}

// Assign will assign the body of the response to some pointer-value, probably a struct
func (assigner *Assigner) Assign(v interface{}) *Errors {
	defer assigner.client.body.Close()
	assigner.decode(v)
	assigner.runAssignmentCallback(v)
	return assigner.client.errors
}

// AssignmentCallback will set a callback on the client that runs post-assignment
func (client *Client) AssignmentCallback(cb assignmentCallback) *Client {
	client.assignmentCallback = cb
	return client
}

// Fetch will use the client's connector to
func (client *Client) Fetch() *Assigner {
	assigner := newAssigner(client)
	client.generateSlug()
	return assigner
}
