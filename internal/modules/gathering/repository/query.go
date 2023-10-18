package repository

var (
	CreateGatheringQuery = `INSERT INTO gatherings
		(creator, member_id, type, name, location, schedule_at)
		VALUES (?, ?, ?, ?, ?, ?);`
	GetGatheringQuery = `SELECT id, creator, member_id, type,
		name, location, schedule_at
		FROM gatherings ORDER BY id DESC LIMIT ? OFFSET ?;`
	GetGatheringByIDQuery = `SELECT id, creator, member_id,
		type, name, location, schedule_at
		FROM gatherings WHERE id = ?;`
	CountGatheringQuery = `SELECT count(*)
		FROM gatherings;`
	GetDetailGatheringByIDQuery = `SELECT m.id, m.first_name,
		m.last_name, m.email, i.status
		FROM members m
		LEFT JOIN attendee a ON m.id = a.member_id
		LEFT JOIN invitations i ON a.gathering_id = i.gathering_id
		AND a.member_id = i.member_id
		WHERE a.gathering_id = ?;`
)
