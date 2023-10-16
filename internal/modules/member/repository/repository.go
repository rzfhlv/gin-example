package repository

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/rzfhlv/gin-example/internal/modules/member/model"
)

type IRepository interface {
	Create(ctx context.Context, member model.Member) (result sql.Result, err error)
	Get(ctx context.Context) (result []model.Member, err error)
	GetByID(ctx context.Context, id int64) (result model.Member, err error)
	Update(ctx context.Context) (err error)
	Delete(ctx context.Context) (err error)
}

type Repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) IRepository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Create(ctx context.Context, member model.Member) (result sql.Result, err error) {
	result, err = r.db.Exec(CreateMemberQuery, member.FirstName, member.LastName, member.Email, member.Password, member.CreatedAt)
	return
}

func (r *Repository) Get(ctx context.Context) (result []model.Member, err error) {
	err = r.db.Select(&result, GetMemberQuery)
	return
}

func (r *Repository) GetByID(ctx context.Context, id int64) (result model.Member, err error) {
	err = r.db.Select(&result, GetMemberByIDQuery, id)
	return
}

func (r *Repository) Update(ctx context.Context) (err error) {
	return
}

func (r *Repository) Delete(ctx context.Context) (err error) {
	return
}
