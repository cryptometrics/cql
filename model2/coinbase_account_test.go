package model2

import (
	"encoding/json"
	"testing"

	"github.com/franela/goblin"
)

// ! This is a generated file, do not edit

func TestCoinbaseAccountTradingEnabled(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("200 status example from coinbase", func() {
		response := []byte(`{"id":"7fd0abc0-e5ad-4cbb-8d54-f2b3f43364da","currency":"USD","balance":"0.0000000000000000","available":"0","hold":"0.0000000000000000","profile_id":"8058d771-2d88-4f0f-ab6e-299c153d4308","trading_enabled":true}`)
		expected := CoinbaseAccount{ID: "7fd0abc0-e5ad-4cbb-8d54-f2b3f43364da", Currency: "USD", Balance: 0.0, Available: 0.0, Hold: 0.0, ProfileID: "8058d771-2d88-4f0f-ab6e-299c153d4308", TradingEnabled: true}
		actual := CoinbaseAccount{}
		if err := json.Unmarshal(response, &actual); err != nil {
			panic(err)
		}
		g.It("should pass with all struct assertions", func() {
			g.Assert(actual.ID).Eql(expected.ID)
			g.Assert(actual.Currency).Eql(expected.Currency)
			g.Assert(actual.Balance).Eql(expected.Balance)
			g.Assert(actual.Available).Eql(expected.Available)
			g.Assert(actual.Hold).Eql(expected.Hold)
			g.Assert(actual.ProfileID).Eql(expected.ProfileID)
			g.Assert(actual.TradingEnabled).Eql(expected.TradingEnabled)
		})
	})
}
