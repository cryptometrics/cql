package model2

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/franela/goblin"
)

// * This is a generated file, do not edit

func TestCoinbaseAccount200_1(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("200 status example from coinbase", func() {
		response := []byte(`{"id":"7fd0abc0-e5ad-4cbb-8d54-f2b3f43364da","currency":"USD","balance":"0.0000000000000000","available":"0","hold":"0.0000000000000000","profile_id":"8058d771-2d88-4f0f-ab6e-299c153d4308","trading_enabled":true}`)
		actual := CoinbaseAccount{}
		g.It("should pass with all struct assertions", func() {
			err := json.Unmarshal(response, &actual)
			if err != nil {
				panic(err)
			}
			fmt.Println(actual)
		})
	})
}
