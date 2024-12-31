package service

import "errors"

// ValidateAmount ensures the amount is positive
func ValidateAmount(amount float64) error {
	if amount <= 0 {
		return errors.New("amount must be greater than zero")
	}
	return nil
}
