package database

import (
	"database/sql"
	"fmt"
)

var DB *sql.DB

// InitDB Initializes the DB connection in a function
func InitDB() (*sql.DB, error) {
	connStr := "postgres://food_delivery_user:6996891@localhost/fooddelivery?sslmode=disable" // connection string
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("error connecting to database", err)
		return nil, err
	}

	// Verify the connection
	if err = db.Ping(); err != nil {
		fmt.Println("error pinging database", err)
		return nil, err
	}

	fmt.Println("Database connection established")
	return db, nil
}
