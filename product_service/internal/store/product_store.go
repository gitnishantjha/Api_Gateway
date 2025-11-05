package store

import (
	"database/sql"

	"github.com/gitnishantjha/Api_Gateway/product_service/internal/models"
)

type ProductStore struct {
	DB *sql.DB
}

func NewProductStore(db *sql.DB) *ProductStore {
	return &ProductStore{DB: db}
}
func (s *ProductStore) GetProducts() ([]models.Product, error) {
	rows, err := s.DB.Query("SELECT id, name, qty FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var p models.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Qty); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func (s *ProductStore) CreateProduct(p models.Product) (int64, error) {
	var id int64

	query := "INSERT INTO products (name, qty) VALUES ($1, $2) RETURNING id"

	err := s.DB.QueryRow(query, p.Name, p.Qty).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *ProductStore) GetProductByID(id int64) (*models.Product, error) {
	var p models.Product

	query := "SELECT id, name, qty FROM products WHERE id = $1"
	row := s.DB.QueryRow(query, id)

	err := row.Scan(&p.ID, &p.Name, &p.Qty)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &p, nil
}
