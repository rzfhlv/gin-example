package usecase

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/rzfhlv/gin-example/internal/modules/user/model"
	"github.com/rzfhlv/gin-example/internal/modules/user/repository"
	"github.com/rzfhlv/gin-example/pkg/hasher"
	pJwt "github.com/rzfhlv/gin-example/pkg/jwt"
	"github.com/rzfhlv/gin-example/pkg/param"
	"golang.org/x/crypto/bcrypt"
)

type IUsecase interface {
	Register(ctx context.Context, register model.Register) (jwt model.JWT, err error)
	Login(ctx context.Context, login model.Login) (jwt model.JWT, err error)
	Logout(ctx context.Context, username string) (err error)
	GetAll(ctx context.Context, param param.Param) (users []model.User, total int64, err error)
	GetByID(ctx context.Context, id int64) (user model.User, err error)
}

type Usecase struct {
	repo    repository.IRepository
	hasher  hasher.HashPassword
	jwtImpl pJwt.JWTInterface
}

func New(repo repository.IRepository, hasher hasher.HashPassword, jwtImpl pJwt.JWTInterface) IUsecase {
	return &Usecase{
		repo:    repo,
		hasher:  hasher,
		jwtImpl: jwtImpl,
	}
}

func (u *Usecase) Register(ctx context.Context, register model.Register) (jwt model.JWT, err error) {
	hashPassword, err := u.hasher.HashedPassword(register.Password)
	if err != nil {
		return
	}
	register.Password = hashPassword

	result, err := u.repo.Register(ctx, register)
	if err != nil {
		return
	}

	register.ID, err = result.LastInsertId()
	if err != nil {
		return
	}

	token, err := u.jwtImpl.Generate(register.ID, register.Username, register.Email)
	if err != nil {
		return
	}

	err = u.repo.Set(ctx, register.Username, token, time.Duration(1*time.Hour))
	if err != nil {
		return
	}

	jwt.Token = token
	jwt.Expired = fmt.Sprintf("%s Hour", os.Getenv("JWT_EXPIRED"))
	return
}

func (u *Usecase) Login(ctx context.Context, login model.Login) (jwt model.JWT, err error) {
	register, err := u.repo.Login(ctx, login)
	if err != nil {
		return
	}

	err = u.hasher.VerifyPassword(register.Password, login.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return
	}

	token, err := u.jwtImpl.Generate(register.ID, register.Username, register.Email)
	if err != nil {
		return
	}

	err = u.repo.Set(ctx, register.Username, token, time.Duration(1*time.Hour))
	if err != nil {
		return
	}

	jwt.Token = token
	jwt.Expired = fmt.Sprintf("%s Hour", os.Getenv("JWT_EXPIRED"))
	return
}

func (u *Usecase) Logout(ctx context.Context, username string) (err error) {
	err = u.repo.Del(ctx, username)
	return
}

func (u *Usecase) GetAll(ctx context.Context, param param.Param) (users []model.User, total int64, err error) {
	users, err = u.repo.GetAll(ctx, param)
	if err != nil {
		return
	}

	if len(users) < 1 {
		users = []model.User{}
	}
	total, err = u.repo.Count(ctx)
	return
}
func (u *Usecase) GetByID(ctx context.Context, id int64) (user model.User, err error) {
	user, err = u.repo.GetByID(ctx, id)
	return
}
