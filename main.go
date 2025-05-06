package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3" // Import SQLite driver
	"log"
	"os"
	"time"
)

type Job struct {
	ID   int
	Name string
	Date int64
}

func AddJob(db *sql.DB, name string, date int64) {
	insertSQL := "INSERT INTO jobs (Name, Date) VALUES (?, ?)"
	_, err := db.Exec(insertSQL, name, date)
	if err != nil {
		log.Fatalf("Error inserting job: %v", err)
	}
	fmt.Printf("Job '%s' added successfully.\n", name)
}

func closeDB(db *sql.DB) {
	if err := db.Close(); err != nil {
		log.Printf("Error closing database connection: %v", err)
	}
}

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

	currentDate := time.Now()
	var resultDate = currentDate.Format("2006-01-02")
	fmt.Println(resultDate)
	AddJob(db, "jobs", currentDate.Unix())
}
