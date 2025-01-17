package utils

import (
	"errors"
)

// ValidatePrice memvalidasi apakah harga yang diberikan lebih besar dari nol
// Mengembalikan error jika harga kurang dari atau sama dengan nol
func ValidatePrice(price float64) error {
	if price <= 0 {
		return errors.New("price must be greater than zero")
	}
	return nil
}
