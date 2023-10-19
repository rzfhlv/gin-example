package repository

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	modelGathering "github.com/rzfhlv/gin-example/internal/modules/gathering/model"
	"github.com/rzfhlv/gin-example/internal/modules/invitation/model"
	"github.com/rzfhlv/gin-example/pkg/param"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	name       string
	args       context.Context
	beforeTest func(s sqlmock.Sqlmock)
	want       error
	wantError  bool
}

var (
	ctx         = context.Background()
	invitations = []model.Invitation{
		{
			ID: 1, MemberID: 1, GatheringID: 1, Status: "accept",
		},
	}
	attendee = model.Attendee{
		MemberID:    1,
		GatheringID: 1,
	}
	gathering = modelGathering.Gathering{
		ID:         1,
		Creator:    "John",
		Type:       "family",
		Name:       "Family Gathering",
		Location:   "Puncak",
		ScheduleAt: "2023-11-10 12:00:00",
	}
	errFoo    = errors.New("foo")
	paramTest = param.Param{
		Limit: 10,
		Page:  1,
	}
)

func TestCreate(t *testing.T) {
	testCase := []testCase{
		{
			name: "Testcase #1: Positive",
			args: ctx,
			beforeTest: func(s sqlmock.Sqlmock) {
				s.ExpectExec(`INSERT INTO invitations
						(member_id, gathering_id, status)
						VALUES (?, ?, ?);`).
					WithArgs(invitations[0].MemberID, invitations[0].GatheringID, invitations[0].Status).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			want:      errFoo,
			wantError: false,
		},
		{
			name: "Testcase #2: Negative",
			args: ctx,
			beforeTest: func(s sqlmock.Sqlmock) {
				s.ExpectExec(`INSERT INTO invitations
						(member_id, gathering_id, status)
						VALUES (?, ?, ?);`).
					WithArgs(invitations[0].MemberID, invitations[0].GatheringID, invitations[0].Status).
					WillReturnError(errFoo)
			},
			want:      errFoo,
			wantError: true,
		},
	}
	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			defer mockDB.Close()

			db := sqlx.NewDb(mockDB, "sqlmock")

			r := &Repository{
				db: db,
			}

			if tt.beforeTest != nil {
				tt.beforeTest(mockSQL)
			}

			result, err := r.Create(tt.args, invitations[0])
			if tt.wantError {
				assert.Error(t, err)
				assert.Empty(t, result)
			} else {
				assert.NoError(t, err)
			}

			if err := mockSQL.ExpectationsWereMet(); err != nil {
				assert.Error(t, err)
			}
		})
	}
}

func TestGet(t *testing.T) {
	testCase := []testCase{
		{
			name: "Testcase #1: Positive",
			args: ctx,
			beforeTest: func(s sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{
					"id", "member_id", "gathering_id", "status",
				}).
					AddRow(invitations[0].ID, invitations[0].MemberID, invitations[0].GatheringID, invitations[0].Status)
				s.ExpectQuery(`SELECT id, member_id, gathering_id, status
						FROM invitations ORDER BY id DESC LIMIT ? OFFSET ?;`).
					WithArgs(paramTest.Limit, paramTest.CalculateOffset()).
					WillReturnRows(rows)
			},
			want:      nil,
			wantError: false,
		},
		{
			name: "Testcase #2: Negative",
			args: ctx,
			beforeTest: func(s sqlmock.Sqlmock) {
				s.ExpectQuery(`SELECT id, member_id, gathering_id, status
						FROM invitations ORDER BY id DESC LIMIT ? OFFSET ?;`).
					WithArgs(paramTest.Limit, paramTest.CalculateOffset()).
					WillReturnError(errFoo)
			},
			want:      errFoo,
			wantError: true,
		},
	}
	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			defer mockDB.Close()

			db := sqlx.NewDb(mockDB, "sqlmock")

			r := &Repository{
				db: db,
			}

			if tt.beforeTest != nil {
				tt.beforeTest(mockSQL)
			}

			invitations, err := r.Get(tt.args, paramTest)
			if tt.wantError {
				assert.Error(t, err)
				assert.Empty(t, invitations)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, invitations)
			}

			if err := mockSQL.ExpectationsWereMet(); err != nil {
				assert.Error(t, err)
			}
		})
	}
}

func TestGetByID(t *testing.T) {
	testCase := []testCase{
		{
			name: "Testcase #1: Positive",
			args: ctx,
			beforeTest: func(s sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{
					"id", "member_id", "gathering_id", "status",
				}).
					AddRow(invitations[0].ID, invitations[0].MemberID, invitations[0].GatheringID, invitations[0].Status)
				s.ExpectQuery(`SELECT id, member_id,
						gathering_id, status
						FROM invitations WHERE id = ?;`).
					WithArgs(invitations[0].ID).
					WillReturnRows(rows)
			},
			want:      nil,
			wantError: false,
		},
		{
			name: "Testcase #2: Negative",
			args: ctx,
			beforeTest: func(s sqlmock.Sqlmock) {
				s.ExpectQuery(`SELECT id, member_id,
						gathering_id, status
						FROM invitations WHERE id = ?;`).
					WithArgs(invitations[0].ID).
					WillReturnError(errFoo)
			},
			want:      errFoo,
			wantError: true,
		},
	}
	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			defer mockDB.Close()

			db := sqlx.NewDb(mockDB, "sqlmock")

			r := &Repository{
				db: db,
			}

			if tt.beforeTest != nil {
				tt.beforeTest(mockSQL)
			}

			invitation, err := r.GetByID(tt.args, invitations[0].ID)
			if tt.wantError {
				assert.Error(t, err)
				assert.Empty(t, invitation)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, invitation)
			}

			if err := mockSQL.ExpectationsWereMet(); err != nil {
				assert.Error(t, err)
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	testCase := []testCase{
		{
			name: "Testcase #1: Positive",
			args: ctx,
			beforeTest: func(s sqlmock.Sqlmock) {
				s.ExpectExec(`UPDATE invitations
						SET status = ?
						WHERE id = ?;`).
					WithArgs(invitations[0].Status, invitations[0].MemberID).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			want:      errFoo,
			wantError: false,
		},
		{
			name: "Testcase #2: Negative",
			args: ctx,
			beforeTest: func(s sqlmock.Sqlmock) {
				s.ExpectExec(`UPDATE invitations
						SET status = ?
						WHERE id = ?;`).
					WithArgs(invitations[0].Status, invitations[0].MemberID).
					WillReturnError(errFoo)
			},
			want:      errFoo,
			wantError: true,
		},
	}
	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			defer mockDB.Close()

			db := sqlx.NewDb(mockDB, "sqlmock")

			r := &Repository{
				db: db,
			}

			if tt.beforeTest != nil {
				tt.beforeTest(mockSQL)
			}

			result, err := r.Update(tt.args, invitations[0], invitations[0].ID)
			if tt.wantError {
				assert.Error(t, err)
				assert.Empty(t, result)
			} else {
				assert.NoError(t, err)
			}

			if err := mockSQL.ExpectationsWereMet(); err != nil {
				assert.Error(t, err)
			}
		})
	}
}

func TestCreateAttendee(t *testing.T) {
	testCase := []testCase{
		{
			name: "Testcase #1: Positive",
			args: ctx,
			beforeTest: func(s sqlmock.Sqlmock) {
				s.ExpectExec(`INSERT INTO attendee
						(member_id, gathering_id)
						VALUES (?, ?);`).
					WithArgs(attendee.MemberID, attendee.GatheringID).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			want:      errFoo,
			wantError: false,
		},
		{
			name: "Testcase #2: Negative",
			args: ctx,
			beforeTest: func(s sqlmock.Sqlmock) {
				s.ExpectExec(`INSERT INTO attendee
						(member_id, gathering_id)
						VALUES (?, ?);`).
					WithArgs(attendee.MemberID, attendee.GatheringID).
					WillReturnError(errFoo)
			},
			want:      errFoo,
			wantError: true,
		},
	}
	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			defer mockDB.Close()

			db := sqlx.NewDb(mockDB, "sqlmock")

			r := &Repository{
				db: db,
			}

			if tt.beforeTest != nil {
				tt.beforeTest(mockSQL)
			}

			err := r.CreateAttendee(tt.args, attendee)
			if tt.wantError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			if err := mockSQL.ExpectationsWereMet(); err != nil {
				assert.Error(t, err)
			}
		})
	}
}

func TestCount(t *testing.T) {
	expectedCount := int64(10)
	testCase := []testCase{
		{
			name: "Testcase #1: Positive",
			args: ctx,
			beforeTest: func(s sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"count"}).
					AddRow(expectedCount)
				s.ExpectQuery("SELECT count(*) FROM invitations;").
					WillReturnRows(rows)
			},
			want:      nil,
			wantError: false,
		},
		{
			name: "Testcase #2: Negative",
			args: ctx,
			beforeTest: func(s sqlmock.Sqlmock) {
				s.ExpectQuery("SELECT count(*) FROM invitations;").
					WillReturnError(errFoo)
			},
			want:      errFoo,
			wantError: true,
		},
	}
	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			defer mockDB.Close()

			db := sqlx.NewDb(mockDB, "sqlmock")

			r := &Repository{
				db: db,
			}

			if tt.beforeTest != nil {
				tt.beforeTest(mockSQL)
			}

			total, err := r.Count(tt.args)
			if tt.wantError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, expectedCount, total)
			}

			if err := mockSQL.ExpectationsWereMet(); err != nil {
				assert.Error(t, err)
			}
		})
	}
}

func TestGetByMemberID(t *testing.T) {
	testCase := []testCase{
		{
			name: "Testcase #1: Positive",
			args: ctx,
			beforeTest: func(s sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{
					"iid", "i.member_id", "i.gathering_id", "i.status", "gid", "g.creator",
					"g.type", "g.name", "g.location", "g.schedule_at",
				}).
					AddRow(invitations[0].ID, invitations[0].MemberID, invitations[0].GatheringID, invitations[0].Status,
						gathering.ID, gathering.Creator, gathering.Type, gathering.Name, gathering.Location, gathering.ScheduleAt)
				s.ExpectQuery(`SELECT i.id as iid, i.member_id,
						i.gathering_id, i.status, g.id as gid, g.creator,
						g.type, g.name, g.location, g.schedule_at 
						FROM invitations i
						LEFT JOIN gatherings g ON i.gathering_id = g.id
						WHERE i.member_id = ?`).
					WithArgs(invitations[0].MemberID).
					WillReturnRows(rows)
			},
			want:      nil,
			wantError: false,
		},
		{
			name: "Testcase #2: Negative",
			args: ctx,
			beforeTest: func(s sqlmock.Sqlmock) {
				s.ExpectQuery(`SELECT i.id as iid, i.member_id,
						i.gathering_id, i.status, g.id as gid, g.creator,
						g.type, g.name, g.location, g.schedule_at 
						FROM invitations i
						LEFT JOIN gatherings g ON i.gathering_id = g.id
						WHERE i.member_id = ?`).
					WithArgs(invitations[0].MemberID).
					WillReturnError(errFoo)
			},
			want:      errFoo,
			wantError: true,
		},
		{
			name: "Testcase #3: Negative",
			args: ctx,
			beforeTest: func(s sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{
					"iid", "i.member_id", "i.gathering_id", "i.status", "gid", "g.creator",
					"g.type", "g.name", "g.location", "g.schedule_at",
				}).
					AddRow(nil, invitations[0].MemberID, invitations[0].GatheringID, invitations[0].Status,
						gathering.ID, gathering.Creator, gathering.Type, gathering.Name, gathering.Location, gathering.ScheduleAt)
				s.ExpectQuery(`SELECT i.id as iid, i.member_id,
						i.gathering_id, i.status, g.id as gid, g.creator,
						g.type, g.name, g.location, g.schedule_at 
						FROM invitations i
						LEFT JOIN gatherings g ON i.gathering_id = g.id
						WHERE i.member_id = ?`).
					WithArgs(invitations[0].MemberID).
					WillReturnRows(rows)
			},
			want:      errFoo,
			wantError: true,
		},
	}
	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			defer mockDB.Close()

			db := sqlx.NewDb(mockDB, "sqlmock")

			r := &Repository{
				db: db,
			}

			if tt.beforeTest != nil {
				tt.beforeTest(mockSQL)
			}

			invitation, err := r.GetByMemberID(tt.args, invitations[0].MemberID)
			if tt.wantError {
				assert.Error(t, err)
				assert.Empty(t, invitation)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, invitation)
			}

			if err := mockSQL.ExpectationsWereMet(); err != nil {
				assert.Error(t, err)
			}
		})
	}
}
