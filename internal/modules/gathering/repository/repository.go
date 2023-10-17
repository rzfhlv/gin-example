package repository

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/rzfhlv/gin-example/internal/modules/gathering/model"
	"github.com/rzfhlv/gin-example/pkg/param"
)

type IRepository interface {
	Create(ctx context.Context, gathering model.Gathering) (result sql.Result, err error)
	Get(ctx context.Context, param param.Param) (gatherings []model.Gathering, err error)
	GetByID(ctx context.Context, id int64) (gathering model.Gathering, err error)
	Update(ctx context.Context) (err error)
	Delete(ctx context.Context) (err error)
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

func (r *Repository) Create(ctx context.Context, gathering model.Gathering) (result sql.Result, err error) {
	result, err = r.db.Exec(CreateGatheringQuery,
		gathering.Creator, gathering.MemberID, gathering.Type,
		gathering.Name, gathering.Location, gathering.ScheduleAtDB)
	return
}

func (r *Repository) Get(ctx context.Context, param param.Param) (gatherings []model.Gathering, err error) {
	err = r.db.Select(&gatherings, GetGatheringQuery, param.Limit, param.CalculateOffset())
	return
}

func (r *Repository) GetByID(ctx context.Context, id int64) (gathering model.Gathering, err error) {
	err = r.db.Get(&gathering, GetGatheringByIDQuery, id)
	return
}

func (r *Repository) Update(ctx context.Context) (err error) {
	return
}

func (r *Repository) Delete(ctx context.Context) (err error) {
	return
}

func (r *Repository) Count(ctx context.Context) (total int64, err error) {
	err = r.db.Get(&total, CountGatheringQuery)
	return
}
