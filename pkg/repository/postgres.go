package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	usersTable = "users"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

// postgres://{user}:{password}@{hostname}:{port}/{database-name}?sslmode=disable
func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	conString := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SSLMode)
	db, err := sqlx.Connect("postgres", conString)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return db, nil
}
