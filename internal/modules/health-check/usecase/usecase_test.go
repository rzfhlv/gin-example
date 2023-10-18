package usecase

import (
	"context"
	"errors"
	"testing"

	mockRepo "github.com/rzfhlv/gin-example/shared/mocks/modules/health-check/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type testCase struct {
	name      string
	wantError error
	isErr     bool
}

var (
	errFoo = errors.New("error")
)

func TestNew(t *testing.T) {
	mockRepo := mockRepo.IRepository{}
	u := New(&mockRepo)
	assert.NotNil(t, u)
}

func TestUsecase(t *testing.T) {
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
			mockRepo.On("Ping", mock.Anything).Return(tt.wantError)

			u := &Usecase{
				repo: &mockRepo,
			}

			err := u.Ping(context.Background())
			assert.EqualValues(t, err, tt.wantError)
		})
	}
}
