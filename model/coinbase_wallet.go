package model

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
	data, err := newUmap(d)
	if err != nil {
		return err
	}

	data.unmarshalString("id", &wallet.ID)
	data.unmarshalString("name", &wallet.Name)
	data.unmarshalString("currency", &wallet.Currency)
	data.unmarshalString("type", &wallet.Type)
	data.unmarshalString("destination_tag_name", &wallet.DestinationTagName)
	data.unmarshalString("destination_tag_regex", &wallet.DestinationTagRegex)
	data.unmarshalString("hold_currency", &wallet.HoldCurrency)
	data.unmarshalBool("primary", &wallet.Primary)
	data.unmarshalBool("active", &wallet.Active)
	data.unmarshalBool("available_on_consumer", &wallet.AvailableOnConsumer)
	data.unmarshalBool("ready", &wallet.Ready)

	err = data.unmarshalFloatFromString("balance", &wallet.Balance)
	if err != nil {
		return err
	}

	err = data.unmarshalFloatFromString("hold_balance", &wallet.HoldBalance)
	if err != nil {
		return err
	}

	wallet.WireDepositInformation = CoinbaseWireDepositInformation{}
	err = data.unmarshalStruct("wire_deposit_information",
		&wallet.WireDepositInformation)
	if err != nil {
		return err
	}

	wallet.SWIFTDepositInformation = CoinbaseSWIFTDepositInformation{}
	err = data.unmarshalStruct("swift_deposit_information",
		&wallet.SWIFTDepositInformation)
	if err != nil {
		return err
	}

	wallet.SEPADepositInformation = CoinbaseSEPADepositInformation{}
	err = data.unmarshalStruct("sepa_deposit_information",
		&wallet.SEPADepositInformation)
	if err != nil {
		return err
	}

	wallet.UKDepositInformation = CoinbaseUKDepositInformation{}
	err = data.unmarshalStruct("uk_deposit_information",
		&wallet.UKDepositInformation)
	if err != nil {
		return err
	}

	return nil
}
