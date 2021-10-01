package coinbase

import (
	"fmt"
	"strconv"
)

func IntPtrStringPtr(i *int) *string {
	if i != nil {
		str := strconv.Itoa(*i)
		return &str
	}
	return nil
}

func FloatPtrStringPtr(f *float64) *string {
	if f == nil {
		return nil
	}
	str := fmt.Sprintf("%f", *f)
	return &str
}

func StringStringPtr(str string) *string {
	return &str
}
