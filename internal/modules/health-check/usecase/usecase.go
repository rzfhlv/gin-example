package usecase

import (
	"context"

	"github.com/rzfhlv/gin-example/internal/modules/health-check/repository"
)

type IUsecase interface {
	Ping(ctx context.Context) (err error)
}

type Usecase struct {
	repo repository.IRepository
}

func New(repo repository.IRepository) IUsecase {
	return &Usecase{
		repo: repo,
	}
}

func (u *Usecase) Ping(ctx context.Context) (err error) {
	err = u.repo.Ping(ctx)
	return
}
