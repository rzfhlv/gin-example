package usecase

import (
	"context"

	"github.com/rzfhlv/gin-example/internal/modules/invitation/model"
	"github.com/rzfhlv/gin-example/internal/modules/invitation/repository"
	"github.com/rzfhlv/gin-example/pkg/param"
)

type IUsecase interface {
	Create(ctx context.Context, invitationPayload model.Invitation) (invitation model.Invitation, err error)
	Get(ctx context.Context, param param.Param) (invitations []model.Invitation, total int64, err error)
	GetByID(ctx context.Context, id int64) (invitation model.Invitation, err error)
	Update(ctx context.Context, invitationPayload model.Invitation, id int64) (invitation model.Invitation, err error)
	Delete(ctx context.Context) (err error)
	GetByMemberID(ctx context.Context, memberID int64) (invitations []model.InvitationDetail, err error)
}

type Usecase struct {
	repo repository.IRepository
}

func New(repo repository.IRepository) IUsecase {
	return &Usecase{
		repo: repo,
	}
}

func (u *Usecase) Create(ctx context.Context, invitationPayload model.Invitation) (invitation model.Invitation, err error) {
	result, err := u.repo.Create(ctx, invitationPayload)
	if err != nil {
		return
	}

	invitationPayload.ID, err = result.LastInsertId()
	if err != nil {
		return
	}

	attendee := model.Attendee{}
	attendee.MemberID = invitationPayload.MemberID
	attendee.GatheringID = invitationPayload.GatheringID
	err = u.repo.CreateAttendee(ctx, attendee)
	if err != nil {
		return
	}

	invitation = invitationPayload
	return
}

func (u *Usecase) Get(ctx context.Context, param param.Param) (invitations []model.Invitation, total int64, err error) {
	invitations, err = u.repo.Get(ctx, param)
	if err != nil {
		return
	}

	if len(invitations) < 1 {
		invitations = []model.Invitation{}
	}
	total, err = u.repo.Count(ctx)
	return
}

func (u *Usecase) GetByID(ctx context.Context, id int64) (invitation model.Invitation, err error) {
	invitation, err = u.repo.GetByID(ctx, id)
	return
}

func (u *Usecase) Update(ctx context.Context, invitationPayload model.Invitation, id int64) (invitation model.Invitation, err error) {
	_, err = u.repo.Update(ctx, invitationPayload, id)
	if err != nil {
		return
	}

	invitationPayload.ID = id
	invitation = invitationPayload
	return
}

func (u *Usecase) Delete(ctx context.Context) (err error) {
	return
}

func (u *Usecase) GetByMemberID(ctx context.Context, memberID int64) (invitations []model.InvitationDetail, err error) {
	invitations, err = u.repo.GetByMemberID(ctx, memberID)
	return
}
