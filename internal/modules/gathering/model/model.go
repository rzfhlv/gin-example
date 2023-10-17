package model

import "time"

type Gathering struct {
	ID           int64     `json:"id,omitempty" db:"id"`
	Creator      string    `json:"creator" db:"creator" binding:"required"`
	Type         string    `json:"type" db:"type" binding:"required"`
	Name         string    `json:"name" db:"name" binding:"required"`
	Location     string    `json:"location" db:"location" binding:"required"`
	ScheduleAt   string    `json:"schedule_at" db:"schedule_at" binding:"required"`
	MemberID     int64     `json:"member_id" db:"member_id"`
	ScheduleAtDB time.Time `json:"-" db:"schedule_at"`
}
