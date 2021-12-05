package coinbase

import (
	"cql/client"
)

// CoinbaseAccounts is an object used to query coinbase account data.
type Transfer struct {
	client.Parent
}

// NewCurrency will return an object to query currency currencys
func NewTransfer(conn client.Connector) *Transfer {
	transfer := new(Transfer)
	client.ConstructParent(&transfer.Parent, conn)
	return transfer
}

// // All gets a list of in-progress and completed transfers of funds in/out of any
// // of the user's accounts.
// //
// // This endpoint requires either the "view" or "trade" permission.
// //
// // * source: https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_gettransfers
// func (transfer *Transfer) All() (m []*model.CoinbaseTransfer, err error) {
// 	return m, transfer.get(ENDPOINT_TRANSFERS).Fetch().Assign(&m).JoinMessages()
// }

// // Find get information on a single transfer.
// //
// // This endpoint requires either the "view" or "trade" permission.
// //
// // * source: https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_gettransfers
// func (transfer *Transfer) Find(id string) (m *model.CoinbaseTransfer, err error) {
// 	req := transfer.get(ENDPOINT_TRANSFER)
// 	return m, req.PathParam("id", id).Fetch().Assign(&m).JoinMessages()
// }

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
// func (transfer *Transfer) MakeCoinbaseAccountDeposit(input *model.MakeCoinbaseAccountDepositInput) (m *model.CoinbaseDeposit, err error) {
// 	return m, transfer.post(ENDPOINT_COINBASE_ACCOUNT_DEPOSITS).
// 		Body(client.NewBody(client.BODY_TYPE_JSON).
// 			SetString("profile_id", input.ProfileID).
// 			SetString("coinbase_account_id", &input.CoinbaseAccountID).
// 			SetString("currency", &input.Currency).
// 			SetFloat("amount", &input.Amount)).
// 		Fetch().Assign(&m).JoinMessages()
// }

// // MakePaymentMethodDeposit will deposits funds from a linked external payment
// // method to the specified profile_id.
// //
// // Deposit funds from a payment method. See the Payment Methods section for
// // retrieving your payment methods.
// //
// // This endpoint requires the "transfer" permission. API key must belong to
// // default profile.
// //
// // * source https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_postdepositpaymentmethod
// func (transfer *Transfer) MakePaymentMethodDeposit(input *model.MakeCoinbasePaymentMethodInput) (m *model.CoinbaseDeposit, err error) {
// 	return m, transfer.post(ENDPOINT_COINBASE_ACCOUNT_DEPOSITS).
// 		Body(client.NewBody(client.BODY_TYPE_JSON).
// 			SetString("profile_id", input.ProfileID).
// 			SetString("payment_method_id", &input.PaymentMethodID).
// 			SetString("currency", &input.Currency).
// 			SetFloat("amount", &input.Amount)).
// 		Fetch().Assign(&m).JoinMessages()
// }

// // PaymentMethods gets a list of the user's linked payment methods
// //
// //
// // * source: https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_getpaymentmethods
// func (transfer *Transfer) PaymentMethods() (m []*model.CoinbasePaymentMethod, err error) {
// 	req := transfer.get(ENDPOINT_TRANSFERS_PAYMENT_METHODS)
// 	return m, req.Fetch().Assign(&m).JoinMessages()
// }
