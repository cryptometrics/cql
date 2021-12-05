package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"cql/coinbase"
	"cql/graph/generated"
	"cql/iex"
	"cql/kraken"
	"cql/model"
)

func (r *mutationResolver) CoinbaseAccountDeposit(ctx context.Context, opts *model.CoinbaseAccountDepositOptions) (*model.CoinbaseDeposit, error) {
	return coinbase.NewTransfer(coinbase.DefaultClient).MakeCoinbaseAccountDeposit(opts)
}

func (r *mutationResolver) CoinbaseConvertCurrency(ctx context.Context, opts model.CoinbaseConversionsOptions) (*model.CoinbaseCurrencyConversion, error) {
	return coinbase.NewConversion(coinbase.DefaultClient).Make(opts)
}

func (r *mutationResolver) CoinbaseGenerateCryptoAddress(ctx context.Context, walletID string) (*model.CoinbaseCryptoAddress, error) {
	return coinbase.NewCoinbaseAccounts(coinbase.DefaultClient).GenerateCryptoAddress(walletID)
}

func (r *mutationResolver) CoinbasePaymentMethodDeposit(ctx context.Context, opts *model.CoinbasePaymentMethodDepositOptions) (*model.CoinbaseDeposit, error) {
	return coinbase.NewTransfer(coinbase.DefaultClient).MakePaymentMethodDeposit(opts)
}

func (r *queryResolver) CoinbaseAccount(ctx context.Context, accountID string) (*model.CoinbaseAccount, error) {
	return coinbase.NewAccounts(coinbase.DefaultClient).Find(accountID)
}

func (r *queryResolver) CoinbaseAccounts(ctx context.Context) ([]*model.CoinbaseAccount, error) {
	return coinbase.NewAccounts(coinbase.DefaultClient).All()
}

func (r *queryResolver) CoinbaseAccountHolds(ctx context.Context, accountID string, opts *model.CoinbaseAccountHoldsOptions) ([]*model.CoinbaseAccountHold, error) {
	return coinbase.NewAccounts(coinbase.DefaultClient).Holds(accountID, opts)
}

func (r *queryResolver) CoinbaseAccountLedger(ctx context.Context, accountID string, opts *model.CoinbaseAccountLedgerOptions) ([]*model.CoinbaseAccountLedger, error) {
	return coinbase.NewAccounts(coinbase.DefaultClient).Ledger(accountID, opts)
}

func (r *queryResolver) CoinbaseAccountTransfers(ctx context.Context, accountID string, opts *model.CoinbaseAccountTransferOptions) ([]*model.CoinbaseAccountTransfer, error) {
	return coinbase.NewAccounts(coinbase.DefaultClient).Transfers(accountID, opts)
}

func (r *queryResolver) CoinbaseCurrencies(ctx context.Context) ([]*model.CoinbaseCurrency, error) {
	return coinbase.NewCurrency(coinbase.DefaultClient).All()
}

func (r *queryResolver) CoinbaseCurrencyConversion(ctx context.Context, conversionID string, opts *model.CoinbaseConversionOptions) (*model.CoinbaseCurrencyConversion, error) {
	return coinbase.NewConversion(coinbase.DefaultClient).Find(conversionID, opts)
}

func (r *queryResolver) CoinbaseCurrency(ctx context.Context, currentID string) (*model.CoinbaseCurrency, error) {
	return coinbase.NewCurrency(coinbase.DefaultClient).Find(currentID)
}

func (r *queryResolver) CoinbasePaymentMethods(ctx context.Context) ([]*model.CoinbasePaymentMethod, error) {
	return coinbase.NewTransfer(coinbase.DefaultClient).PaymentMethods()
}

func (r *queryResolver) IexRules(ctx context.Context, value string) ([]*model.IexRule, error) {
	return iex.NewRulesEngine(iex.DefaultClient).Rules(value)
}

func (r *queryResolver) IexRulesSchema(ctx context.Context) (*model.IexRulesSchema, error) {
	return iex.NewRulesEngine(iex.DefaultClient).RulesSchema()
}

func (r *queryResolver) KrakenServerTime(ctx context.Context) (*model.KrakenServerTime, error) {
	return kraken.NewMarketData(kraken.DefaultClient).ServerTime()
}

func (r *queryResolver) KrakenSystemStatus(ctx context.Context) (*model.KrakenSystemStatus, error) {
	return kraken.NewMarketData(kraken.DefaultClient).SystemStatus()
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) CoinbaseWallets(ctx context.Context) ([]*model.CoinbaseWallet, error) {
	return coinbase.NewCoinbaseAccounts(coinbase.DefaultClient).Wallets()
}
