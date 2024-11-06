package repository

import (
	"database/sql"

	"github.com/project-aplikasi-restorant/model"
)

type UserRepository struct {
	DB *sql.DB
}

func (repo *UserRepository) GetUserByUsername(username string) (model.User, error) {
	var user model.User
	query := "SELECT id, username, password, role FROM users WHERE username=$1"
	err := repo.DB.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Password, &user.Role)
	return user, err
}
