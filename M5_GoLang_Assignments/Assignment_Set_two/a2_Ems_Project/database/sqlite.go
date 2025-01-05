package database

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

// Initialize SQLite database connection
func InitDB() {
	var err error
	DB, err = sql.Open("sqlite", "./inventory.db")
	if err != nil {
		log.Fatalf("Failed to connect to SQLite database: %v", err)
	}

	// Create `products` table if it doesn't exist
	query := `
	CREATE TABLE IF NOT EXISTS products (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		price REAL NOT NULL,
		stock INTEGER NOT NULL,
		category_id INTEGER NOT NULL
	);
	`
	_, err = DB.Exec(query)
	if err != nil {
		log.Fatalf("Failed to create products table: %v", err)
	}
	log.Println("Database initialized successfully.")
}
