package protomodel

import "cql/serial"

// * This is a generated file, do not edit

// KrakenSystemStatus holds data concerning the current system status or trading
// mode.
type KrakenSystemStatus struct {
	Error       []string                 `json:"error"`
	ProtoResult KrakenSystemStatusResult `json:"result"`
}

// UnmarshalJSON will deserialize bytes into a KrakenSystemStatus model
func (krakenSystemStatus *KrakenSystemStatus) UnmarshalJSON(d []byte) error {
	const (
		errorJsonTag  = "error"
		resultJsonTag = "result"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	data.UnmarshalStringSlice(errorJsonTag, &krakenSystemStatus.Error)
	krakenSystemStatus.ProtoResult = KrakenSystemStatusResult{}
	if err := data.UnmarshalStruct(resultJsonTag, &krakenSystemStatus.ProtoResult); err != nil {
		return err
	}
	return nil
}
