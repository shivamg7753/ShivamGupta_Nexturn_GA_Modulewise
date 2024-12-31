package main

import (
	"employee/handler"
	"fmt"
)

func main() {
	for {
		fmt.Println("\nEmployee Management System")
		fmt.Println("1. Add Employee")
		fmt.Println("2. Search Employee")
		fmt.Println("3. List Employees by Department")
		fmt.Println("4. Count Employees by Department")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			handler.AddEmployee()
		case 2:
			handler.SearchEmployee()
		case 3:
			handler.ListEmployeesByDepartment()
		case 4:
			handler.CountEmployeesByDepartment()
		case 5:
			fmt.Println("Exiting program. Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
