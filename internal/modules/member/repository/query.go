package repository

var (
	CreateMemberQuery = `INSERT INTO members
		(first_name, last_name, email, password, created_at)
		VALUES (?, ?, ?, ?, ?);`
	GetMemberQuery = `SELECT id, first_name,
		last_name, email, created_at
		FROM members;`
	GetMemberByIDQuery = `SELECT id, first_name,
		last_name, email, created_at
		FROM members WHERE id = ?;`
)
