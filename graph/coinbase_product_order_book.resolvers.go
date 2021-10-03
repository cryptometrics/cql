package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"cql/graph/generated"
	"cql/model"
	"fmt"
)

func (r *coinbaseProductOrderBookResolver) Bids(ctx context.Context, obj *model.CoinbaseProductOrderBook) ([]*model.CoinbaseProductOrderBookBidAsk, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *coinbaseProductOrderBookResolver) Asks(ctx context.Context, obj *model.CoinbaseProductOrderBook) ([]*model.CoinbaseProductOrderBookBidAsk, error) {
	panic(fmt.Errorf("not implemented"))
}

// CoinbaseProductOrderBook returns generated.CoinbaseProductOrderBookResolver implementation.
func (r *Resolver) CoinbaseProductOrderBook() generated.CoinbaseProductOrderBookResolver {
	return &coinbaseProductOrderBookResolver{r}
}

type coinbaseProductOrderBookResolver struct{ *Resolver }
