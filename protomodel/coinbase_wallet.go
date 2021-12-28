package protomodel

// * This is a generated file, do not edit

// CoinbaseWallet represents a user's available Coinbase wallet (These are the// wallets/accounts that are used for buying and selling on www.coinbase.com)
type CoinbaseWallet struct {
	Active                       bool                            `json:"active"`
	AvailableOnConsumer          bool                            `json:"available_on_consumer"`
	Balance                      float64                         `json:"balance"`
	Currency                     string                          `json:"currency"`
	DestinationTagName           string                          `json:"destination_tag_name"`
	DestinationTagRegex          string                          `json:"destination_tag_regex"`
	HoldBalance                  float64                         `json:"hold_balance"`
	HoldCurrency                 string                          `json:"hold_currency"`
	Id                           string                          `json:"id"`
	Name                         string                          `json:"name"`
	Primary                      bool                            `json:"primary"`
	ProtoSepaDepositInformation  CoinbaseSepaDepositInformation  `json:"sepa_deposit_information"`
	ProtoSwiftDepositInformation CoinbaseSwiftDepositInformation `json:"swift_deposit_information"`
	ProtoUkDepositInformation    CoinbaseUkDepositInformation    `json:"uk_deposit_information"`
	ProtoWireDepositInformation  CoinbaseWireDepositInformation  `json:"wire_deposit_information"`
	Ready                        bool                            `json:"ready"`
	Type                         string                          `json:"type"`
}
