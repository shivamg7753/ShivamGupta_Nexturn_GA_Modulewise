package service

import (
	"errors"
	"employee/model"
)

// ValidateEmployee ensures that the employee input is valid
func ValidateEmployee(id, age int) error {
	// Check for unique ID
	for _, emp := range model.Employees {
		if emp.ID == id {
			return errors.New("employee ID must be unique")
		}
	}
	// Check age validation
	if age <= 18 {
		return errors.New("employee age must be greater than 18")
	}
	return nil
}
