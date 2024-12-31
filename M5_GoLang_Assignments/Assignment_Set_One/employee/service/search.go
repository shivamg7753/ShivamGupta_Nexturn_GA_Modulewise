package service

import (
	"employee/model"
	"errors"
	"strings"
)

// FindEmployeeByID searches for an employee by ID
func FindEmployeeByID(id int) (*model.Employee, error) {
	for _, emp := range model.Employees {
		if emp.ID == id {
			return &emp, nil
		}
	}
	return nil, errors.New("employee not found")
}

// FindEmployeeByName searches for an employee by name
func FindEmployeeByName(name string) (*model.Employee, error) {
	for _, emp := range model.Employees {
		if strings.EqualFold(emp.Name, name) {
			return &emp, nil
		}
	}
	return nil, errors.New("employee not found")
}
