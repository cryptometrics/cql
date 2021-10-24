package coinbase

import (
	"cql/client2"
	"cql/model"
)

// Transfer is an object that can be used to query on and execute transfers using the coinbase API.
type Transfer struct {
	parent
}

// NewTransfer returns a new Transfer object that will query using the provided
// client.
func NewTransfer(conn client2.Connector) *Transfer {
	transfer := new(Transfer)
	construct(&transfer.parent, conn)
	return transfer
}

// MakeCoinbaseAccountDeposit will deposit funds from a www.coinbase.com wallet
// to the specified profile_id.
//
// Deposit funds from a coinbase account. You can move funds between your
// Coinbase accounts and your Coinbase Exchange trading accounts within your
// daily limits. Moving funds between Coinbase and Coinbase Exchange is instant
// and free. See the Coinbase Accounts section for retrieving your Coinbase
// accounts.
//
// This endpoint requires the "transfer" permission.
//
// * source https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_postdepositcoinbaseaccount
func (transfer *Transfer) MakeCoinbaseAccountDeposit(input *model.MakeCoinbaseAccountDepositInput) (m *model.CoinbaseDeposit, err error) {
	return m, transfer.post(ENDPOINT_COINBASE_ACCOUNT_DEPOSITS).
		Body(client2.NewBody(client2.BODY_TYPE_JSON).
			SetString("profile_id", input.ProfileID).
			SetString("coinbase_account_id", &input.CoinbaseAccountID).
			SetString("currency", &input.Currency).
			SetFloat("amount", &input.Amount)).
		Fetch().Assign(&m).JoinMessages()
}

// MakePaymentMethodDeposit will deposits funds from a linked external payment
// method to the specified profile_id.
//
// Deposit funds from a payment method. See the Payment Methods section for
// retrieving your payment methods.
//
// This endpoint requires the "transfer" permission. API key must belong to
// default profile.
//
// * source https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_postdepositpaymentmethod
func (transfer *Transfer) MakePaymentMethodDeposit(input *model.MakeCoinbasePaymentMethodInput) (m *model.CoinbaseDeposit, err error) {
	return m, transfer.post(ENDPOINT_COINBASE_ACCOUNT_DEPOSITS).
		Body(client2.NewBody(client2.BODY_TYPE_JSON).
			SetString("profile_id", input.ProfileID).
			SetString("payment_method_id", &input.PaymentMethodID).
			SetString("currency", &input.Currency).
			SetFloat("amount", &input.Amount)).
		Fetch().Assign(&m).JoinMessages()
}

// PaymentMethods gets a list of the user's linked payment methods
//
//
// * source: https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_getpaymentmethods
func (transfer *Transfer) PaymentMethods() (m []*model.CoinbasePaymentMethod, err error) {
	req := transfer.get(ENDPOINT_TRANSFERS_PAYMENT_METHODS)
	return m, req.Fetch().Assign(&m).JoinMessages()
}
