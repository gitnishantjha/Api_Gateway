package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gitnishantjha/Api_Gateway/product_service/internal/models"
	"github.com/gitnishantjha/Api_Gateway/product_service/internal/store"
)

// ProductHandler holds the store
type ProductHandler struct {
	Store *store.ProductStore
}

// NewProductHandler creates a new ProductHandler
func NewProductHandler(s *store.ProductStore) *ProductHandler {
	return &ProductHandler{Store: s}
}

// ServeHTTP routes requests to the correct method handler
func (h *ProductHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getProducts(w, r)
	case http.MethodPost:
		h.addProduct(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *ProductHandler) getProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.Store.GetProducts()
	if err != nil {
		// This is proper error handling!
		log.Printf("Error getting products: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func (h *ProductHandler) addProduct(w http.ResponseWriter, r *http.Request) {
	var p models.Product
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	id, err := h.Store.CreateProduct(p)
	if err != nil {
		log.Printf("Error creating product: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	p.ID = id // Set the ID from the database

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(p)
	log.Printf("Added product: %s (ID: %d)", p.Name, p.ID)
}
