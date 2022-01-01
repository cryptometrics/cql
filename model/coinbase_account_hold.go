package model

import "github.com/cryptometrics/cql/protomodel"

// * This file was initialized by the meta-program, but is open to modification

// CoinbaseHold represents the hold on an account that belong to the same
// profile as the API key. Holds are placed on an account for any active orders
// or pending withdraw requests. As an order is filled, the hold amount is
// updated. If an order is canceled, any remaining hold is removed. For
// withdrawals, the hold is removed after it is completed.
type CoinbaseAccountHold struct{ protomodel.CoinbaseAccountHold }
