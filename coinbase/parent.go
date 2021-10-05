package coinbase

import "cql/client"

type parent struct {
	conn client.Connector
}

func construct(p *parent, conn client.Connector) {
	p.conn = conn
}
