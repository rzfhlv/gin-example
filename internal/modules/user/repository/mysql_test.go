package repository

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/rzfhlv/gin-example/internal/modules/user/model"
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
	users = []model.User{
		{
			ID: 1, Username: "johndoe", Email: "john@test.com", CreatedAt: time.Now(),
		},
	}
	errFoo    = errors.New("foo")
	paramTest = param.Param{
		Limit: 10,
		Page:  1,
	}
	register = model.Register{
		ID:        1,
		Username:  "johndoe",
		Email:     "johndoe@test.com",
		Password:  "password",
		CreatedAt: time.Now(),
	}
	login = model.Login{
		Username: "johndoe",
		Password: "password",
	}
)

func TestRegister(t *testing.T) {
	testCase := []testCase{
		{
			name: "Testcase #1: Positive",
			args: ctx,
			beforeTest: func(s sqlmock.Sqlmock) {
				s.ExpectExec("INSERT INTO users (username, email, password, created_at) VALUES (?, ?, ?, ?);").
					WithArgs(register.Username, register.Email, register.Password, register.CreatedAt).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			want:      errFoo,
			wantError: false,
		},
		{
			name: "Testcase #2: Negative",
			args: ctx,
			beforeTest: func(s sqlmock.Sqlmock) {
				s.ExpectExec("INSERT INTO users (username, email, password, created_at) VALUES (?, ?, ?, ?);").
					WithArgs(register.Username, register.Email, register.Password, register.CreatedAt).
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

			result, err := r.Register(tt.args, register)
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

func TestLogin(t *testing.T) {
	testCase := []testCase{
		{
			name: "Testcase #1: Positive",
			args: ctx,
			beforeTest: func(s sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{
					"id", "username", "email", "password", "created_at",
				}).
					AddRow(register.ID, register.Username, register.Email, register.Password, register.CreatedAt)
				s.ExpectQuery("SELECT id, username, email, password, created_at FROM users WHERE username = ?;").
					WithArgs(register.Username).
					WillReturnRows(rows)
			},
			want:      nil,
			wantError: false,
		},
		{
			name: "Testcase #2: Negative",
			args: ctx,
			beforeTest: func(s sqlmock.Sqlmock) {
				s.ExpectQuery("SELECT id, username, email, password, created_at FROM users WHERE username = ?;").
					WithArgs(register.Username).
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

			register, err := r.Login(tt.args, login)
			if tt.wantError {
				assert.Error(t, err)
				assert.Empty(t, register)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, register)
			}

			if err := mockSQL.ExpectationsWereMet(); err != nil {
				assert.Error(t, err)
			}
		})
	}
}

func TestGetAll(t *testing.T) {
	testCase := []testCase{
		{
			name: "Testcase #1: Positive",
			args: ctx,
			beforeTest: func(s sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{
					"id", "username", "email", "created_at",
				}).
					AddRow(users[0].ID, users[0].Username, users[0].Email, users[0].CreatedAt)
				s.ExpectQuery("SELECT id, username, email, created_at FROM users ORDER BY id DESC LIMIT ? OFFSET ?;").
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
				s.ExpectQuery("SELECT id, username, email, created_at FROM users ORDER BY id DESC LIMIT ? OFFSET ?;").
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

			users, err := r.GetAll(tt.args, paramTest)
			if tt.wantError {
				assert.Error(t, err)
				assert.Empty(t, users)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, users)
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
					"id", "username", "email", "created_at",
				}).
					AddRow(users[0].ID, users[0].Username, users[0].Email, users[0].CreatedAt)
				s.ExpectQuery("SELECT id, username, email, created_at FROM users WHERE id = ?;").
					WithArgs(users[0].ID).
					WillReturnRows(rows)
			},
			want:      nil,
			wantError: false,
		},
		{
			name: "Testcase #2: Negative",
			args: ctx,
			beforeTest: func(s sqlmock.Sqlmock) {
				s.ExpectQuery("SELECT id, username, email, created_at FROM users WHERE id = ?;").
					WithArgs(users[0].ID).
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

			user, err := r.GetByID(tt.args, users[0].ID)
			if tt.wantError {
				assert.Error(t, err)
				assert.Empty(t, user)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, user)
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
				s.ExpectQuery("SELECT count(*) FROM users;").
					WillReturnRows(rows)
			},
			want:      nil,
			wantError: false,
		},
		{
			name: "Testcase #2: Negative",
			args: ctx,
			beforeTest: func(s sqlmock.Sqlmock) {
				s.ExpectQuery("SELECT count(*) FROM users;").
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
