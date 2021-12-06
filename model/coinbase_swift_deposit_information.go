package model

import "cryptometrics/cql/protomodel"

// * This file was initialized by schema/generate.py, but is open to extension

// CoinbaseSwiftDepositInformation information regarding a wallet's deposits.
// SWIFT stands for Society for Worldwide Interbank Financial
// Telecommunications. Basically, it's a computer network that connects over 900
// banks around the world – and enables them to transfer money. ING is part of
// this network. There is no fee for accepting deposits into your account with
// ING.
type CoinbaseSwiftDepositInformation struct {
	protomodel.CoinbaseSwiftDepositInformation
}
