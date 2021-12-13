package coinbase

import (
	"github.com/cryptometrics/cql/client"
	"github.com/cryptometrics/cql/model"
)

// Fees is a structure used to maintain state while querying on trading
// Fees linked to coinbase.
type Fees struct {
	client.Parent
	conn client.Connector
}

// NewFees will return a new Fees structure to query on trading Fees
func NewFees(conn client.Connector) *Fees {
	Fees := new(Fees)
	client.ConstructParent(&Fees.Parent, conn)
	return Fees
}

// All will get fees rates and 30 days trailing volume.
//
// This request will return your current maker & taker fee rates, as well as
// your 30-day trailing volume. Quoted rates are subject to change. More
// information on fees can found on our support page.
//
// * source: https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_getfees
func (fees *Fees) All() (m *model.CoinbaseFees, err error) {
	req := fees.Get(FeesEndpoint)
	return m, req.Fetch().Assign(&m).JoinMessages()
}
