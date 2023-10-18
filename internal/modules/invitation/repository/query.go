package repository

var (
	CreateInvitationQuery = `INSERT INTO invitations
		(member_id, gathering_id, status)
		VALUES (?, ?, ?);`
	GetInvitationQuery = `SELECT id, member_id,
		gathering_id, status
		FROM invitations ORDER BY id DESC LIMIT ? OFFSET ?;`
	GetInvitationByIDQuery = `SELECT id, member_id,
		gathering_id, status
		FROM invitations WHERE id = ?;`
	UpdateInvitationQuery = `UPDATE invitations
		SET status = ?
		WHERE id = ?;`
	CreateAttendeeQuery = `INSERT INTO attendee
		(member_id, gathering_id)
		VALUES (?, ?);`
	CountInvitationQuery = `SELECT count(*)
		FROM invitations;`
	GetInvitationByMemberIDQuery = `SELECT i.id as iid, i.member_id,
		i.gathering_id, i.status, g.id as gid, g.creator,
		g.type, g.name, g.location, g.schedule_at 
		FROM invitations i
		LEFT JOIN gatherings g ON i.gathering_id = g.id
		WHERE i.member_id = ?`
)
