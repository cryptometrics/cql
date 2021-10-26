package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"cql/coinbase"
	generated1 "cql/graph/generated"
	model1 "cql/model"
)

func (r *mutationResolver) CoinbaseConvertCurrency(ctx context.Context, from string, to string, amount float64, opts *model1.CoinbaseCurrencyConversionOpts) (*model1.CoinbaseCurrencyConversion, error) {
	return coinbase.NewConversion(coinbase.DefaultClient).Make(from, to, amount, opts)
}

func (r *mutationResolver) CoinbaseGenerateCryptoAddress(ctx context.Context, id string) (*model1.CoinbaseDepositAddress, error) {
	return coinbase.NewCoinbaseAccounts(coinbase.DefaultClient).GenerateCryptoAddress(id)
}

func (r *mutationResolver) MakeCoinbaseAccountDeposit(ctx context.Context, input *model1.MakeCoinbaseAccountDepositInput) (*model1.CoinbaseDeposit, error) {
	return coinbase.NewTransfer(coinbase.DefaultClient).MakeCoinbaseAccountDeposit(input)
}

func (r *mutationResolver) CreateCoinbaseLimitOrder(ctx context.Context, input *model1.CoinbaseOrderInput) (*model1.CoinbaseOrder, error) {
	// private := coinbase.NewPrivate()
	return nil, nil // private.CreateLimitOrder(input)
}

func (r *mutationResolver) MakeCoinbasePaymentMethodDeposit(ctx context.Context, input *model1.MakeCoinbasePaymentMethodInput) (*model1.CoinbaseDeposit, error) {
	return coinbase.NewTransfer(coinbase.DefaultClient).MakePaymentMethodDeposit(input)
}

func (r *queryResolver) CoinbaseAccountLedger(ctx context.Context, id string, opts *model1.CoinbaseAccountLedgerOptions) ([]*model1.CoinbaseAccountLedger, error) {
	return coinbase.NewAccounts(coinbase.DefaultClient).Ledger(id, opts)
}

func (r *queryResolver) CoinbaseAccountHold(ctx context.Context, id string, opts *model1.CoinbaseAccountHoldOptions) ([]*model1.CoinbaseAccountHold, error) {
	return coinbase.NewAccounts(coinbase.DefaultClient).Holds(id, opts)
}

func (r *queryResolver) CoinbaseAccount(ctx context.Context, id string) (*model1.CoinbaseAccount, error) {
	return coinbase.NewAccounts(coinbase.DefaultClient).Find(id)
}

func (r *queryResolver) CoinbaseAccounts(ctx context.Context, test *string) ([]*model1.CoinbaseAccount, error) {
	return coinbase.NewAccounts(coinbase.DefaultClient).All()
}

func (r *queryResolver) CoinbaseAccountTransfer(ctx context.Context, id string, opts *model1.CoinbaseAccountTransferOptions) ([]*model1.CoinbaseAccountTransfer, error) {
	return coinbase.NewAccounts(coinbase.DefaultClient).Transfers(id, opts)
}

func (r *queryResolver) CoinbaseCurrencyConversion(ctx context.Context, id string, opts *model1.CoinbaseCurrencyConversionOpts) (*model1.CoinbaseCurrencyConversion, error) {
	return coinbase.NewConversion(coinbase.DefaultClient).Find(id, opts)
}

func (r *queryResolver) CoinbaseCurrencies(ctx context.Context) ([]*model1.CoinbaseCurrency, error) {
	return coinbase.NewCurrency(coinbase.DefaultClient).All()
}

func (r *queryResolver) CoinbaseCurrency(ctx context.Context, id string) (*model1.CoinbaseCurrency, error) {
	return coinbase.NewCurrency(coinbase.DefaultClient).Find(id)
}

func (r *queryResolver) CoinbasePaymentMethods(ctx context.Context) ([]*model1.CoinbasePaymentMethod, error) {
	return coinbase.NewTransfer(coinbase.DefaultClient).PaymentMethods()
}

func (r *queryResolver) CoinbaseTransfers(ctx context.Context) ([]*model1.CoinbaseTransfer, error) {
	return coinbase.NewTransfer(coinbase.DefaultClient).All()
}

func (r *queryResolver) CoinbaseTransfer(ctx context.Context, id string) (*model1.CoinbaseTransfer, error) {
	return coinbase.NewTransfer(coinbase.DefaultClient).Find(id)
}

func (r *queryResolver) CoinbaseWallets(ctx context.Context, filler *string) ([]*model1.CoinbaseWallet, error) {
	return coinbase.NewCoinbaseAccounts(coinbase.DefaultClient).Wallets()
}

func (r *queryResolver) ClinbaseClientOrder(ctx context.Context, clientOid string) (*model1.CoinbaseOrder, error) {
	// private := coinbase.NewPrivate()
	return nil, nil // private.ClientOrder(clientOid)
}

func (r *queryResolver) CoinbaseOrder(ctx context.Context, id string) (*model1.CoinbaseOrder, error) {
	// private := coinbase.NewPrivate()
	return nil, nil // private.Order(id)
}

func (r *queryResolver) CoinbaseProducts(ctx context.Context, test *string) ([]*model1.CoinbaseProduct, error) {
	// md := coinbase.NewMarketData()
	return nil, nil // md.Products()
}

func (r *queryResolver) CoinbaseProduct(ctx context.Context, id string) (*model1.CoinbaseProduct, error) {
	// md := coinbase.NewMarketData()
	return nil, nil // md.Product(id)
}

func (r *queryResolver) CoinbaseProductDailyStats(ctx context.Context, id string) (*model1.CoinbaseProductDailyStats, error) {
	// md := coinbase.NewMarketData()
	return nil, nil // md.ProductDailyStats(id)
}

func (r *queryResolver) CoinbaseProductHistoricalRate(ctx context.Context, id string, start string, end string, granularity int) ([]*model1.CoinbaseProductHistoricalRate, error) {
	// md := coinbase.NewMarketData()
	return nil, nil // md.ProductHistoricalRates(id, start, end, granularity)
}

func (r *queryResolver) CoinbaseProductOrderBook(ctx context.Context, id string, level string) (*model1.CoinbaseProductOrderBook, error) {
	// md := coinbase.NewMarketData()
	return nil, nil // md.ProductOrderBook(id, level)
}

func (r *queryResolver) CoinbaseProductTicker(ctx context.Context, id string) (*model1.CoinbaseProductTicker, error) {
	// md := coinbase.NewMarketData()
	return nil, nil // md.ProductTicker(id)
}

func (r *queryResolver) CoinbaseProductTrade(ctx context.Context, id string) ([]*model1.CoinbaseProductTrade, error) {
	// md := coinbase.NewMarketData()
	return nil, nil // md.ProductTrades(id)
}

func (r *queryResolver) CoinbaseTime(ctx context.Context, test *string) (*model1.CoinbaseTime, error) {
	// md := coinbase.NewMarketData()
	return nil, nil // md.Time()
}

// Mutation returns generated1.MutationResolver implementation.
func (r *Resolver) Mutation() generated1.MutationResolver { return &mutationResolver{r} }

// Query returns generated1.QueryResolver implementation.
func (r *Resolver) Query() generated1.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
