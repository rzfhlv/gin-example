package model

import "time"

type Register struct {
	ID        int64     `json:"id" db:"id"`
	Username  string    `json:"username" db:"username" binding:"required"`
	Email     string    `json:"email" db:"email" binding:"required"`
	Password  string    `json:"password" db:"password" binding:"required"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type User struct {
	ID        int64     `json:"id" db:"id"`
	Username  string    `json:"username" db:"username"`
	Email     string    `json:"email" db:"email"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type JWT struct {
	Token   string `json:"token"`
	Expired string `json:"expired"`
}

type Login struct {
	Username string `json:"username" db:"username" binding:"required"`
	Password string `json:"password" db:"password" binding:"required"`
}
