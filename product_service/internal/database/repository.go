package repository

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // The PostgreSQL driver
)

func InitDB(connStr string) *sql.DB {

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	// Create the products table with Postgres-specific syntax
	// SERIAL PRIMARY KEY is the auto-incrementing integer in Postgres
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS products (
		id SERIAL PRIMARY KEY,
		name TEXT,
		qty INTEGER
	);`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}

	log.Println("PostgreSQL database initialized and table created successfully.")
	return db
}
