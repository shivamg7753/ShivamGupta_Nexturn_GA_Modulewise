package main

import (
	"log"
	"net/http"

	"a2_ems_project/database"
	"a2_ems_project/handlers"
	"a2_ems_project/middleware"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize the database
	database.InitDB()

	// Set up the router
	r := mux.NewRouter()

	// Apply global middleware
	r.Use(middleware.LoggingMiddleware)
	r.Use(middleware.BasicAuthMiddleware)

	// Define required fields for validation
	addProductRequiredFields := []string{"name", "description", "price", "stock", "category_id"}
	updateProductRequiredFields := []string{"name", "description", "price", "stock", "category_id"}

	// Routes with validation middleware
	r.Handle("/product", middleware.ValidationMiddleware(addProductRequiredFields)(http.HandlerFunc(handlers.AddProduct))).Methods("POST")
	r.Handle("/product/{id}", middleware.ValidationMiddleware(updateProductRequiredFields)(http.HandlerFunc(handlers.UpdateProduct))).Methods("PUT")

	// Routes without validation middleware
	r.HandleFunc("/product/{id}", handlers.FetchProduct).Methods("GET")
	r.HandleFunc("/product/{id}", handlers.DeleteProduct).Methods("DELETE")

	// Start the server
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
