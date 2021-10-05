package model

import "time"

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
	data, err := newUmap(d)
	if err != nil {
		return err
	}

	data.unmarshalString("id", &address.ID)
	data.unmarshalString("address", &address.Address)
	data.unmarshalString("name", &address.Name)
	data.unmarshalString("network", &address.Network)
	data.unmarshalString("uri_scheme", &address.URIScheme)
	data.unmarshalString("resource", &address.Resource)
	data.unmarshalString("resource_path", &address.ResourcePath)
	data.unmarshalString("legacy_address", &address.LegacyAddress)
	data.unmarshalString("destination_tag", &address.DestinationTag)
	data.unmarshalString("deposit_uri", &address.DepositURI)
	data.unmarshalString("callback_url", &address.CallbackURL)
	data.unmarshalBool("exchange_deposit_address", &address.ExchangeDepositAddress)

	err = data.unmarshalTime(time.RFC3339Nano, "created_at", &address.CreatedAt)
	if err != nil {
		return err
	}

	err = data.unmarshalTime(time.RFC3339Nano, "updated_at", &address.UpdatedAt)
	if err != nil {
		return err
	}

	address.AddressInfo = CoinbaseDepositAddressInfo{}
	if err := data.unmarshalStruct("address_info", &address.AddressInfo); err != nil {
		return err
	}

	err = data.unmarshalStructSlice("warnings", &address.Warnings,
		&CoinbaseDepositAddressWarning{})
	if err != nil {
		return err
	}
	return nil
}
