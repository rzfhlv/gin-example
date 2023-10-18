package model

import (
	"time"
)

type Gathering struct {
	ID           int64     `json:"id,omitempty" db:"id"`
	Creator      string    `json:"creator" db:"creator" binding:"required"`
	Type         string    `json:"type" db:"type" binding:"required"`
	Name         string    `json:"name" db:"name" binding:"required"`
	Location     string    `json:"location" db:"location" binding:"required"`
	ScheduleAt   string    `json:"schedule_at" db:"schedule_at" binding:"required"`
	MemberID     int64     `json:"-" db:"member_id"`
	ScheduleAtDB time.Time `json:"-" db:"schedule_at"`
}

type GatheringDetail struct {
	Gathering
	Attendees []Attendee `json:"attendees"`
}

type Attendee struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Status    string `json:"status"`
}
