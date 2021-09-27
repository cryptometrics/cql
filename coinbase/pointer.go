package coinbase

import "strconv"

func IntString(i *int) *string {
	if i != nil {
		str := strconv.Itoa(*i)
		return &str
	}
	return nil
}
