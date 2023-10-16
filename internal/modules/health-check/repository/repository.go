package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type IRepository interface {
	Ping(ctx context.Context) (err error)
}

type Repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) IRepository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Ping(ctx context.Context) (err error) {
	err = r.db.Ping()
	return
}
