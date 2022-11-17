package repository

import (
	"github.com/jmoiron/sqlx"
	"storage/internal/model"
)

type Authorization interface {
	Create(u model.User) (int, error)
	GetUser(username, password string) (model.User, error)
}

type FileStorage interface {
	UploadDB(fileName, Key string) error
	GetKey(key string) (string, error)
	DeleteFile(key string) error
}

type Repository struct {
	Authorization
	FileStorage
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		FileStorage:   NewFileClient(db),
	}
}
