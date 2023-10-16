package usecase

import (
	"context"

	"github.com/rzfhlv/gin-example/internal/modules/member/model"
	"github.com/rzfhlv/gin-example/internal/modules/member/repository"
	"github.com/rzfhlv/gin-example/pkg/hasher"
)

type IUsecase interface {
	Create(ctx context.Context, member model.Member) (result model.Member, err error)
	Get(ctx context.Context) (result []model.Member, err error)
	GetByID(ctx context.Context, id int64) (result model.Member, err error)
	Update(ctx context.Context) (err error)
	Delete(ctx context.Context) (err error)
}

type Usecase struct {
	repo   repository.IRepository
	hasher hasher.HashPassword
}

func New(repo repository.IRepository, hasher hasher.HashPassword) IUsecase {
	return &Usecase{
		repo:   repo,
		hasher: hasher,
	}
}

func (u *Usecase) Create(ctx context.Context, member model.Member) (result model.Member, err error) {
	hashPassword, err := u.hasher.HashedPassword(member.Password)
	if err != nil {
		return
	}
	member.Password = hashPassword

	data, err := u.repo.Create(ctx, member)
	if err != nil {
		return
	}

	member.ID, err = data.LastInsertId()
	if err != nil {
		return
	}

	result = member
	return
}

func (u *Usecase) Get(ctx context.Context) (result []model.Member, err error) {
	result, err = u.repo.Get(ctx)
	return
}

func (u *Usecase) GetByID(ctx context.Context, id int64) (result model.Member, err error) {
	result, err = u.repo.GetByID(ctx, id)
	return
}

func (u *Usecase) Update(ctx context.Context) (err error) {
	return
}

func (u *Usecase) Delete(ctx context.Context) (err error) {
	return
}
