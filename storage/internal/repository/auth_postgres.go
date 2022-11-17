package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"storage/internal/model"
)

const (
	UserTable = "users"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{
		db: db,
	}
}

func (r *AuthPostgres) Create(u model.User) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) VALUES ($1, $2, $3) RETURNING id", UserTable)

	row := r.db.QueryRow(query, u.Name, u.Username, u.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (model.User, error) {
	var user model.User
	err := r.db.Get(&user, "select id from users where username = $1 and password_hash=$2", username, password)

	return user, err
}
