package main

import (
	"fmt"
	"inventory/service"
)

func main() {
	for {
		fmt.Println("\nInventory Management System")
		fmt.Println("1. Add Product")
		fmt.Println("2. Update Stock")
		fmt.Println("3. Search Product")
		fmt.Println("4. Display Inventory")
		fmt.Println("5. Sort by Price")
		fmt.Println("6. Sort by Stock")
		fmt.Println("7. Exit")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			var id, stock int
			var name string
			var price float64

			fmt.Print("Enter Product ID: ")
			fmt.Scanln(&id)
			fmt.Print("Enter Product Name: ")
			fmt.Scanln(&name)
			fmt.Print("Enter Product Price: ")
			fmt.Scanln(&price)
			fmt.Print("Enter Product Stock: ")
			fmt.Scanln(&stock)

			if err := service.AddProduct(id, name, price, stock); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Product added successfully!")
			}
		case 2:
			var id, stock int
			fmt.Print("Enter Product ID to update stock: ")
			fmt.Scanln(&id)
			fmt.Print("Enter New Stock: ")
			fmt.Scanln(&stock)

			if err := service.UpdateStock(id, stock); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Stock updated successfully!")
			}
		case 3:
			var query string
			fmt.Print("Enter Product ID or Name to search: ")
			fmt.Scanln(&query)

			product, err := service.SearchProduct(query)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("Product Found: %+v\n", *product)
			}
		case 4:
			service.DisplayInventory()
		case 5:
			service.SortInventoryByPrice()
			fmt.Println("Inventory sorted by price.")
		case 6:
			service.SortInventoryByStock()
			fmt.Println("Inventory sorted by stock.")
		case 7:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
