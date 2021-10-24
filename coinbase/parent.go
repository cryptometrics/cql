package coinbase

import "cql/client2"

type parent struct {
	conn client2.Connector
}

func construct(p *parent, conn client2.Connector) {
	p.conn = conn
}

func (parent *parent) get(ep Endpoint) *client2.Request {
	return client2.New(parent.conn, client2.GET, ep)
}

func (parent *parent) post(ep Endpoint) *client2.Request {
	return client2.New(parent.conn, client2.POST, ep)
}
