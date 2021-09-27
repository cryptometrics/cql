package model

import (
	"encoding/json"
	"testing"
	"time"

	. "github.com/franela/goblin"
)

func TestCoinbaseCurrencyUnmarshalJSON(t *testing.T) {
	g := Goblin(t)
	g.Describe("CoinbaseCurrency#UnmarshalJSON", func() {
		test := func(desc string, buf []byte, expected CoinbaseCurrency) {
			g.It(desc, func() {
				v := CoinbaseCurrency{}
				if err := json.Unmarshal(buf, &v); err != nil {
					panic(err)
				}
				g.Assert(v).Equal(expected)
			})
		}
		var buf []byte
		var expected CoinbaseCurrency

		buf = []byte(`{
			"id": "BTC",
			"name": "Bitcoin",
			"min_size": "0.00000001",
			"status": "online",
			"message": "",
			"max_precision": "0.00000001",
			"convertible_to": ["USDC"],
			"details": {
				"type": "crypto",
				"symbol": "",
				"network_confirmations": 6,
				"sort_order": 3,
				"crypto_address_link": "meep",
				"crypto_transaction_link": "moop",
				"push_payment_methods": [
					"crypto"
				],
				"group_types": [
					"btc",
					"crypto"
				]
			},
			"display_name": "",
			"processing_time_seconds": 0,
			"min_withdrawal_amount": 0,
			"max_withdrawal_amount": 1000
		}`)
		details := CoinbaseCurrencyDetails{
			Type:                  "crypto",
			Symbol:                "",
			NetworkConfirmations:  6,
			SortOrder:             3,
			CryptoAddressLink:     "meep",
			CryptoTransactionLink: "moop",
			PushPaymentMethods:    []string{"crypto"},
			GroupTypes:            []string{"btc", "crypto"},
		}
		expected = CoinbaseCurrency{
			ID:                    "BTC",
			Name:                  "Bitcoin",
			MinSize:               0.00000001,
			Status:                "online",
			Message:               "",
			MaxPrecision:          0.00000001,
			ConvertibleTo:         []string{"USDC"},
			Details:               details,
			DisplayName:           "",
			ProcessingTimeSeconds: 0,
			MinWithdrawalAmount:   0,
			MaxWithdrawalAmount:   1000,
		}
		test("all fields", buf, expected)
	})
}

func TestCoinbaseProductHistoricalRateUnmarshalJSON(t *testing.T) {
	g := Goblin(t)
	g.Describe("CoinbaseProductHistoricalRate#UnmarshalJSON", func() {
		test := func(desc string, buf []byte, exp CoinbaseProductHistoricalRate) {
			g.It(desc, func() {
				v := CoinbaseProductHistoricalRate{}
				if err := json.Unmarshal(buf, &v); err != nil {
					panic(err)
				}
				g.Assert(v).Equal(exp)
			})
		}
		var buf []byte
		var expected CoinbaseProductHistoricalRate

		buf = []byte(`[ 1415398768, 0.32, 4.2, 0.35, 4.2, 12.3 ]`)

		expected = CoinbaseProductHistoricalRate{
			Time:   time.Date(2014, 11, 07, 22, 19, 28, 0, time.UTC),
			Low:    0.32,
			High:   4.2,
			Open:   0.35,
			Close:  4.2,
			Volume: 12.3,
		}
		test("level 2, single", buf, expected)
	})
}

func TestCoinbaseProductOrderBookUnmarshalJSON(t *testing.T) {
	g := Goblin(t)
	g.Describe("CoinbaseProductOrderBook#UnmarshalJSON", func() {
		test := func(desc string, buf []byte, expected CoinbaseProductOrderBook) {
			g.It(desc, func() {
				v := CoinbaseProductOrderBook{}
				if err := json.Unmarshal(buf, &v); err != nil {
					panic(err)
				}
				g.Assert(v).Equal(expected)
			})
		}
		var buf []byte
		var expected CoinbaseProductOrderBook

		buf = []byte(`{
			"sequence": 3,
			"bids": [
					[ "295.96", "4.39088265", 2 ]
			],
			"asks": [
					[ "295.97", "25.23542881", 12 ]
			]
		}`)

		expected = CoinbaseProductOrderBook{
			Sequence: 3,
			Bids: []CoinbaseProductOrderBookBidAsk{
				{295.96, 4.39088265, Int(2), nil}},
			Asks: []CoinbaseProductOrderBookBidAsk{
				{295.97, 25.23542881, Int(12), nil}},
		}
		test("level 2, single", buf, expected)

		buf = []byte(`{
			"sequence": 3,
			"bids": [
					[ "295.96","0.05088265","3b0f1225-7f84-490b-a29f-0faef9de823a" ]
			],
			"asks": [
					[ "295.97","5.72036512","da863862-25f4-4868-ac41-005d11ab0a5f" ]
			]
		}`)
		expected = CoinbaseProductOrderBook{
			Sequence: 3,
			Bids: []CoinbaseProductOrderBookBidAsk{{295.96, 0.05088265, nil,
				String("3b0f1225-7f84-490b-a29f-0faef9de823a")}},
			Asks: []CoinbaseProductOrderBookBidAsk{{295.97, 5.72036512, nil,
				String("da863862-25f4-4868-ac41-005d11ab0a5f")}},
		}
		test("level 3, single", buf, expected)

		buf = []byte(`{
			"sequence": 3,
			"bids": [
					[ "295.96", "4.39088265", 2 ],
					[ "295.96", "4.39088265", 2 ]
			],
			"asks": [
					[ "295.97", "25.23542881", 12 ],
					[ "295.97", "25.23542881", 12 ]
			]
		}`)

		expected = CoinbaseProductOrderBook{
			Sequence: 3,
			Bids: []CoinbaseProductOrderBookBidAsk{
				{295.96, 4.39088265, Int(2), nil},
				{295.96, 4.39088265, Int(2), nil},
			},
			Asks: []CoinbaseProductOrderBookBidAsk{
				{295.97, 25.23542881, Int(12), nil},
				{295.97, 25.23542881, Int(12), nil},
			},
		}
		test("level 2, multiple", buf, expected)
	})
}

func TestCoinbaseProductTickerUnmarshalJSON(t *testing.T) {
	g := Goblin(t)
	g.Describe("CoinbaseProductTicker#UnmarshalJSON", func() {
		test := func(desc string, buf []byte, expected CoinbaseProductTicker) {
			g.It(desc, func() {
				v := CoinbaseProductTicker{}
				if err := json.Unmarshal(buf, &v); err != nil {
					panic(err)
				}
				g.Assert(v).Equal(expected)
			})
		}
		var buf []byte
		var expected CoinbaseProductTicker

		buf = []byte(`{
	 		"trade_id": 4729088,
	 		"price": "333.99",
	 		"size": "0.193",
	 		"bid": "333.98",
	 		"ask": "333.99",
	 		"volume": "5957.11914015",
	 		"time": "2015-11-14T20:46:03.511254Z"
	 	}`)
		expected = CoinbaseProductTicker{
			TradeID: 4729088,
			Price:   333.99,
			Size:    0.193,
			Bid:     333.98,
			Ask:     333.99,
			Volume:  5957.11914015,
		}

		var err error
		expected.Time, err = time.Parse("2006-01-02T15:04:05.000000Z", "2015-11-14T20:46:03.511254Z")
		if err != nil {
			panic(err)
		}
		test("all fields", buf, expected)
	})
}

func TestCoinbaseProductTradeUnmarshalJSON(t *testing.T) {
	g := Goblin(t)
	g.Describe("CoinbaseProductTrade#UnmarshalJSON", func() {
		test := func(desc string, buf []byte, expected CoinbaseProductTrade) {
			g.It(desc, func() {
				v := CoinbaseProductTrade{}
				if err := json.Unmarshal(buf, &v); err != nil {
					panic(err)
				}
				g.Assert(v).Equal(expected)
			})
		}
		var buf []byte
		var expected CoinbaseProductTrade

		buf = []byte(`		{
	    "time": "2014-11-07T22:19:28.578544Z",
	    "trade_id": 74,
	    "price": "10.00000000",
	    "size": "0.01000000",
	    "side": "buy"
		}`)

		expected = CoinbaseProductTrade{
			TradeID: 74,
			Price:   10.00000000,
			Size:    0.01000000,
			Side:    "buy",
		}

		var err error
		expected.Time, err = time.Parse("2006-01-02T15:04:05.000000Z", "2014-11-07T22:19:28.578544Z")
		if err != nil {
			panic(err)
		}
		test("all fields", buf, expected)
	})
}

func TestCoinbaseProductUnmarshalJSON(t *testing.T) {
	g := Goblin(t)
	g.Describe("CoinbaseProduct#UnmarshalJSON", func() {
		test := func(desc string, buf []byte, expected CoinbaseProduct) {
			g.It(desc, func() {
				v := CoinbaseProduct{}
				if err := json.Unmarshal(buf, &v); err != nil {
					panic(err)
				}
				g.Assert(v).Equal(expected)
			})
		}
		var buf []byte
		var expected CoinbaseProduct

		buf = []byte(`{
			"id": "BTC-USD",
			"display_name": "BTC/USD",
			"base_currency": "BTC",
			"quote_currency": "USD",
			"base_increment": "0.00000001",
			"quote_increment": "0.01000000",
			"base_min_size": "0.00100000",
			"base_max_size": "280.00000000",
			"min_market_funds": "5",
			"max_market_funds": "1000000",
			"status": "online",
			"status_message": "",
			"cancel_only": false,
			"limit_only": false,
			"post_only": false,
			"trading_disabled": false,
			"fx_stablecoin": false
		}`)
		expected = CoinbaseProduct{
			ID:              "BTC-USD",
			DisplayName:     "BTC/USD",
			BaseCurrency:    "BTC",
			QuoteCurrency:   "USD",
			BaseIncrement:   0.00000001,
			QuoteIncrement:  0.01000000,
			BaseMinSize:     0.00100000,
			BaseMaxSize:     280.00000000,
			MinMarketFunds:  5,
			MaxMarketFunds:  1000000,
			Status:          "online",
			StatusMessage:   "",
			CancelOnly:      false,
			LimitOnly:       false,
			PostOnly:        false,
			TradingDisabled: false,
			FXStablecoin:    false,
		}
		test("all fields", buf, expected)
	})
}
