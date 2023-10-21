package repository

import (
	"context"
	"database/sql"

	"github.com/rzfhlv/gin-example/internal/modules/user/model"
	"github.com/rzfhlv/gin-example/pkg/param"
)

func (r *Repository) Register(ctx context.Context, register model.Register) (result sql.Result, err error) {
	result, err = r.db.Exec(RegisterUserQuery, register.Username, register.Email, register.Password, register.CreatedAt)
	return
}

func (r *Repository) Login(ctx context.Context, login model.Login) (register model.Register, err error) {
	err = r.db.Get(&register, LoginUserQuery, login.Username)
	return
}

func (r *Repository) GetAll(ctx context.Context, param param.Param) (users []model.User, err error) {
	err = r.db.Select(&users, GetUserQuery, param.Limit, param.CalculateOffset())
	return
}

func (r *Repository) GetByID(ctx context.Context, id int64) (user model.User, err error) {
	err = r.db.Get(&user, GetUserByIDQuery, id)
	return
}

func (r *Repository) Count(ctx context.Context) (total int64, err error) {
	err = r.db.Get(&total, CountUserQuery)
	return
}
