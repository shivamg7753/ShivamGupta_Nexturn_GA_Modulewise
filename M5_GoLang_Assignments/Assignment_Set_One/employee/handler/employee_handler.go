package handler

import (
	"employee/model"
	"employee/service"
	"fmt"
	"strconv"
	"strings"
)

// AddEmployee handles adding a new employee
func AddEmployee() {
	var id, age int
	var name, department string

	fmt.Print("Enter ID: ")
	fmt.Scan(&id)
	fmt.Print("Enter Name: ")
	fmt.Scan(&name)
	fmt.Print("Enter Age: ")
	fmt.Scan(&age)
	fmt.Print("Enter Department: ")
	fmt.Scan(&department)

	// Validate input
	if err := service.ValidateEmployee(id, age); err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Add employee to the list
	model.Employees = append(model.Employees, model.Employee{
		ID:         id,
		Name:       name,
		Age:        age,
		Department: department,
	})
	fmt.Println("Employee added successfully!")
}

// SearchEmployee handles searching for an employee by ID or name
func SearchEmployee() {
	var searchType, searchTerm string
	fmt.Print("Search by (id/name): ")
	fmt.Scan(&searchType)
	fmt.Print("Enter search term: ")
	fmt.Scan(&searchTerm)

	var foundEmployee *model.Employee
	var err error

	if strings.ToLower(searchType) == "id" {
		id, _ := strconv.Atoi(searchTerm)
		foundEmployee, err = service.FindEmployeeByID(id)
	} else if strings.ToLower(searchType) == "name" {
		foundEmployee, err = service.FindEmployeeByName(searchTerm)
	} else {
		fmt.Println("Invalid search type. Please use 'id' or 'name'.")
		return
	}

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Employee found: %+v\n", *foundEmployee)
}

// ListEmployeesByDepartment handles listing employees in a department
func ListEmployeesByDepartment() {
	var department string
	fmt.Print("Enter Department: ")
	fmt.Scan(&department)

	fmt.Printf("Employees in department %s:\n", department)
	found := false
	for _, emp := range model.Employees {
		if strings.EqualFold(emp.Department, department) {
			fmt.Printf("%+v\n", emp)
			found = true
		}
	}
	if !found {
		fmt.Println("No employees found in this department.")
	}
}

// CountEmployeesByDepartment handles counting employees in a department
func CountEmployeesByDepartment() {
	const departments = "HR, IT"
	var department string
	fmt.Print("Enter Department (HR/IT): ")
	fmt.Scan(&department)

	if !strings.Contains(departments, strings.ToUpper(department)) {
		fmt.Println("Invalid department!")
		return
	}

	count := 0
	for _, emp := range model.Employees {
		if strings.EqualFold(emp.Department, department) {
			count++
		}
	}

	fmt.Printf("Total employees in %s: %d\n", department, count)
}
