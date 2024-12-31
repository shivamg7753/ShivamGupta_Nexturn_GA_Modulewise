package main

import (
	"log"
	"net/http"

	"a1_bms_project/database"
	"a1_bms_project/handlers"
	"a1_bms_project/middleware"
)

func main() {
	// Initialize the database
	database.Initialize()

	// Define routes
	http.HandleFunc("/blog", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			handlers.CreateBlog(w, r)
		case "GET":
			handlers.GetBlog(w, r)
		case "PUT":
			handlers.UpdateBlog(w, r)
		case "DELETE":
			handlers.DeleteBlog(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/blogs", handlers.GetAllBlogs)

	// Wrap with middleware
	loggedRouter := middleware.LoggingMiddleware(http.DefaultServeMux)

	// Start the server
	log.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", loggedRouter); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
