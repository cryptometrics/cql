package kraken

import (
	"cryptometrics/cql/client"

	"github.com/cryptometrics/cql/model"
)

// MarketData is a structure used to maintain state while querying on kraken
// market data
type MarketData struct {
	client.Parent
	conn client.Connector
}

// NewMarketData will return a new MarketDAta struct to query market data on
func NewMarketData(conn client.Connector) *MarketData {
	marketData := new(MarketData)
	client.ConstructParent(&marketData.Parent, conn)
	return marketData
}

// Get the server's time.
//
// * source: https://docs.kraken.com/rest/#operation/getServerTime
func (marketData *MarketData) ServerTime() (m *model.KrakenServerTime, err error) {
	req := marketData.Get(ServerTimeEndpoint)
	return m, req.Fetch().Assign(&m).JoinMessages()
}

// Get the current system status or trading mode.
//
// * source: https://docs.kraken.com/rest/#operation/getSystemStatus
func (marketData *MarketData) SystemStatus() (m *model.KrakenSystemStatus, err error) {
	req := marketData.Get(SystemStatusEndpoint)
	return m, req.Fetch().Assign(&m).JoinMessages()
}
