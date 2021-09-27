package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"cql/coinbase"
	"cql/graph/generated"
	"cql/graph/model"
	model1 "cql/model"
	"fmt"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
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

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
