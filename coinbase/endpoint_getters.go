package coinbase

import "cql/client"

type endpointGetters map[Endpoint]func(args client.EndpointArgs) string
