package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	"github.com/rzfhlv/gin-example/internal/modules/user/model"
	"github.com/rzfhlv/gin-example/pkg/param"
)

type IRepository interface {
	Register(ctx context.Context, register model.Register) (result sql.Result, err error)
	Login(ctx context.Context, login model.Login) (register model.Register, err error)
	GetAll(ctx context.Context, param param.Param) (users []model.User, err error)
	GetByID(ctx context.Context, id int64) (user model.User, err error)
	Count(ctx context.Context) (total int64, err error)
	Set(ctx context.Context, key, value string, ttl time.Duration) (err error)
	Get(ctx context.Context, key string) (value string, err error)
	Del(ctx context.Context, key string) (err error)
}

type Repository struct {
	db    *sqlx.DB
	redis *redis.Client
}

func New(db *sqlx.DB, redis *redis.Client) IRepository {
	return &Repository{
		db:    db,
		redis: redis,
	}
}
