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
)
