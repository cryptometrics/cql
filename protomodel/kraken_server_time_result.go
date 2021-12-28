package protomodel

import "time"

// * This is a generated file, do not edit

// KrakenServerTimeResult holds data concerning the server time
type KrakenServerTimeResult struct {
	// RFC 1123 time format
	Rfc1123 time.Time `json:"rfc1123"`

	// Unix timestamp
	Unixtime int `json:"unixtime"`
}
