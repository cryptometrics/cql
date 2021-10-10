package coinbase

import (
	"cql/client"
	"cql/model"
	"cql/null"
)

// Transfer is an object that can be used to query on and execute transfers using the coinbase API.
type Transfer struct {
	parent
}

// NewTransfer returns a new Transfer object that will query using the provided client.
func NewTransfer(conn client.Connector) *Transfer {
	transfer := new(Transfer)
	construct(&transfer.parent, conn)
	return transfer
}

func (transfer *Transfer) fetchCoinbaseAccountDeposit(input *model.CoinbaseDepositInput) *client.FetchResponse {
	return transfer.conn.Fetch(&client.Request{
		Method:   client.POST,
		Endpoint: CoinbaseAccountDepositsEP,
		Body: client.JSONBody(map[string]null.Interface{
			"profile_id": null.NewInterfaceString(func() (s *string) {
				return input.ProfileID
			}()),
			"amount":              null.NewInterfaceFloat(&input.Amount),
			"coinbase_account_id": null.NewInterfaceString(&input.CoinbaseAccountID),
			"currency":            null.NewInterfaceString(&input.Currency),
		}),
	})
}

// Deposits funds from a www.coinbase.com wallet to the specified profile_id.
//
// Deposit funds from a coinbase account. You can move funds between your Coinbase accounts and your
// Coinbase Exchange trading accounts within your daily limits. Moving funds between Coinbase and
// Coinbase Exchange is instant and free. See the Coinbase Accounts section for retrieving your
// Coinbase accounts.
//
// This endpoint requires the "transfer" permission.
func (transfer *Transfer) DepositFromCoinbaseAccount(input *model.CoinbaseDepositInput) (m *model.CoinbaseDeposit, err error) {
	return m, transfer.fetchCoinbaseAccountDeposit(input).Assign(&m).Error
}
