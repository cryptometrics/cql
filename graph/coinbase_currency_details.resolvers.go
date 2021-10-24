package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"cql/graph/generated"
	"cql/model"
	"fmt"
)

func (r *coinbaseCurrencyDetailsResolver) PushPaymentMethods(ctx context.Context, obj *model.CoinbaseCurrencyDetails) ([]*string, error) {
	panic(fmt.Errorf("not implemented"))
}

// CoinbaseCurrencyDetails returns generated.CoinbaseCurrencyDetailsResolver implementation.
func (r *Resolver) CoinbaseCurrencyDetails() generated.CoinbaseCurrencyDetailsResolver {
	return &coinbaseCurrencyDetailsResolver{r}
}

type coinbaseCurrencyDetailsResolver struct{ *Resolver }
