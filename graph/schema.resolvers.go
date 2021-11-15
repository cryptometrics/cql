package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"cql/graph/generated"
	"cql/model"
	"cql/protomodel"
	"fmt"
)

func (r *mutationResolver) CoinbaseConvertCurrency(ctx context.Context, from string, to string, amount float64, opts *model.CoinbaseCurrencyConversionOpts) (*model.CoinbaseCurrencyConversion, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CoinbaseGenerateCryptoAddress(ctx context.Context, id string) (*model.CoinbaseDepositAddress, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) MakeCoinbaseAccountDeposit(ctx context.Context, input *model.MakeCoinbaseAccountDepositInput) (*model.CoinbaseDeposit, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateCoinbaseLimitOrder(ctx context.Context, input *model.CoinbaseOrderInput) (*model.CoinbaseOrder, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) MakeCoinbasePaymentMethodDeposit(ctx context.Context, input *model.MakeCoinbasePaymentMethodInput) (*model.CoinbaseDeposit, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CoinbaseAccountLedger(ctx context.Context, id string, opts *model.CoinbaseAccountLedgerOptions) ([]*model.CoinbaseAccountLedger, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CoinbaseAccountHold(ctx context.Context, id string, opts *model.CoinbaseAccountHoldOptions) ([]*model.CoinbaseAccountHold, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CoinbaseAccount(ctx context.Context, id string) (*protomodel.CoinbaseAccount, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CoinbaseAccounts(ctx context.Context, test *string) ([]*protomodel.CoinbaseAccount, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CoinbaseAccountTransfer(ctx context.Context, id string, opts *model.CoinbaseAccountTransferOptions) ([]*model.CoinbaseAccountTransfer, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CoinbaseCurrencyConversion(ctx context.Context, id string, opts *model.CoinbaseCurrencyConversionOpts) (*model.CoinbaseCurrencyConversion, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CoinbaseCurrencies(ctx context.Context) ([]*model.CoinbaseCurrency, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CoinbaseCurrency(ctx context.Context, id string) (*model.CoinbaseCurrency, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CoinbasePaymentMethods(ctx context.Context) ([]*model.CoinbasePaymentMethod, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CoinbaseTransfers(ctx context.Context) ([]*model.CoinbaseTransfer, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CoinbaseTransfer(ctx context.Context, id string) (*model.CoinbaseTransfer, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CoinbaseWallets(ctx context.Context, filler *string) ([]*model.CoinbaseWallet, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ClinbaseClientOrder(ctx context.Context, clientOid string) (*model.CoinbaseOrder, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CoinbaseOrder(ctx context.Context, id string) (*model.CoinbaseOrder, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CoinbaseProducts(ctx context.Context, test *string) ([]*model.CoinbaseProduct, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CoinbaseProduct(ctx context.Context, id string) (*model.CoinbaseProduct, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CoinbaseProductDailyStats(ctx context.Context, id string) (*model.CoinbaseProductDailyStats, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CoinbaseProductHistoricalRate(ctx context.Context, id string, start string, end string, granularity int) ([]*model.CoinbaseProductHistoricalRate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CoinbaseProductOrderBook(ctx context.Context, id string, level string) (*model.CoinbaseProductOrderBook, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CoinbaseProductTicker(ctx context.Context, id string) (*model.CoinbaseProductTicker, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CoinbaseProductTrade(ctx context.Context, id string) ([]*model.CoinbaseProductTrade, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CoinbaseTime(ctx context.Context, test *string) (*model.CoinbaseTime, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
