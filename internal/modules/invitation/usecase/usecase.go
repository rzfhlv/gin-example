package usecase

import (
	"context"

	"github.com/rzfhlv/gin-example/internal/modules/invitation/model"
	"github.com/rzfhlv/gin-example/internal/modules/invitation/repository"
)

type IUsecase interface {
	Create(ctx context.Context, invitation model.Invitation) (result model.Invitation, err error)
	Get(ctx context.Context) (result []model.Invitation, err error)
	GetByID(ctx context.Context, id int64) (result model.Invitation, err error)
	Update(ctx context.Context, invitation model.Invitation, id int64) (result model.Invitation, err error)
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

func (u *Usecase) Create(ctx context.Context, invitation model.Invitation) (result model.Invitation, err error) {
	data, err := u.repo.Create(ctx, invitation)
	if err != nil {
		return
	}

	invitation.ID, err = data.LastInsertId()
	if err != nil {
		return
	}

	result = invitation
	return
}

func (u *Usecase) Get(ctx context.Context) (result []model.Invitation, err error) {
	result, err = u.repo.Get(ctx)
	return
}

func (u *Usecase) GetByID(ctx context.Context, id int64) (result model.Invitation, err error) {
	result, err = u.repo.GetByID(ctx, id)
	return
}

func (u *Usecase) Update(ctx context.Context, invitation model.Invitation, id int64) (result model.Invitation, err error) {
	data, err := u.repo.Update(ctx, invitation, id)
	if err != nil {
		return
	}

	invitation.ID, err = data.LastInsertId()
	if err != nil {
		return
	}

	if invitation.Status == "accept" {
		attendee := model.Attendee{}
		attendee.MemberID = invitation.MemberID
		attendee.GatheringID = invitation.GatheringID
		err = u.repo.CreateAttendee(ctx, attendee)
		if err != nil {
			return
		}
	}

	result = invitation
	return
}

func (u *Usecase) Delete(ctx context.Context) (err error) {
	return
}
