package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3" // Import SQLite driver
)

func main() {
	fmt.Println("Starting application...")

	dbFileName := "test.db"

	if err := os.Remove(dbFileName); err != nil && !os.IsNotExist(err) {
		log.Fatalf("Error removing database file: %v", err)
	}

	db, err := sql.Open("sqlite3", dbFileName)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer closeDB(db)

	fmt.Println("Successfully connected to the database.")

	const createTableSQL = `
	CREATE TABLE IF NOT EXISTS jobs (
		ID INTEGER PRIMARY KEY AUTOINCREMENT,
		Name TEXT,
		Date INTEGER
	);`

	fmt.Println("Creating or checking 'jobs' table...")
	if _, err := db.Exec(createTableSQL); err != nil {

		log.Fatalf("Error creating or checking 'jobs' table: %v", err)
	}

	fmt.Println("Table 'jobs' is ready.")
	fmt.Println("Application setup complete.")
}

// Helper function to close the database connection and handle errors.
func closeDB(db *sql.DB) {
	if err := db.Close(); err != nil {
		log.Printf("Error closing database connection: %v", err)
	}
}
