package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"cql/graph/generated"
	"cql/model"
)

func (r *coinbaseDepositAddressResolver) Warnings(ctx context.Context, obj *model.CoinbaseDepositAddress) ([]*model.CoinbaseDepositAddressWarning, error) {
	return obj.Warnings.UntypedSlice().([]*model.CoinbaseDepositAddressWarning), nil
}

func (r *coinbaseDepositAddressResolver) ExchagneDepositAddress(ctx context.Context, obj *model.CoinbaseDepositAddress) (*bool, error) {
	return &obj.ExchangeDepositAddress, nil
}

// CoinbaseDepositAddress returns generated.CoinbaseDepositAddressResolver implementation.
func (r *Resolver) CoinbaseDepositAddress() generated.CoinbaseDepositAddressResolver {
	return &coinbaseDepositAddressResolver{r}
}

type coinbaseDepositAddressResolver struct{ *Resolver }
