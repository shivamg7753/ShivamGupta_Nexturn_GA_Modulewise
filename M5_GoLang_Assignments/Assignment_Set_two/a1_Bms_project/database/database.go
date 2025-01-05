package database

import (
	"database/sql"
	"log"

	 _ "modernc.org/sqlite"
)

var DB *sql.DB

// Initialize initializes the SQLite database connection
func Initialize() {
	var err error
	DB, err = sql.Open("sqlite", "./blogs.db")
	if err != nil {
		log.Fatalf("Failed to connect to SQLite: %v", err)
	}

	createTableQuery := `
	CREATE TABLE IF NOT EXISTS blogs (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		content TEXT NOT NULL,
		author TEXT NOT NULL,
		timestamp DATETIME DEFAULT CURRENT_TIMESTAMP
	);`
	_, err = DB.Exec(createTableQuery)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
	log.Println("Database initialized successfully.")
}
