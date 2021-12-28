package protomodel

import (
	"time"

	"github.com/cryptometrics/cql/scalar"
)

// * This is a generated file, do not edit

// KrakenSystemStatusResult holds data concerning the current system status or// trading mode.
type KrakenSystemStatusResult struct {
	// Current timestamp (RFC3339)
	Timestamp time.Time `json:"timestamp"`

	// Enum: "online" "maintenance" "cancel_only" "post_only" Current system status
	Status scalar.SystemStatus `json:"status"`
}
