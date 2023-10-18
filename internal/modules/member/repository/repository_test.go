package repository

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/rzfhlv/gin-example/internal/modules/member/model"
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
	ctx     = context.Background()
	members = []model.Member{
		{
			ID: 1, FirstName: "John", LastName: "Doe", Email: "john@test.com", Password: "password", CreatedAt: time.Now(),
		},
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
				s.ExpectExec("INSERT INTO members (first_name, last_name, email, password, created_at) VALUES (?, ?, ?, ?, ?);").
					WithArgs(members[0].FirstName, members[0].LastName, members[0].Email, members[0].Password, members[0].CreatedAt).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			want:      errFoo,
			wantError: false,
		},
		{
			name: "Testcase #2: Negative",
			args: ctx,
			beforeTest: func(s sqlmock.Sqlmock) {
				s.ExpectExec("INSERT INTO members (first_name, last_name, email, password, created_at) VALUES (?, ?, ?, ?, ?);").
					WithArgs(members[0].FirstName, members[0].LastName, members[0].Email, members[0].Password, members[0].CreatedAt).
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

			result, err := r.Create(tt.args, members[0])
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
					"id", "first_name", "last_name", "email", "created_at",
				}).
					AddRow(members[0].ID, members[0].FirstName, members[0].LastName, members[0].Email, members[0].CreatedAt)
				s.ExpectQuery("SELECT id, first_name, last_name, email, created_at FROM members ORDER BY id DESC LIMIT ? OFFSET ?;").
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
				s.ExpectQuery("SELECT id, first_name, last_name, email, created_at FROM members ORDER BY id DESC LIMIT ? OFFSET ?;").
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

			members, err := r.Get(tt.args, paramTest)
			if tt.wantError {
				assert.Error(t, err)
				assert.Empty(t, members)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, members)
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
					"id", "first_name", "last_name", "email", "created_at",
				}).
					AddRow(members[0].ID, members[0].FirstName, members[0].LastName, members[0].Email, members[0].CreatedAt)
				s.ExpectQuery("SELECT id, first_name, last_name, email FROM members WHERE id = ?;").
					WithArgs(members[0].ID).
					WillReturnRows(rows)
			},
			want:      nil,
			wantError: false,
		},
		{
			name: "Testcase #2: Negative",
			args: ctx,
			beforeTest: func(s sqlmock.Sqlmock) {
				s.ExpectQuery("SELECT id, first_name, last_name, email FROM members WHERE id = ?;").
					WithArgs(members[0].ID).
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

			member, err := r.GetByID(tt.args, members[0].ID)
			if tt.wantError {
				assert.Error(t, err)
				assert.Empty(t, member)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, member)
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
				s.ExpectQuery("SELECT count(*) FROM members;").
					WillReturnRows(rows)
			},
			want:      nil,
			wantError: false,
		},
		{
			name: "Testcase #2: Negative",
			args: ctx,
			beforeTest: func(s sqlmock.Sqlmock) {
				s.ExpectQuery("SELECT count(*) FROM members;").
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
