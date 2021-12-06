package protomodel

import (
	"cryptometrics/cql/scalar"
	"cryptometrics/cql/serial"
	"time"
)

// * This is a generated file, do not edit

// KrakenSystemStatusResult holds data concerning the current system status or
// trading mode.
type KrakenSystemStatusResult struct {
	// Current timestamp (RFC3339)
	Timestamp time.Time `json:"timestamp"`

	// Enum: "online" "maintenance" "cancel_only" "post_only" Current system status
	Status scalar.SystemStatus `json:"status"`
}

// UnmarshalJSON will deserialize bytes into a KrakenSystemStatusResult model
func (krakenSystemStatusResult *KrakenSystemStatusResult) UnmarshalJSON(d []byte) error {
	const (
		statusJsonTag    = "status"
		timestampJsonTag = "timestamp"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	data.UnmarshalSystemStatus(statusJsonTag, &krakenSystemStatusResult.Status)
	err = data.UnmarshalTime(time.RFC3339Nano, timestampJsonTag, &krakenSystemStatusResult.Timestamp)
	if err != nil {
		return err
	}
	return nil
}
