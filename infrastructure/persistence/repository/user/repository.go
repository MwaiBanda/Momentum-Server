package user

import (
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func GetInstance(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}
