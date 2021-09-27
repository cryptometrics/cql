package coinbase

import (
	"cql/client"
	"cql/model"
	"strconv"
)

type MarketData struct{}

// NewMarketData will return a new object to query the coinbase api for market
// data.
func NewMarketData() *MarketData {
	return &MarketData{}
}

func (md *MarketData) currencies(gen client.Connector) (m []*model.CoinbaseCurrency, err error) {
	err = gen.Decode(&m, CurrenciesEP)
	return
}

func (md *MarketData) currency(gen client.Connector, id string) (m *model.CoinbaseCurrency, err error) {
	err = gen.Decode(&m, CurrencyEP, id)
	return
}

func (md *MarketData) products(gen client.Connector) (m []*model.CoinbaseProduct, err error) {
	err = gen.Decode(&m, ProductsEP)
	return
}

func (md *MarketData) product(gen client.Connector, id string) (m *model.CoinbaseProduct, err error) {
	err = gen.Decode(&m, ProductEP, id)
	return
}

func (md *MarketData) productHistoricalRates(gen client.Connector, id, start, end string, granularity int) (m []*model.CoinbaseProductHistoricalRate, err error) {
	err = gen.Decode(&m, ProductHistoricalRatesEP, id, start, end, strconv.Itoa(granularity))
	return
}

func (md *MarketData) productOrderBook(gen client.Connector, id, level string) (m *model.CoinbaseProductOrderBook, err error) {
	err = gen.Decode(&m, ProductOrderBookEP, id, level)
	return
}

func (md *MarketData) productTicker(gen client.Connector, id string) (m *model.CoinbaseProductTicker, err error) {
	err = gen.Decode(&m, ProductTickerEP, id)
	return
}

func (md *MarketData) productTrades(gen client.Connector, id string) (m []*model.CoinbaseProductTrade, err error) {
	err = gen.Decode(&m, ProductTradesEP, id)
	return
}

// Currencies returns a list of all known currencies on coinbase
func (md *MarketData) Currencies() ([]*model.CoinbaseCurrency, error) {
	return md.currencies(newClient)
}

// Currency returns a single MD currency object for a specified id (e.g. BTC)
func (md *MarketData) Currency(id string) (*model.CoinbaseCurrency, error) {
	return md.currency(newClient, id)
}

// Products returns a list of available currency pairs for trading
func (md *MarketData) Products() ([]*model.CoinbaseProduct, error) {
	return md.products(newClient)
}

// Product returns a single MD product to get market data for a specific
// currency pair (e.g. BTC-USD)
func (md *MarketData) Product(id string) (*model.CoinbaseProduct, error) {
	return md.product(newClient, id)
}

// ProductHistoricalRates list of historic rates for a product. Rates are
// returned in grouped buckets based on requested granularity
func (md *MarketData) ProductHistoricalRates(id, start, end string, granularity int) ([]*model.CoinbaseProductHistoricalRate, error) {
	return md.productHistoricalRates(newClient, id, start, end, granularity)
}

// ProductOrderBook returns a list of open orders for a product. The amount of
// detail shown can be customized with the level parameter
func (md *MarketData) ProductOrderBook(id, level string) (*model.CoinbaseProductOrderBook, error) {
	return md.productOrderBook(newClient, id, level)
}

// ProductTicker returns a Snapshot information about the last trade (tick),
// best bid/ask and 24h volume
func (md *MarketData) ProductTicker(id string) (*model.CoinbaseProductTicker, error) {
	return md.productTicker(newClient, id)
}

// ProductTrade returns a list the latest trades for a product
func (md *MarketData) ProductTrades(id string) ([]*model.CoinbaseProductTrade, error) {
	return md.productTrades(newClient, id)
}
