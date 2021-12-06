package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"cryptometrics/cql/coinbase"
	"cryptometrics/cql/graph/generated"
	"cryptometrics/cql/iex"
	"cryptometrics/cql/kraken"

	"github.com/cryptometrics/cql/model"
)

func (r *mutationResolver) CoinbaseAccountDeposit(ctx context.Context, opts *model.CoinbaseAccountDepositOptions) (*model.CoinbaseDeposit, error) {
	return coinbase.NewTransfer(coinbase.DefaultClient).CoinbaseAccountDeposit(opts)
}

func (r *mutationResolver) CoinbaseConvertCurrency(ctx context.Context, opts model.CoinbaseConversionsOptions) (*model.CoinbaseCurrencyConversion, error) {
	return coinbase.NewConversion(coinbase.DefaultClient).Make(opts)
}

func (r *mutationResolver) CoinbaseGenerateCryptoAddress(ctx context.Context, walletID string) (*model.CoinbaseCryptoAddress, error) {
	return coinbase.NewCoinbaseAccounts(coinbase.DefaultClient).GenerateCryptoAddress(walletID)
}

func (r *mutationResolver) CoinbasePaymentMethodDeposit(ctx context.Context, opts *model.CoinbasePaymentMethodDepositOptions) (*model.CoinbaseDeposit, error) {
	return coinbase.NewTransfer(coinbase.DefaultClient).PaymentMethodDeposit(opts)
}

func (r *mutationResolver) CoinbasePaymentMethodWithdrawal(ctx context.Context, opts *model.CoinbasePaymentMethodWithdrawalOptions) (*model.CoinbaseWithdrawal, error) {
	return coinbase.NewTransfer(coinbase.DefaultClient).PaymentMethodWithdrawal(opts)
}

func (r *mutationResolver) CoinbaseAccountWithdrawal(ctx context.Context, opts *model.CoinbaseAccountWithdrawalOptions) (*model.CoinbaseWithdrawal, error) {
	return coinbase.NewTransfer(coinbase.DefaultClient).AccountWithdrawal(opts)
}

func (r *mutationResolver) CoinbaseCryptoWithdrawal(ctx context.Context, opts *model.CoinbaseCryptoWithdrawalOptions) (*model.CoinbaseWithdrawal, error) {
	return coinbase.NewTransfer(coinbase.DefaultClient).CryptoWithdrawal(opts)
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

func (r *queryResolver) CoinbaseTransfers(ctx context.Context) ([]*model.CoinbaseAccountTransfer, error) {
	return coinbase.NewTransfer(coinbase.DefaultClient).All()
}

func (r *queryResolver) CoinbaseTransfer(ctx context.Context, transferID string) (*model.CoinbaseAccountTransfer, error) {
	return coinbase.NewTransfer(coinbase.DefaultClient).Find(transferID)
}

func (r *queryResolver) CoinbaseWallets(ctx context.Context) ([]*model.CoinbaseWallet, error) {
	return coinbase.NewCoinbaseAccounts(coinbase.DefaultClient).Wallets()
}

func (r *queryResolver) CoinbaseWithdrawalFeeEstimate(ctx context.Context, opts *model.CoinbaseWithdrawalFeeEstimateOptions) (*model.CoinbaseWithdrawalFeeEstimate, error) {
	return coinbase.NewTransfer(coinbase.DefaultClient).WithdrawalFeeEstimate(opts)
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
