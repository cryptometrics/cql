package model

// * This file was initialized by the meta-program, but is open to modification
import "github.com/cryptometrics/cql/protomodel"

// CoinbaseSepaDepositInformation information regarding a wallet's deposits. A
// SEPA credit transfer is a single transfer of Euros from one person or
// organisation to another. For example, this could be to pay the deposit for a
// holiday rental or to settle an invoice. A SEPA direct debit is a recurring
// payment, for example to pay monthly rent or for a service like a mobile phone
// contract.
type CoinbaseSepaDepositInformation struct {
	protomodel.CoinbaseSepaDepositInformation
}
