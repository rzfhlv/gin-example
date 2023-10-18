package usecase

import (
	"context"
	"time"

	"github.com/rzfhlv/gin-example/internal/modules/gathering/model"
	"github.com/rzfhlv/gin-example/internal/modules/gathering/repository"
	"github.com/rzfhlv/gin-example/pkg/param"
)

type IUsecase interface {
	Create(ctx context.Context, gathering model.Gathering) (result model.Gathering, err error)
	Get(ctx context.Context, param param.Param) (gatherings []model.Gathering, total int64, err error)
	GetByID(ctx context.Context, id int64) (gathering model.Gathering, err error)
	Update(ctx context.Context) (err error)
	Delete(ctx context.Context) (err error)
	GetDetailByID(ctx context.Context, id int64) (gathering model.GatheringDetail, err error)
}

type Usecase struct {
	repo repository.IRepository
}

func New(repo repository.IRepository) IUsecase {
	return &Usecase{
		repo: repo,
	}
}

func (u *Usecase) Create(ctx context.Context, gatheringPayload model.Gathering) (gathering model.Gathering, err error) {
	scheduleAt, err := time.Parse("2006-01-02 03:04:05", gatheringPayload.ScheduleAt)
	if err != nil {
		return
	}
	gatheringPayload.ScheduleAtDB = scheduleAt
	result, err := u.repo.Create(ctx, gatheringPayload)
	if err != nil {
		return
	}

	gatheringPayload.ID, err = result.LastInsertId()
	if err != nil {
		return
	}

	gathering = gatheringPayload
	return
}

func (u *Usecase) Get(ctx context.Context, param param.Param) (gatherings []model.Gathering, total int64, err error) {
	gatherings, err = u.repo.Get(ctx, param)
	if err != nil {
		return
	}

	if len(gatherings) < 1 {
		gatherings = []model.Gathering{}
	}
	total, err = u.repo.Count(ctx)
	return
}

func (u *Usecase) GetByID(ctx context.Context, id int64) (gathering model.Gathering, err error) {
	gathering, err = u.repo.GetByID(ctx, id)
	return
}

func (u *Usecase) Update(ctx context.Context) (err error) {
	return
}

func (u *Usecase) Delete(ctx context.Context) (err error) {
	return
}

func (u *Usecase) GetDetailByID(ctx context.Context, id int64) (gathering model.GatheringDetail, err error) {
	gatheringByID, err := u.repo.GetByID(ctx, id)
	if err != nil {
		return
	}

	gathering, err = u.repo.GetDetailByID(ctx, id)
	gathering.Gathering = gatheringByID
	return
}
