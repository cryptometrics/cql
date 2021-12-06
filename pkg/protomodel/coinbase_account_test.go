package protomodel

import (
	"encoding/json"
	"testing"

	"github.com/franela/goblin"
)

// * This is a generated file, do not edit

func TestCoinbaseAccount200_1(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("200 status example from coinbase", func() {
		response := []byte(`{"id":"7fd0abc0-e5ad-4cbb-8d54-f2b3f43364da","currency":"USD","balance":"0.0000000000000000","available":"0","hold":"0.0000000000000000","profile_id":"8058d771-2d88-4f0f-ab6e-299c153d4308","trading_enabled":true}`)
		expected := CoinbaseAccount{Id: "7fd0abc0-e5ad-4cbb-8d54-f2b3f43364da", Currency: "USD", Balance: 0.0, Available: 0.0, Hold: 0.0, ProfileId: "8058d771-2d88-4f0f-ab6e-299c153d4308", TradingEnabled: true}
		actual := CoinbaseAccount{}
		g.It("should pass with all struct assertions", func() {
			err := json.Unmarshal(response, &actual)
			if err != nil {
				panic(err)
			}
			g.Assert(actual.Id).Eql(expected.Id)
			g.Assert(actual.Currency).Eql(expected.Currency)
			g.Assert(actual.Balance).Eql(expected.Balance)
			g.Assert(actual.Available).Eql(expected.Available)
			g.Assert(actual.Hold).Eql(expected.Hold)
			g.Assert(actual.ProfileId).Eql(expected.ProfileId)
			g.Assert(actual.TradingEnabled).Eql(expected.TradingEnabled)
		})
	})
}
