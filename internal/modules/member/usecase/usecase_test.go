package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/rzfhlv/gin-example/internal/modules/member/model"
	"github.com/rzfhlv/gin-example/pkg/param"
	mockRepo "github.com/rzfhlv/gin-example/shared/mocks/modules/member/repository"
	mockHasher "github.com/rzfhlv/gin-example/shared/mocks/pkg/hasher"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type testCase struct {
	name                       string
	wantError, wantHasherError error
	isErr                      bool
}

var (
	errFoo        = errors.New("error")
	memberPayload = model.Member{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@test.com",
		Password:  "password",
		CreatedAt: time.Now(),
	}
)

type CustomResult struct {
	lastInsertID int64
	rowsAffected int64
}

func (r *CustomResult) LastInsertId() (int64, error) {
	return r.lastInsertID, nil
}

func (r *CustomResult) RowsAffected() (int64, error) {
	return r.rowsAffected, nil
}

func TestNew(t *testing.T) {
	mockRepo := mockRepo.IRepository{}
	mockHaser := mockHasher.HashPassword{}

	u := New(&mockRepo, &mockHaser)
	assert.NotNil(t, u)
}

func TestCreate(t *testing.T) {
	testCase := []testCase{
		{
			name: "Testcase #1: Positive", wantError: nil, wantHasherError: nil, isErr: false,
		},
		{
			name: "Testcase #2: Negative", wantError: errFoo, wantHasherError: nil, isErr: true,
		},
		{
			name: "Testcase #2: Negative", wantError: nil, wantHasherError: errFoo, isErr: true,
		},
	}
	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := mockRepo.IRepository{}
			mockHasher := mockHasher.HashPassword{}
			mockHasher.On("HashedPassword", mock.Anything).Return("", tt.wantHasherError)
			mockRepo.On("Create", mock.Anything, mock.Anything).Return(&CustomResult{lastInsertID: 1, rowsAffected: 1}, tt.wantError)

			u := &Usecase{
				repo:   &mockRepo,
				hasher: &mockHasher,
			}

			_, err := u.Create(context.Background(), memberPayload)
			if tt.wantError != nil {
				assert.EqualValues(t, err, tt.wantError)
			} else {
				assert.EqualValues(t, err, tt.wantHasherError)
			}
		})
	}
}

func TestGet(t *testing.T) {
	expectedCount := int64(10)
	testCase := []testCase{
		{
			name: "Testcase #1: Positive", wantError: nil, isErr: false,
		},
		{
			name: "Testcase #2: Negative", wantError: errFoo, isErr: true,
		},
	}
	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := mockRepo.IRepository{}
			mockHasher := mockHasher.HashPassword{}
			mockRepo.On("Get", mock.Anything, mock.Anything).Return([]model.Member{}, tt.wantError)
			mockRepo.On("Count", mock.Anything).Return(expectedCount, tt.wantError)

			u := &Usecase{
				repo:   &mockRepo,
				hasher: &mockHasher,
			}

			_, _, err := u.Get(context.Background(), param.Param{})
			assert.EqualValues(t, err, tt.wantError)
		})
	}
}

func TestGetByID(t *testing.T) {
	testCase := []testCase{
		{
			name: "Testcase #1: Positive", wantError: nil, isErr: false,
		},
		{
			name: "Testcase #2: Negative", wantError: errFoo, isErr: true,
		},
	}
	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := mockRepo.IRepository{}
			mockHasher := mockHasher.HashPassword{}
			mockRepo.On("GetByID", mock.Anything, mock.Anything).Return(model.Member{}, tt.wantError)

			u := &Usecase{
				repo:   &mockRepo,
				hasher: &mockHasher,
			}

			_, err := u.GetByID(context.Background(), memberPayload.ID)
			assert.EqualValues(t, err, tt.wantError)
		})
	}
}
