package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"cql/coinbase"
	"cql/graph/generated"
	"cql/graph/model"
	model1 "cql/model"
	"cql/svc"
	"fmt"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CoinbaseMarketDataCurrencies(ctx context.Context, test *string) ([]*model1.CoinbaseMarketDataCurrency, error) {
	md := &coinbase.MarketData{}
	return md.Currencies()
}

func (r *queryResolver) CoinbaseMarketDataCurrency(ctx context.Context, id string) (*model1.CoinbaseMarketDataCurrency, error) {
	md := &coinbase.MarketData{}
	return md.Currency(id)
}

func (r *queryResolver) CoinbaseBuyPrice(ctx context.Context, currencyPair string) (*model1.CoinbaseBuyPrice, error) {
	p := svc.CoinbaseBuyPrice{}
	return p.Price(currencyPair)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
