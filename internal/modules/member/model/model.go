package model

import "time"

type Member struct {
	ID        int64     `json:"id,omitempty" db:"id"`
	FirstName string    `json:"first_name" db:"first_name" binding:"required"`
	LastName  string    `json:"last_name" db:"last_name" binding:"required"`
	Email     string    `json:"email" db:"email" binding:"required"`
	Password  string    `json:"password,omitempty" binding:"required"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
