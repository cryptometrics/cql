package protomodel

import (
	"time"

	"github.com/cryptometrics/cql/serial"
)

// * This is a generated file, do not edit

// KrakenServerTimeResult holds data concerning the server time
type KrakenServerTimeResult struct {
	// RFC 1123 time format
	Rfc1123 time.Time `json:"rfc1123"`

	// Unix timestamp
	Unixtime int `json:"unixtime"`
}

// UnmarshalJSON will deserialize bytes into a KrakenServerTimeResult model
func (krakenServerTimeResult *KrakenServerTimeResult) UnmarshalJSON(d []byte) error {
	const (
		rfc1123JsonTag  = "rfc1123"
		unixtimeJsonTag = "unixtime"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	data.UnmarshalInt(unixtimeJsonTag, &krakenServerTimeResult.Unixtime)
	err = data.UnmarshalTime(KrakenRFC1123Layout, rfc1123JsonTag, &krakenServerTimeResult.Rfc1123)
	if err != nil {
		return err
	}
	return nil
}
