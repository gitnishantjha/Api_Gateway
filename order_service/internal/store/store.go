package store

import (
	"database/sql"

	"github.com/gitnishantjha/Api_Gateway/order_service/internal/models"
)

type OrderStore struct {
	DB *sql.DB
}

func NewOrderStore(db *sql.DB) *OrderStore {
	return &OrderStore{
		DB: db,
	}
}

func (s *OrderStore) GetOrders() ([]models.Order, error) {
	rows, err := s.DB.Query("select id, product_id,quantity from orders")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []models.Order

	for rows.Next() {
		var o models.Order
		if err := rows.Scan(&o.ID, &o.ProductID, &o.Quantity); err != nil {
			return nil, err
		}
		orders = append(orders, o)
	}
	return orders, nil
}

func (s *OrderStore) CreateOrder(o models.Order) (int64, error) {
	var id int64
	query := "INSERT INTO orders (product_id, quantity) VALUES ($1, $2) RETURNING id"
	err := s.DB.QueryRow(query, o.ProductID, o.Quantity).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
