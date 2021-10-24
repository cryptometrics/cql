package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"cql/graph/generated"
	"cql/model"
	"fmt"
)

func (r *coinbaseDepositAddressResolver) Warnings(ctx context.Context, obj *model.CoinbaseDepositAddress) ([]*model.CoinbaseDepositAddressWarning, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *coinbaseDepositAddressResolver) ExchagneDepositAddress(ctx context.Context, obj *model.CoinbaseDepositAddress) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

// CoinbaseDepositAddress returns generated.CoinbaseDepositAddressResolver implementation.
func (r *Resolver) CoinbaseDepositAddress() generated.CoinbaseDepositAddressResolver {
	return &coinbaseDepositAddressResolver{r}
}

type coinbaseDepositAddressResolver struct{ *Resolver }
