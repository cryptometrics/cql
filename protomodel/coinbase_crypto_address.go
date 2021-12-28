package protomodel

import "time"

// * This is a generated file, do not edit

// CoinbaseCryptoAddress is used for a one-time crypto address for depositing// crypto.
type CoinbaseCryptoAddress struct {
	Address          string                          `json:"address"`
	CallbackUrl      string                          `json:"callback_url"`
	CreateAt         time.Time                       `json:"create_at"`
	DepositUri       string                          `json:"deposit_uri"`
	DestinationTag   string                          `json:"destination_tag"`
	Id               string                          `json:"id"`
	LegacyAddress    string                          `json:"legacy_address"`
	Name             string                          `json:"name"`
	Network          string                          `json:"network"`
	ProtoAddressInfo CoinbaseCryptoAddressInfo       `json:"address_info"`
	ProtoWarnings    []*CoinbaseCryptoAddressWarning `json:"warnings"`
	Resource         string                          `json:"resource"`
	ResourcePath     string                          `json:"resource_path"`
	UpdatedAt        time.Time                       `json:"updated_at"`
	UriScheme        string                          `json:"uri_scheme"`
}
