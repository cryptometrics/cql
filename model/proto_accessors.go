package model

// * This is a generated file, do not edit

// Details converts the native protomodel ProtoDetails to a local
// CoinbaseAccountTransferDetails.
func (m *CoinbaseAccountTransfer) Details() *CoinbaseAccountTransferDetails {
	return &CoinbaseAccountTransferDetails{
		CoinbaseAccountTransferDetails: m.ProtoDetails}
}

// BankCountry converts the native protomodel ProtoBankCountry to a local
// CoinbaseBankCountry.
func (m *CoinbaseSwiftDepositInformation) BankCountry() *CoinbaseBankCountry {
	return &CoinbaseBankCountry{
		CoinbaseBankCountry: m.ProtoBankCountry}
}

// Result converts the native protomodel ProtoResult to a local
// KrakenServerTimeResult.
func (m *KrakenServerTime) Result() *KrakenServerTimeResult {
	return &KrakenServerTimeResult{
		KrakenServerTimeResult: m.ProtoResult}
}

// Schema converts the native protomodel ProtoSchema to a local
// []*IexRulesScheme.
func (m *IexRulesSchema) Schema() (slice []*IexRulesScheme) {
	for _, v := range m.ProtoSchema {
		slice = append(slice, &IexRulesScheme{IexRulesScheme: *v})
	}
	return
}

// Details converts the native protomodel ProtoDetails to a local
// CoinbaseAccountLedgerDetails.
func (m *CoinbaseAccountLedger) Details() *CoinbaseAccountLedgerDetails {
	return &CoinbaseAccountLedgerDetails{
		CoinbaseAccountLedgerDetails: m.ProtoDetails}
}

// BankCountry converts the native protomodel ProtoBankCountry to a local
// CoinbaseBankCountry.
func (m *CoinbaseUkDepositInformation) BankCountry() *CoinbaseBankCountry {
	return &CoinbaseBankCountry{
		CoinbaseBankCountry: m.ProtoBankCountry}
}

// Result converts the native protomodel ProtoResult to a local
// KrakenSystemStatusResult.
func (m *KrakenSystemStatus) Result() *KrakenSystemStatusResult {
	return &KrakenSystemStatusResult{
		KrakenSystemStatusResult: m.ProtoResult}
}

// BankCountry converts the native protomodel ProtoBankCountry to a local
// CoinbaseBankCountry.
func (m *CoinbaseSepaDepositInformation) BankCountry() *CoinbaseBankCountry {
	return &CoinbaseBankCountry{
		CoinbaseBankCountry: m.ProtoBankCountry}
}

// BankCountry converts the native protomodel ProtoBankCountry to a local
// CoinbaseBankCountry.
func (m *CoinbaseWireDepositInformation) BankCountry() *CoinbaseBankCountry {
	return &CoinbaseBankCountry{
		CoinbaseBankCountry: m.ProtoBankCountry}
}

// WireDepositInformation converts the native protomodel
// ProtoWireDepositInformation to a local CoinbaseWireDepositInformation.
func (m *CoinbaseWallet) WireDepositInformation() *CoinbaseWireDepositInformation {
	return &CoinbaseWireDepositInformation{
		CoinbaseWireDepositInformation: m.ProtoWireDepositInformation}
}

// SwiftDepositInformation converts the native protomodel
// ProtoSwiftDepositInformation to a local CoinbaseSwiftDepositInformation.
func (m *CoinbaseWallet) SwiftDepositInformation() *CoinbaseSwiftDepositInformation {
	return &CoinbaseSwiftDepositInformation{
		CoinbaseSwiftDepositInformation: m.ProtoSwiftDepositInformation}
}

// SepaDepositInformation converts the native protomodel
// ProtoSepaDepositInformation to a local CoinbaseSepaDepositInformation.
func (m *CoinbaseWallet) SepaDepositInformation() *CoinbaseSepaDepositInformation {
	return &CoinbaseSepaDepositInformation{
		CoinbaseSepaDepositInformation: m.ProtoSepaDepositInformation}
}

// UkDepositInformation converts the native protomodel ProtoUkDepositInformation
// to a local CoinbaseUkDepositInformation.
func (m *CoinbaseWallet) UkDepositInformation() *CoinbaseUkDepositInformation {
	return &CoinbaseUkDepositInformation{
		CoinbaseUkDepositInformation: m.ProtoUkDepositInformation}
}
