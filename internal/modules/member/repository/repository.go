package repository

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/rzfhlv/gin-example/internal/modules/member/model"
	"github.com/rzfhlv/gin-example/pkg/param"
)

type IRepository interface {
	Create(ctx context.Context, member model.Member) (result sql.Result, err error)
	Get(ctx context.Context, param param.Param) (members []model.Member, err error)
	GetByID(ctx context.Context, id int64) (member model.Member, err error)
	Count(ctx context.Context) (total int64, err error)
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

func (r *Repository) Get(ctx context.Context, param param.Param) (members []model.Member, err error) {
	err = r.db.Select(&members, GetMemberQuery, param.Limit, param.CalculateOffset())
	return
}

func (r *Repository) GetByID(ctx context.Context, id int64) (member model.Member, err error) {
	err = r.db.Get(&member, GetMemberByIDQuery, id)
	return
}

func (r *Repository) Count(ctx context.Context) (total int64, err error) {
	err = r.db.Get(&total, CountMemberQuery)
	return
}
