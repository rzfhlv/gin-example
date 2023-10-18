package usecase

import (
	"context"

	"github.com/rzfhlv/gin-example/internal/modules/member/model"
	"github.com/rzfhlv/gin-example/internal/modules/member/repository"
	"github.com/rzfhlv/gin-example/pkg/hasher"
	"github.com/rzfhlv/gin-example/pkg/param"
)

type IUsecase interface {
	Create(ctx context.Context, memberPayload model.Member) (member model.Member, err error)
	Get(ctx context.Context, param param.Param) (members []model.Member, total int64, err error)
	GetByID(ctx context.Context, id int64) (member model.Member, err error)
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

func (u *Usecase) Create(ctx context.Context, memberPayload model.Member) (member model.Member, err error) {
	hashPassword, err := u.hasher.HashedPassword(memberPayload.Password)
	if err != nil {
		return
	}
	memberPayload.Password = hashPassword

	data, err := u.repo.Create(ctx, memberPayload)
	if err != nil {
		return
	}

	memberPayload.ID, err = data.LastInsertId()
	if err != nil {
		return
	}

	member = memberPayload
	return
}

func (u *Usecase) Get(ctx context.Context, param param.Param) (members []model.Member, total int64, err error) {
	members, err = u.repo.Get(ctx, param)
	if err != nil {
		return
	}

	if len(members) < 1 {
		members = []model.Member{}
	}
	total, err = u.repo.Count(ctx)
	return
}

func (u *Usecase) GetByID(ctx context.Context, id int64) (member model.Member, err error) {
	member, err = u.repo.GetByID(ctx, id)
	return
}
