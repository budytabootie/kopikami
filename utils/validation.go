package utils

import (
	"errors"
)

func ValidatePrice(price float64) error {
	if price <= 0 {
		return errors.New("price must be greater than zero")
	}
	return nil
}
