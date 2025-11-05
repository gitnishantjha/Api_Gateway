package repository

import (
	"database/sql"
	"log"
)

func InitDB(connStr string) *sql.DB {
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	// Create the orders table
	//  add a foreign key to the products table to ensure data integrity
	createTableSQL := `CREATE TABLE IF NOT EXISTS orders(
	id SERIAL PRIMARY KEY,
	product_id INTEGER,
	quantity INTEGER,
	CONSTRAINT fk_product
			FOREIGN KEY(product_id) 
			REFERENCES products(id)
	);`

	_, err = db.Exec(createTableSQL)

	if err != nil {
		log.Fatalf("Failed to create orders table: %v", err)
	}

	log.Println("PostgreSQL database initialized and orders table created successfully.")
	return db
}
