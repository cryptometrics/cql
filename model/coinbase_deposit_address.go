package model

import (
	"cql/serial"
	"time"
)

type CoinbaseDepositAddress struct {
	ID                     string                         `json:"id"`
	Address                string                         `json:"address"`
	AddressInfo            CoinbaseDepositAddressInfo     `join:"address_info"`
	Name                   string                         `json:"name"`
	CreatedAt              time.Time                      `json:"created_at"`
	UpdatedAt              time.Time                      `json:"updated_at"`
	Network                string                         `json:"network"`
	URIScheme              string                         `json:"uri_scheme"`
	Resource               string                         `json:"resource"`
	ResourcePath           string                         `json:"resource_path"`
	Warnings               CoinbaseDepositAddressWarnings `json:"warnings"`
	LegacyAddress          string                         `json:"legacy_address"`
	DestinationTag         string                         `json:"destination_tag"`
	DepositURI             string                         `json:"deposit_uri"`
	CallbackURL            string                         `json:"callback_url"`
	ExchangeDepositAddress bool                           `json:"exchange_deposit_address"`
}

// UnmarshalJSON is an override required to parst strings from coinbases api
// into floats, specifically min_size and max_precision
func (address *CoinbaseDepositAddress) UnmarshalJSON(d []byte) error {
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}

	data.UnmarshalString("id", &address.ID)
	data.UnmarshalString("address", &address.Address)
	data.UnmarshalString("name", &address.Name)
	data.UnmarshalString("network", &address.Network)
	data.UnmarshalString("uri_scheme", &address.URIScheme)
	data.UnmarshalString("resource", &address.Resource)
	data.UnmarshalString("resource_path", &address.ResourcePath)
	data.UnmarshalString("legacy_address", &address.LegacyAddress)
	data.UnmarshalString("destination_tag", &address.DestinationTag)
	data.UnmarshalString("deposit_uri", &address.DepositURI)
	data.UnmarshalString("callback_url", &address.CallbackURL)
	data.UnmarshalBool("exchange_deposit_address", &address.ExchangeDepositAddress)

	err = data.UnmarshalTime(time.RFC3339Nano, "created_at", &address.CreatedAt)
	if err != nil {
		return err
	}

	err = data.UnmarshalTime(time.RFC3339Nano, "updated_at", &address.UpdatedAt)
	if err != nil {
		return err
	}

	address.AddressInfo = CoinbaseDepositAddressInfo{}
	if err := data.UnmarshalStruct("address_info", &address.AddressInfo); err != nil {
		return err
	}

	err = data.UnmarshalStructSlice("warnings", &address.Warnings,
		&CoinbaseDepositAddressWarning{})
	if err != nil {
		return err
	}
	return nil
}
