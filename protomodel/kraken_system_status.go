package protomodel

// * This is a generated file, do not edit

// KrakenSystemStatus holds data concerning the current system status or trading// mode.
type KrakenSystemStatus struct {
	Error       []string                 `json:"error"`
	ProtoResult KrakenSystemStatusResult `json:"result"`
}
