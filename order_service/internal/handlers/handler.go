package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gitnishantjha/Api_Gateway/order_service/internal/models"
	"github.com/gitnishantjha/Api_Gateway/order_service/internal/store"
)

type OrderHandler struct {
	Store *store.OrderStore
}

func NewOrderHandler(s *store.OrderStore) *OrderHandler {
	return &OrderHandler{Store: s}
}
func (h *OrderHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getOrders(w, r)
	case http.MethodPost:
		h.addOrder(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
func (h *OrderHandler) getOrders(w http.ResponseWriter, _ *http.Request) {
	orders, err := h.Store.GetOrders()
	if err != nil {
		log.Printf("Error getting orders: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}

func (h *OrderHandler) addOrder(w http.ResponseWriter, r *http.Request) {
	var o models.Order
	if err := json.NewDecoder(r.Body).Decode(&o); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// NOTE: A service would first check if the product_id exists
	// and if there is enough quantity. We are skipping that for simplicity.

	id, err := h.Store.CreateOrder(o)
	if err != nil {
		log.Printf("Error creating order: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	o.ID = id // Set the ID from the database

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(o)
	log.Printf("Added order: %d", o.ID)
}
