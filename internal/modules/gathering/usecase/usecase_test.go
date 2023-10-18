package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/rzfhlv/gin-example/internal/modules/gathering/model"
	"github.com/rzfhlv/gin-example/pkg/param"
	mockRepo "github.com/rzfhlv/gin-example/shared/mocks/modules/gathering/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type testCase struct {
	name      string
	wantError error
	isErr     bool
}

var (
	errFoo           = errors.New("error")
	gatheringPayload = model.Gathering{
		ID:           1,
		Creator:      "John Doe",
		Type:         "family",
		Name:         "Family Gathering",
		Location:     "Puncak",
		ScheduleAt:   "2023-11-10 12:00:00",
		MemberID:     1,
		ScheduleAtDB: time.Now(),
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

	u := New(&mockRepo)
	assert.NotNil(t, u)
}

func TestCreate(t *testing.T) {
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
			mockRepo.On("Create", mock.Anything, mock.Anything).Return(&CustomResult{lastInsertID: 1, rowsAffected: 1}, tt.wantError)

			u := &Usecase{
				repo: &mockRepo,
			}

			_, err := u.Create(context.Background(), gatheringPayload)
			assert.EqualValues(t, err, tt.wantError)
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
			mockRepo.On("Get", mock.Anything, mock.Anything).Return([]model.Gathering{}, tt.wantError)
			mockRepo.On("Count", mock.Anything).Return(expectedCount, tt.wantError)

			u := &Usecase{
				repo: &mockRepo,
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
			mockRepo.On("GetByID", mock.Anything, mock.Anything).Return(model.Gathering{}, tt.wantError)

			u := &Usecase{
				repo: &mockRepo,
			}

			_, err := u.GetByID(context.Background(), gatheringPayload.ID)
			assert.EqualValues(t, err, tt.wantError)
		})
	}
}

func TestGetDetailByID(t *testing.T) {
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
			mockRepo.On("GetByID", mock.Anything, mock.Anything).Return(model.Gathering{}, tt.wantError)
			mockRepo.On("GetDetailByID", mock.Anything, mock.Anything).Return(model.GatheringDetail{}, tt.wantError)

			u := &Usecase{
				repo: &mockRepo,
			}

			_, err := u.GetDetailByID(context.Background(), gatheringPayload.ID)
			assert.EqualValues(t, err, tt.wantError)
		})
	}
}
