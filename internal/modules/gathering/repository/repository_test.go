package repository

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/rzfhlv/gin-example/internal/modules/gathering/model"
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
	ctx        = context.Background()
	gatherings = []model.Gathering{
		{
			ID: 1, Creator: "John", Type: "family", Name: "Family Gathering",
			Location: "Jakarta", ScheduleAtDB: time.Now(), MemberID: 1,
		},
	}
	errFoo    = errors.New("foo")
	paramTest = param.Param{
		Limit: 10,
		Page:  1,
	}
	detailGatherings = []model.Attendee{
		{
			ID: 1, FirstName: "John", LastName: "Doe", Email: "john@test.com", Status: "accept",
		},
	}
)

func TestCreate(t *testing.T) {
	testCase := []testCase{
		{
			name: "Testcase #1: Positive",
			args: ctx,
			beforeTest: func(s sqlmock.Sqlmock) {
				s.ExpectExec("INSERT INTO gatherings (creator, member_id, type, name, location, schedule_at) VALUES (?, ?, ?, ?, ?, ?);").
					WithArgs(gatherings[0].Creator, gatherings[0].MemberID, gatherings[0].Type, gatherings[0].Name, gatherings[0].Location, gatherings[0].ScheduleAtDB).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			want:      errFoo,
			wantError: false,
		},
		{
			name: "Testcase #2: Negative",
			args: ctx,
			beforeTest: func(s sqlmock.Sqlmock) {
				s.ExpectExec("INSERT INTO gatherings (creator, member_id, type, name, location, schedule_at) VALUES (?, ?, ?, ?, ?, ?);").
					WithArgs(gatherings[0].Creator, gatherings[0].MemberID, gatherings[0].Type, gatherings[0].Name, gatherings[0].Location, gatherings[0].ScheduleAtDB).
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

			result, err := r.Create(tt.args, gatherings[0])
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
					"id", "creator", "member_id", "type", "name", "location", "schedule_at",
				}).
					AddRow(gatherings[0].ID, gatherings[0].Creator, gatherings[0].MemberID, gatherings[0].Type,
						gatherings[0].Name, gatherings[0].Location, gatherings[0].ScheduleAtDB)
				s.ExpectQuery("SELECT id, creator, member_id, type, name, location, schedule_at FROM gatherings ORDER BY id DESC LIMIT ? OFFSET ?;").
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
				s.ExpectQuery("SELECT id, creator, member_id, type, name, location, schedule_at FROM gatherings ORDER BY id DESC LIMIT ? OFFSET ?;").
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

			gatherings, err := r.Get(tt.args, paramTest)
			if tt.wantError {
				assert.Error(t, err)
				assert.Empty(t, gatherings)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, gatherings)
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
					"id", "creator", "member_id", "type", "name", "location", "schedule_at",
				}).
					AddRow(gatherings[0].ID, gatherings[0].Creator, gatherings[0].MemberID, gatherings[0].Type,
						gatherings[0].Name, gatherings[0].Location, gatherings[0].ScheduleAtDB)
				s.ExpectQuery("SELECT id, creator, member_id, type, name, location, schedule_at FROM gatherings WHERE id = ?;").
					WithArgs(gatherings[0].ID).
					WillReturnRows(rows)
			},
			want:      nil,
			wantError: false,
		},
		{
			name: "Testcase #2: Negative",
			args: ctx,
			beforeTest: func(s sqlmock.Sqlmock) {
				s.ExpectQuery("SELECT id, creator, member_id, type, name, location, schedule_at FROM gatherings WHERE id = ?;").
					WithArgs(gatherings[0].ID).
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

			gathering, err := r.GetByID(tt.args, gatherings[0].ID)
			if tt.wantError {
				assert.Error(t, err)
				assert.Empty(t, gathering)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, gathering)
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
				s.ExpectQuery("SELECT count(*) FROM gatherings;").
					WillReturnRows(rows)
			},
			want:      nil,
			wantError: false,
		},
		{
			name: "Testcase #2: Negative",
			args: ctx,
			beforeTest: func(s sqlmock.Sqlmock) {
				s.ExpectQuery("SELECT count(*) FROM gatherings;").
					WithArgs(gatherings[0].ID).
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

func TestGetDetailByID(t *testing.T) {
	testCase := []testCase{
		{
			name: "Testcase #1: Positive",
			args: ctx,
			beforeTest: func(s sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{
					"m.id", "m.first_name", "m.last_name", "m.email", "i.status",
				}).
					AddRow(detailGatherings[0].ID, detailGatherings[0].FirstName,
						detailGatherings[0].LastName, detailGatherings[0].Email, detailGatherings[0].Status)
				s.ExpectQuery(`SELECT m.id, m.first_name, m.last_name, m.email, i.status
				FROM members m
				LEFT JOIN attendee a ON m.id = a.member_id
				LEFT JOIN invitations i ON a.gathering_id = i.gathering_id
				AND a.member_id = i.member_id
				WHERE a.gathering_id = ?;`).
					WithArgs(gatherings[0].ID).
					WillReturnRows(rows)
			},
			want:      nil,
			wantError: false,
		},
		{
			name: "Testcase #2: Negative",
			args: ctx,
			beforeTest: func(s sqlmock.Sqlmock) {
				s.ExpectQuery(`SELECT m.id, m.first_name, m.last_name, m.email, i.status
				FROM members m
				LEFT JOIN attendee a ON m.id = a.member_id
				LEFT JOIN invitations i ON a.gathering_id = i.gathering_id
				AND a.member_id = i.member_id
				WHERE a.gathering_id = ?;`).
					WithArgs(gatherings[0].ID).
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
					"m.id", "m.first_name", "m.last_name", "m.email", "i.status",
				}).
					AddRow(nil, detailGatherings[0].FirstName,
						detailGatherings[0].LastName, detailGatherings[0].Email, detailGatherings[0].Status)
				s.ExpectQuery(`SELECT m.id, m.first_name, m.last_name, m.email, i.status
				FROM members m
				LEFT JOIN attendee a ON m.id = a.member_id
				LEFT JOIN invitations i ON a.gathering_id = i.gathering_id
				AND a.member_id = i.member_id
				WHERE a.gathering_id = ?;`).
					WithArgs(gatherings[0].ID).
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

			gatherings, err := r.GetDetailByID(tt.args, gatherings[0].ID)
			if tt.wantError {
				assert.Error(t, err)
				assert.Empty(t, gatherings)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, gatherings)
			}

			if err := mockSQL.ExpectationsWereMet(); err != nil {
				assert.Error(t, err)
			}
		})
	}
}
