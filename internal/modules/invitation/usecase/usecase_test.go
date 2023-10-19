package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/rzfhlv/gin-example/internal/modules/invitation/model"
	"github.com/rzfhlv/gin-example/pkg/param"
	mockRepo "github.com/rzfhlv/gin-example/shared/mocks/modules/invitation/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type testCase struct {
	name                                      string
	wantError, wantAttendeeError, wantIDError error
	isErr                                     bool
	result                                    CustomResult
}

var (
	errFoo            = errors.New("error")
	invitationPayload = model.Invitation{
		ID:          1,
		MemberID:    1,
		GatheringID: 1,
		Status:      "accept",
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
			name: "Testcase #1: Positive", wantError: nil, wantAttendeeError: nil, wantIDError: nil, isErr: false, result: CustomResult{lastInsertID: 1, rowsAffected: 1, err: nil},
		},
		{
			name: "Testcase #2: Negative", wantError: errFoo, wantAttendeeError: nil, wantIDError: nil, isErr: true, result: CustomResult{lastInsertID: 1, rowsAffected: 1, err: nil},
		},
		{
			name: "Testcase #3: Negative", wantError: nil, wantAttendeeError: errFoo, wantIDError: nil, isErr: true, result: CustomResult{lastInsertID: 1, rowsAffected: 1, err: nil},
		},
		{
			name: "Testcase #4: Negative", wantError: nil, wantAttendeeError: nil, wantIDError: errFoo, isErr: true, result: CustomResult{lastInsertID: 1, rowsAffected: 1, err: errFoo},
		},
	}
	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := mockRepo.IRepository{}
			mockRepo.On("Create", mock.Anything, mock.Anything).Return(&tt.result, tt.wantError)
			mockRepo.On("CreateAttendee", mock.Anything, mock.Anything).Return(tt.wantAttendeeError)

			u := &Usecase{
				repo: &mockRepo,
			}

			_, err := u.Create(context.Background(), invitationPayload)
			if tt.wantAttendeeError != nil {
				assert.EqualValues(t, err, tt.wantAttendeeError)
			} else if tt.wantError != nil {
				assert.EqualValues(t, err, tt.wantError)
			} else {
				assert.EqualValues(t, err, tt.wantIDError)
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
			mockRepo.On("Get", mock.Anything, mock.Anything).Return([]model.Invitation{}, tt.wantError)
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
			mockRepo.On("GetByID", mock.Anything, mock.Anything).Return(model.Invitation{}, tt.wantError)

			u := &Usecase{
				repo: &mockRepo,
			}

			_, err := u.GetByID(context.Background(), invitationPayload.ID)
			assert.EqualValues(t, err, tt.wantError)
		})
	}
}

func TestUpdate(t *testing.T) {
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
			mockRepo.On("Update", mock.Anything, mock.Anything, mock.Anything).Return(&CustomResult{lastInsertID: 1, rowsAffected: 1}, tt.wantError)

			u := &Usecase{
				repo: &mockRepo,
			}

			_, err := u.Update(context.Background(), invitationPayload, invitationPayload.ID)
			assert.EqualValues(t, err, tt.wantError)
		})
	}
}

func TestGetByMemberID(t *testing.T) {
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
			mockRepo.On("GetByMemberID", mock.Anything, mock.Anything).Return([]model.InvitationDetail{}, tt.wantError)

			u := &Usecase{
				repo: &mockRepo,
			}

			_, err := u.GetByMemberID(context.Background(), invitationPayload.MemberID)
			assert.EqualValues(t, err, tt.wantError)
		})
	}
}
