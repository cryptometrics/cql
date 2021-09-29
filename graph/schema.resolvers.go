package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"cql/coinbase"
	"cql/graph/generated"
	model1 "cql/model"
)

func (r *queryResolver) CoinbaseAccountHistory(ctx context.Context, id string, before *int, after *int, startDate *string, endDate *string, limit *int) ([]*model1.CoinbaseAccountHistory, error) {
	private := coinbase.NewPrivate()
	return private.AccountHistory(id, before, after, startDate, endDate, limit)
}

func (r *queryResolver) CoinbaseAccountHold(ctx context.Context, id string, before *int, after *int, limit *int) ([]*model1.CoinbaseAccountHold, error) {
	private := coinbase.NewPrivate()
	return private.AccountHolds(id, before, after, limit)
}

func (r *queryResolver) CoinbaseAccount(ctx context.Context, id string) (*model1.CoinbaseAccount, error) {
	private := coinbase.NewPrivate()
	return private.Account(id)
}

func (r *queryResolver) CoinbaseAccounts(ctx context.Context, test *string) ([]*model1.CoinbaseAccount, error) {
	private := coinbase.NewPrivate()
	return private.Accounts()
}

func (r *queryResolver) CoinbaseCurrencies(ctx context.Context, test *string) ([]*model1.CoinbaseCurrency, error) {
	md := coinbase.NewMarketData()
	return md.Currencies()
}

func (r *queryResolver) CoinbaseCurrency(ctx context.Context, id string) (*model1.CoinbaseCurrency, error) {
	md := coinbase.NewMarketData()
	return md.Currency(id)
}

func (r *queryResolver) CoinbaseProducts(ctx context.Context, test *string) ([]*model1.CoinbaseProduct, error) {
	md := coinbase.NewMarketData()
	return md.Products()
}

func (r *queryResolver) CoinbaseProduct(ctx context.Context, id string) (*model1.CoinbaseProduct, error) {
	md := coinbase.NewMarketData()
	return md.Product(id)
}

func (r *queryResolver) CoinbaseProductDailyStats(ctx context.Context, id string) (*model1.CoinbaseProductDailyStats, error) {
	md := coinbase.NewMarketData()
	return md.ProductDailyStats(id)
}

func (r *queryResolver) CoinbaseProductHistoricalRate(ctx context.Context, id string, start string, end string, granularity int) ([]*model1.CoinbaseProductHistoricalRate, error) {
	md := coinbase.NewMarketData()
	return md.ProductHistoricalRates(id, start, end, granularity)
}

func (r *queryResolver) CoinbaseProductOrderBook(ctx context.Context, id string, level string) (*model1.CoinbaseProductOrderBook, error) {
	md := coinbase.NewMarketData()
	return md.ProductOrderBook(id, level)
}

func (r *queryResolver) CoinbaseProductTicker(ctx context.Context, id string) (*model1.CoinbaseProductTicker, error) {
	md := coinbase.NewMarketData()
	return md.ProductTicker(id)
}

func (r *queryResolver) CoinbaseProductTrade(ctx context.Context, id string) ([]*model1.CoinbaseProductTrade, error) {
	md := coinbase.NewMarketData()
	return md.ProductTrades(id)
}

func (r *queryResolver) CoinbaseTime(ctx context.Context, test *string) (*model1.CoinbaseTime, error) {
	md := coinbase.NewMarketData()
	return md.Time()
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
