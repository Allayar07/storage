package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DbName   string
	SSlmode  string
}

func NewPostgres(cfg Config) (*sqlx.DB, error) {
	str := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DbName, cfg.Password, cfg.SSlmode)
	db, err := sqlx.Open("postgres", str)
	if err != nil {
		return nil, err
	}

	return db, nil
}