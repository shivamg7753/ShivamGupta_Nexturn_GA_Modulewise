package service 

import (
	"errors"
	"fmt"
	"inventory/model"
	"sort"
)

func AddProduct(id int, name string, price float64, stock int) error {
	for  _, product := range model.Inventory {
		if product.ID == id {
			return errors.New("product ID already exists")
		}
	}

	model.Inventory = append(model.Inventory, model.Product {
		ID: id,
		Name: name,
		Price: price,
		Stock: stock,
	})
	return nil 
}

func UpdateStock(id int, newStock int) error {
	if newStock < 0 {
		return errors.New("stock cannot be negative")
	}
	for i, product := range model.Inventory {
		if product.ID == id {
			model.Inventory[i].Stock = newStock
			return nil 
		}
	}
	return errors.New("product not found")
}

func SearchProduct(query string) (*model.Product, error){
	for _, product := range model.Inventory {
		if fmt.Sprintf("%d", product.ID) == query || product.Name == query {
			return &product, nil 
		}
	}
	return nil, errors.New("product not found")
}

func DisplayInventory() {
	fmt.Println("ID\tName\t\tPrice\tStock")
	fmt.Println("----------------------------------")
	for _, product := range model.Inventory {
		fmt.Printf("%d\t%s\t\t%.2f\t%d\n", product.ID, product.Name, product.Price, product.Stock)
	}
}

func SortInventoryByPrice() {
	sort.Slice(model.Inventory, func(i, j int) bool {
		return model.Inventory[i].Price < model.Inventory[j].Price
	})
}

// SortInventoryByStock sorts the inventory by stock
func SortInventoryByStock() {
	sort.Slice(model.Inventory, func(i, j int) bool {
		return model.Inventory[i].Stock < model.Inventory[j].Stock
	})
}