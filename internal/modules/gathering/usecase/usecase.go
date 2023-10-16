package usecase

import (
	"context"

	"github.com/rzfhlv/gin-example/internal/modules/gathering/model"
	"github.com/rzfhlv/gin-example/internal/modules/gathering/repository"
)

type IUsecase interface {
	Create(ctx context.Context, gathering model.Gathering) (result model.Gathering, err error)
	Get(ctx context.Context) (result []model.Gathering, err error)
	GetByID(ctx context.Context, id int64) (result model.Gathering, err error)
	Update(ctx context.Context) (err error)
	Delete(ctx context.Context) (err error)
}

type Usecase struct {
	repo repository.IRepository
}

func New(repo repository.IRepository) IUsecase {
	return &Usecase{
		repo: repo,
	}
}

func (u *Usecase) Create(ctx context.Context, gathering model.Gathering) (result model.Gathering, err error) {
	data, err := u.repo.Create(ctx, gathering)
	if err != nil {
		return
	}

	gathering.ID, err = data.LastInsertId()
	if err != nil {
		return
	}

	result = gathering
	return
}

func (u *Usecase) Get(ctx context.Context) (result []model.Gathering, err error) {
	result, err = u.repo.Get(ctx)
	return
}

func (u *Usecase) GetByID(ctx context.Context, id int64) (result model.Gathering, err error) {
	result, err = u.repo.GetByID(ctx, id)
	return
}

func (u *Usecase) Update(ctx context.Context) (err error) {
	return
}

func (u *Usecase) Delete(ctx context.Context) (err error) {
	return
}
