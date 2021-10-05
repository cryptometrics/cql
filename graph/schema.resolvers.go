package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"cql/coinbase"
	generated1 "cql/graph/generated"
	model1 "cql/model"
)

func (r *mutationResolver) CoinbaseGenerateCryptoAddress(ctx context.Context, id string) (*model1.CoinbaseDepositAddress, error) {
	return coinbase.NewCoinbaseAccounts().GenerateCryptoAddress(id)
}

func (r *mutationResolver) CreateCoinbaseLimitOrder(ctx context.Context, input *model1.CoinbaseOrderInput) (*model1.CoinbaseOrder, error) {
	private := coinbase.NewPrivate()
	return private.CreateLimitOrder(input)
}

func (r *queryResolver) CoinbaseAccountLedger(ctx context.Context, id string, opts *model1.CoinbaseAccountLedgerOptions) ([]*model1.CoinbaseAccountLedger, error) {
	return coinbase.NewAccounts().Ledger(id, opts)
}

func (r *queryResolver) CoinbaseAccountHold(ctx context.Context, id string, opts *model1.CoinbaseAccountHoldOptions) ([]*model1.CoinbaseAccountHold, error) {
	return coinbase.NewAccounts().Holds(id, opts)
}

func (r *queryResolver) CoinbaseAccount(ctx context.Context, id string) (*model1.CoinbaseAccount, error) {
	return coinbase.NewAccounts().Find(id)
}

func (r *queryResolver) CoinbaseAccounts(ctx context.Context, test *string) ([]*model1.CoinbaseAccount, error) {
	return coinbase.NewAccounts().All()
}

func (r *queryResolver) CoinbaseAccountTransfer(ctx context.Context, id string, opts *model1.CoinbaseAccountTransferOptions) ([]*model1.CoinbaseAccountTransfer, error) {
	return coinbase.NewAccounts().Transfers(id, opts)
}

func (r *queryResolver) CoinbaseWallets(ctx context.Context, filler *string) ([]*model1.CoinbaseWallet, error) {
	return coinbase.NewCoinbaseAccounts().Wallets()
}

func (r *queryResolver) ClinbaseClientOrder(ctx context.Context, clientOid string) (*model1.CoinbaseOrder, error) {
	private := coinbase.NewPrivate()
	return private.ClientOrder(clientOid)
}

func (r *queryResolver) CoinbaseCurrencies(ctx context.Context, test *string) ([]*model1.CoinbaseCurrency, error) {
	md := coinbase.NewMarketData()
	return md.Currencies()
}

func (r *queryResolver) CoinbaseCurrency(ctx context.Context, id string) (*model1.CoinbaseCurrency, error) {
	md := coinbase.NewMarketData()
	return md.Currency(id)
}

func (r *queryResolver) CoinbaseOrder(ctx context.Context, id string) (*model1.CoinbaseOrder, error) {
	private := coinbase.NewPrivate()
	return private.Order(id)
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

// Mutation returns generated1.MutationResolver implementation.
func (r *Resolver) Mutation() generated1.MutationResolver { return &mutationResolver{r} }

// Query returns generated1.QueryResolver implementation.
func (r *Resolver) Query() generated1.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
