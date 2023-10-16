package model

type Invitation struct {
	ID          int64  `json:"id" db:"id"`
	MemberID    int64  `json:"member_id" db:"member_id" binding:"required"`
	GatheringID int64  `json:"gathering_id" db:"gathering_id" binding:"required"`
	Status      string `json:"status" db:"status" binding:"required"`
}

type Attendee struct {
	MemberID    int64 `json:"member_id" db:"member_id"`
	GatheringID int64 `json:"gathering_id" db:"gathering_id"`
}
