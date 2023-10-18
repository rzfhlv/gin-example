package repository

var (
	CreateMemberQuery = `INSERT INTO members
		(first_name, last_name, email, password, created_at)
		VALUES (?, ?, ?, ?, ?);`
	GetMemberQuery = `SELECT id, first_name,
		last_name, email, created_at
		FROM members ORDER BY id DESC LIMIT ? OFFSET ?;`
	GetMemberByIDQuery = `SELECT id, first_name,
		last_name, email
		FROM members WHERE id = ?;`
	CountMemberQuery = `SELECT count(*)
		FROM members;`
)
