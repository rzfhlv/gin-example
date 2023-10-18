package repository

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/rzfhlv/gin-example/internal/modules/invitation/model"
	"github.com/rzfhlv/gin-example/pkg/param"
)

type IRepository interface {
	Create(ctx context.Context, invitation model.Invitation) (result sql.Result, err error)
	Get(ctx context.Context, param param.Param) (invitations []model.Invitation, err error)
	GetByID(ctx context.Context, id int64) (invitation model.Invitation, err error)
	Update(ctx context.Context, invitation model.Invitation, id int64) (result sql.Result, err error)
	CreateAttendee(ctx context.Context, attendee model.Attendee) (err error)
	Count(ctx context.Context) (total int64, err error)
	GetByMemberID(ctx context.Context, memberID int64) (invitations []model.InvitationDetail, err error)
}

type Repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) IRepository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Create(ctx context.Context, invitation model.Invitation) (result sql.Result, err error) {
	result, err = r.db.Exec(CreateInvitationQuery, invitation.MemberID, invitation.GatheringID, invitation.Status)
	return
}

func (r *Repository) Get(ctx context.Context, param param.Param) (invitations []model.Invitation, err error) {
	err = r.db.Select(&invitations, GetInvitationQuery, param.Limit, param.CalculateOffset())
	return
}

func (r *Repository) GetByID(ctx context.Context, id int64) (invitation model.Invitation, err error) {
	err = r.db.Get(&invitation, GetInvitationByIDQuery, id)
	return
}

func (r *Repository) Update(ctx context.Context, invitation model.Invitation, id int64) (result sql.Result, err error) {
	result, err = r.db.Exec(UpdateInvitationQuery, invitation.Status, id)
	return
}

func (r *Repository) CreateAttendee(ctx context.Context, attendee model.Attendee) (err error) {
	_, err = r.db.Exec(CreateAttendeeQuery, attendee.MemberID, attendee.GatheringID)
	return
}

func (r *Repository) Count(ctx context.Context) (total int64, err error) {
	err = r.db.Get(&total, CountInvitationQuery)
	return
}

func (r *Repository) GetByMemberID(ctx context.Context, memberID int64) (invitations []model.InvitationDetail, err error) {
	rows, err := r.db.Query(GetInvitationByMemberIDQuery, memberID)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var invitation = model.InvitationDetail{}
		err = rows.Scan(&invitation.ID, &invitation.MemberID, &invitation.GatheringID,
			&invitation.Status, &invitation.Gathering.ID, &invitation.Gathering.Creator,
			&invitation.Gathering.Type, &invitation.Gathering.Name,
			&invitation.Gathering.Location, &invitation.Gathering.ScheduleAt)
		if err != nil {
			return
		}
		invitations = append(invitations, invitation)
	}
	return
}
