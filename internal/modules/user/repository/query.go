package repository

var (
	RegisterUserQuery = `INSERT INTO users
		(username, email, password, created_at)
		VALUES (?, ?, ?, ?);`
	LoginUserQuery = `SELECT id, username,
		email, password, created_at
		FROM users WHERE username = ?;`
	GetUserQuery = `SELECT id, username,
		email, created_at
		FROM users ORDER BY id DESC LIMIT ? OFFSET ?;`
	GetUserByIDQuery = `SELECT id, username
		email, created_at
		FROM users WHERE id = ?;`
	CountUserQuery = `SELECT count(*)
		FROM users;`
)
