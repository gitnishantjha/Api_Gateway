package main

import (
	"log"
	"net/http"

	repository "github.com/gitnishantjha/Api_Gateway/order_service/internal/database"
	"github.com/gitnishantjha/Api_Gateway/order_service/internal/handlers"
	"github.com/gitnishantjha/Api_Gateway/order_service/internal/store"
	_ "github.com/lib/pq" // PostgreSQL driver
)

func main() {
	// The port for this service is hardcoded in your gateway config
	port := "9003"

	// The same PostgreSQL connection string
	connStr := "postgres://myuser:mysecretpassword@localhost:5432/productdb?sslmode=disable"

	// 1. Initialize the database (this will create the 'orders' table)
	db := repository.InitDB(connStr)
	defer db.Close()

	// 2. Create the store
	orderStore := store.NewOrderStore(db)

	// 3. Create the handlers
	orderHandler := handlers.NewOrderHandler(orderStore)

	// 4. Register the handler
	http.Handle("/", orderHandler)

	log.Printf("Starting Order Service on port %s (connected to PostgreSQL)", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
