package model

import "cql/serial"

type CoinbaseWallet struct {
	ID                      string                          `json:"id"`
	Name                    string                          `json:"name"`
	Balance                 float64                         `json:"balance"`
	Currency                string                          `json:"currency"`
	Type                    string                          `json:"type"`
	Primary                 bool                            `json:"primary"`
	Active                  bool                            `json:"active"`
	AvailableOnConsumer     bool                            `json:"available_on_consumer"`
	Ready                   bool                            `json:"ready"`
	WireDepositInformation  CoinbaseWireDepositInformation  `json:"wire_deposit_information"`
	SWIFTDepositInformation CoinbaseSWIFTDepositInformation `json:"swift_deposit_information"`
	SEPADepositInformation  CoinbaseSEPADepositInformation  `json:"sepa_deposit_information"`
	UKDepositInformation    CoinbaseUKDepositInformation    `json:"uk_deposit_information"`
	DestinationTagName      string                          `json:"destination_tag_name"`
	DestinationTagRegex     string                          `json:"destination_tag_regex"`
	HoldBalance             float64                         `json:"hold_balance"`
	HoldCurrency            string                          `json:"hold_currency"`
}

// UnmarshalJSON is an override required to parst strings from coinbases api
// into floats, specifically min_size and max_precision
func (wallet *CoinbaseWallet) UnmarshalJSON(d []byte) error {
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}

	data.UnmarshalString("id", &wallet.ID)
	data.UnmarshalString("name", &wallet.Name)
	data.UnmarshalString("currency", &wallet.Currency)
	data.UnmarshalString("type", &wallet.Type)
	data.UnmarshalString("destination_tag_name", &wallet.DestinationTagName)
	data.UnmarshalString("destination_tag_regex", &wallet.DestinationTagRegex)
	data.UnmarshalString("hold_currency", &wallet.HoldCurrency)
	data.UnmarshalBool("primary", &wallet.Primary)
	data.UnmarshalBool("active", &wallet.Active)
	data.UnmarshalBool("available_on_consumer", &wallet.AvailableOnConsumer)
	data.UnmarshalBool("ready", &wallet.Ready)

	err = data.UnmarshalFloatFromString("balance", &wallet.Balance)
	if err != nil {
		return err
	}

	err = data.UnmarshalFloatFromString("hold_balance", &wallet.HoldBalance)
	if err != nil {
		return err
	}

	wallet.WireDepositInformation = CoinbaseWireDepositInformation{}
	err = data.UnmarshalStruct("wire_deposit_information",
		&wallet.WireDepositInformation)
	if err != nil {
		return err
	}

	wallet.SWIFTDepositInformation = CoinbaseSWIFTDepositInformation{}
	err = data.UnmarshalStruct("swift_deposit_information",
		&wallet.SWIFTDepositInformation)
	if err != nil {
		return err
	}

	wallet.SEPADepositInformation = CoinbaseSEPADepositInformation{}
	err = data.UnmarshalStruct("sepa_deposit_information",
		&wallet.SEPADepositInformation)
	if err != nil {
		return err
	}

	wallet.UKDepositInformation = CoinbaseUKDepositInformation{}
	err = data.UnmarshalStruct("uk_deposit_information",
		&wallet.UKDepositInformation)
	if err != nil {
		return err
	}

	return nil
}
