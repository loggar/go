package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // go get github.com/lib/pq
)

// ex: postgres://user:password@localhost:57031/app?sslmode=disable
func main() {
	// Set up the database connection string.
	const connStr = "host=localhost port=57031 user=app password=password dbname=app sslmode=disable"

	// Open a database connection.
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer db.Close()

	// Check the database connection.
	if err = db.Ping(); err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	// Execute the query to get the current time in AEST.
	var currentTime string
	err = db.QueryRow("SELECT NOW() AT TIME ZONE 'UTC' AT TIME ZONE 'Australia/Sydney'").Scan(&currentTime)
	if err != nil {
		log.Fatalf("Error executing query: %v", err)
	}

	// Print the result.
	fmt.Println("Current time in AEST:", currentTime)
}
