package iex

import (
	"github.com/cryptometrics/cql/client"

	"github.com/cryptometrics/cql/model"
)

// RulesEngine is a structure used to maintain state while querying on iex
// rules engine data
type RulesEngine struct {
	client.Parent
	conn client.Connector
}

// NewRulesEngine will return a new RulesEngine struct to query rules
func NewRulesEngine(conn client.Connector) *RulesEngine {
	rulesEngine := new(RulesEngine)
	client.ConstructParent(&rulesEngine.Parent, conn)
	return rulesEngine
}

// Rule for automated alerts
//
// * source: https://iexcloud.io/docs/api/#lookup-values
func (rulesEngine *RulesEngine) Rules(value string) (m []*model.IexRule, err error) {
	req := rulesEngine.Get(RulesEndpoint)
	return m, req.PathParam("value", value).Fetch().Assign(&m).JoinMessages()
}

// Pull the latest schema for data points, notification types, and operators
// used to construct rules.
//
// * source: https://iexcloud.io/docs/api/#rules-schema
func (rulesEngine *RulesEngine) RulesSchema() (m *model.IexRulesSchema, err error) {
	req := rulesEngine.Get(RulesSchemaEndpoint)
	return m, req.Fetch().Assign(&m).JoinMessages()
}
