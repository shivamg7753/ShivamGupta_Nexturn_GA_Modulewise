package model 

type Product struct {
	ID  int 
	Name string 
	Price float64 
	Stock int 
}

var Inventory []Product 