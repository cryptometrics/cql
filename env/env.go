package env

import "os"

type Variable int

const (
	_ Variable = iota
	CoinbaseProURL
	CoinbaseProAccessPassphrase
	CoinbaseProAccessKey
	CoinbaseProSecret
)

// Name will return the variable name of the env variable.
func (variable Variable) Name() string {
	return map[Variable]string{
		CoinbaseProURL:              "CB_PRO_URL",
		CoinbaseProAccessPassphrase: "CB_PRO_ACCESS_PASSPHRASE",
		CoinbaseProAccessKey:        "CB_PRO_ACCESS_KEY",
		CoinbaseProSecret:           "CB_PRO_SECRET",
	}[variable]
}

// Get will return the environment variable for a variable-type constant as a
// string.
func (variable Variable) Get() string {
	return os.Getenv(variable.Name())
}
