package model

import "github.com/cryptometrics/cql/protomodel"

// * This file was initialized by schema/generate.py, but is open to extension

// Rule to evaluate thousands of data points per second and build event-driven,
// automated alerts using Rules Engine. You can access Rules Engine through the
// IEX Console or through our API using the guidelines below.
type IexRule struct{ protomodel.IexRule }
