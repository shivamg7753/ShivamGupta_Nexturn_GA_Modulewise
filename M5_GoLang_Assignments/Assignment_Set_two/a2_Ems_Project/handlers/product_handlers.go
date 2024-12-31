package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"a2_ems_project/database"
	"a2_ems_project/models"
	"github.com/gorilla/mux"
)

// AddProduct handles POST /product
func AddProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	query := `INSERT INTO products (name, description, price, stock, category_id) VALUES (?, ?, ?, ?, ?)`
	_, err := database.DB.Exec(query, product.Name, product.Description, product.Price, product.Stock, product.CategoryID)
	if err != nil {
		http.Error(w, "Failed to add product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Product added successfully")
}

// FetchProduct handles GET /product/{id}
func FetchProduct(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	query := `SELECT id, name, description, price, stock, category_id FROM products WHERE id = ?`

	row := database.DB.QueryRow(query, id)
	var product models.Product
	err := row.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Stock, &product.CategoryID)
	if err == sql.ErrNoRows {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Failed to fetch product", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(product)
}

// UpdateProduct handles PUT /product/{id}
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	query := `UPDATE products SET name = ?, description = ?, price = ?, stock = ?, category_id = ? WHERE id = ?`
	_, err := database.DB.Exec(query, product.Name, product.Description, product.Price, product.Stock, product.CategoryID, id)
	if err != nil {
		http.Error(w, "Failed to update product", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "Product updated successfully")
}

// DeleteProduct handles DELETE /product/{id}
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	query := `DELETE FROM products WHERE id = ?`
	_, err := database.DB.Exec(query, id)
	if err != nil {
		http.Error(w, "Failed to delete product", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "Product deleted successfully")
}
