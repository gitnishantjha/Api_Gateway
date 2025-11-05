package main

import (
	"flag"
	"log"
	"net/http"

	repository "github.com/gitnishantjha/Api_Gateway/product_service/internal/database"
	"github.com/gitnishantjha/Api_Gateway/product_service/internal/handlers"
	"github.com/gitnishantjha/Api_Gateway/product_service/internal/store"
	_ "github.com/lib/pq" // PostgreSQL driver
)

func main() {
	port := flag.String("port", "9001", "Port to listen on")
	flag.Parse()
	connStr := "postgres://myuser:mysecretpassword@localhost:5432/productdb?sslmode=disable"

	// 1. Initialize the database
	db := repository.InitDB(connStr)
	defer db.Close()

	// 2. Create the store
	productStore := store.NewProductStore(db)

	// 3. Create the handlers
	productHandler := handlers.NewProductHandler(productStore)

	http.Handle("/", productHandler)

	log.Printf("Starting Product Service on port %s (connected to PostgreSQL)", *port)
	if err := http.ListenAndServe(":"+*port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
