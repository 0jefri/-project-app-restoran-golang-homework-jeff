package repository

import (
	"database/sql"

	"github.com/project-aplikasi-restorant/model"
)

type OrderRepository struct {
	DB *sql.DB
}

func (repo *OrderRepository) CreateOrder(order model.Order) error {
	query := `INSERT INTO orders (user_id, menu_id, status, discount) VALUES ($1, $2, $3, $4)`
	_, err := repo.DB.Exec(query, order.UserID, order.MenuID, order.Status, order.Discount)
	return err
}
