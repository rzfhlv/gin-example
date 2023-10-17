package repository

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestRepository(t *testing.T) {
	tests := []struct {
		name       string
		args       context.Context
		beferoTest func(sqlmock.Sqlmock)
		want       error
		wantError  bool
	}{
		{
			name: "Ping MySQL",
			args: context.Background(),
			beferoTest: func(s sqlmock.Sqlmock) {
				s.ExpectExec("SELECT 1").WillReturnError(nil)
			},
			want:      nil,
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			db := sqlx.NewDb(mockDB, "sqlmock")
			r := Repository{
				db: db,
			}

			if tt.beferoTest != nil {
				tt.beferoTest(mockSQL)
			}

			err := r.Ping(tt.args)
			if tt.wantError {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.want.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
