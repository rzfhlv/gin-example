package repository

var (
	CreateInvitationQuery = `INSERT INTO invitations
		(member_id, gathering_id, status)
		VALUES (?, ?, ?);`
	GetInvitationQuery = `SELECT id, member_id,
		gathering_id, status
		FROM invitations;`
	GetInvitationByIDQuery = `SELECT id, member_id,
		gathering_id, status
		FROM invitations WHERE id = ?;`
	UpdateInvitationQuery = `UPDATE invitations
		SET status = ?
		WHERE id = ?;`
	CreateAttendeeQuery = `INSERT INTO attendee
		(member_id, gathering_id)
		VALUES (?, ?);`
)
