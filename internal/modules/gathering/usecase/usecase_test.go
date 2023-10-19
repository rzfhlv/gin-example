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
	name                   string
	wantError, wantIDError error
	isErr                  bool
	result                 CustomResult
	payload                model.Gathering
}

var (
	scheduleAt, _    = time.Parse("2006-01-02 03:04:05", "2023-11-10 12:00:00")
	errFoo           = errors.New("error")
	errTime          = &time.ParseError{Layout: "2006-01-02 03:04:05", Value: "2023-11-10", LayoutElem: "03", ValueElem: "", Message: ""}
	gatheringPayload = model.Gathering{
		ID:           1,
		Creator:      "John Doe",
		Type:         "family",
		Name:         "Family Gathering",
		Location:     "Puncak",
		ScheduleAt:   "2023-11-10 12:00:00",
		MemberID:     1,
		ScheduleAtDB: scheduleAt,
	}
	gatheringPayloadFail = model.Gathering{
		ID:           1,
		Creator:      "John Doe",
		Type:         "family",
		Name:         "Family Gathering",
		Location:     "Puncak",
		ScheduleAt:   "2023-11-10",
		MemberID:     1,
		ScheduleAtDB: scheduleAt,
	}
)

type CustomResult struct {
	lastInsertID int64
	rowsAffected int64
	err          error
}

func (r *CustomResult) LastInsertId() (int64, error) {
	return r.lastInsertID, r.err
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
			name: "Testcase #1: Positive", wantError: nil, wantIDError: nil, isErr: false, result: CustomResult{lastInsertID: 1, rowsAffected: 1, err: nil}, payload: gatheringPayload,
		},
		{
			name: "Testcase #2: Negative", wantError: errFoo, wantIDError: nil, isErr: false, result: CustomResult{lastInsertID: 1, rowsAffected: 1, err: nil}, payload: gatheringPayload,
		},
		{
			name: "Testcase #3: Negative", wantError: errTime, wantIDError: nil, isErr: false, result: CustomResult{lastInsertID: 1, rowsAffected: 1, err: nil}, payload: gatheringPayloadFail,
		},
		{
			name: "Testcase #4: Negative", wantError: nil, wantIDError: errFoo, isErr: true, result: CustomResult{lastInsertID: 1, rowsAffected: 1, err: errFoo}, payload: gatheringPayload,
		},
	}
	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := mockRepo.IRepository{}
			mockRepo.On("Create", mock.Anything, mock.Anything).Return(&tt.result, tt.wantError)

			u := &Usecase{
				repo: &mockRepo,
			}

			_, err := u.Create(context.Background(), tt.payload)
			if tt.isErr {
				assert.EqualValues(t, err, tt.wantIDError)
			} else {
				assert.EqualValues(t, err, tt.wantError)
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
