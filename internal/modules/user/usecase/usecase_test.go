package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/rzfhlv/gin-example/internal/modules/user/model"
	"github.com/rzfhlv/gin-example/pkg/param"
	mockRepo "github.com/rzfhlv/gin-example/shared/mocks/modules/user/repository"
	mockHasher "github.com/rzfhlv/gin-example/shared/mocks/pkg/hasher"
	mockJwt "github.com/rzfhlv/gin-example/shared/mocks/pkg/jwt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

type testCase struct {
	name                                                                string
	wantError, wantIDError, wantJwtError, wantRedisError, wantHashError error
	result                                                              CustomResult
	payload                                                             model.Register
	isErr                                                               bool
}

var (
	token    string
	errFoo   = errors.New("error")
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
	mockHasher := mockHasher.HashPassword{}
	mockJwt := mockJwt.JWTInterface{}

	u := New(&mockRepo, &mockHasher, &mockJwt)
	assert.NotNil(t, u)
}

func TestRegister(t *testing.T) {
	testCase := []testCase{
		{
			name: "Testcase #1: Positive", wantError: nil, wantIDError: nil, wantJwtError: nil, wantRedisError: nil, wantHashError: nil, result: CustomResult{lastInsertID: 1, rowsAffected: 1, err: nil}, payload: register, isErr: false,
		},
		{
			name: "Testcase #2: Negative", wantError: errFoo, wantIDError: nil, wantJwtError: errFoo, wantRedisError: nil, wantHashError: nil, result: CustomResult{lastInsertID: 1, rowsAffected: 1, err: nil}, payload: register, isErr: true,
		},
		{
			name: "Testcase #3: Negative", wantError: nil, wantIDError: errFoo, wantJwtError: nil, wantRedisError: nil, wantHashError: nil, result: CustomResult{lastInsertID: 1, rowsAffected: 1, err: errFoo}, payload: register, isErr: true,
		},
		{
			name: "Testcase #4: Negative", wantError: nil, wantIDError: nil, wantJwtError: errFoo, wantRedisError: nil, wantHashError: nil, result: CustomResult{lastInsertID: 1, rowsAffected: 1, err: nil}, payload: register, isErr: true,
		},
		{
			name: "Testcase #5: Negative", wantError: nil, wantIDError: nil, wantJwtError: nil, wantRedisError: errFoo, wantHashError: nil, result: CustomResult{lastInsertID: 1, rowsAffected: 1, err: nil}, payload: register, isErr: true,
		},
		{
			name: "Testcase #6: Negative", wantError: nil, wantIDError: nil, wantJwtError: nil, wantRedisError: nil, wantHashError: errFoo, result: CustomResult{lastInsertID: 1, rowsAffected: 1, err: nil}, payload: register, isErr: true,
		},
	}
	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := mockRepo.IRepository{}
			mockHasher := mockHasher.HashPassword{}
			mockJwt := mockJwt.JWTInterface{}

			mockRepo.On("Register", mock.Anything, mock.Anything).Return(&tt.result, tt.wantError)
			mockRepo.On("Set", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.wantRedisError)
			mockHasher.On("HashedPassword", mock.Anything).Return("", tt.wantHashError)
			mockJwt.On("Generate", mock.Anything, mock.Anything, mock.Anything).Return(token, tt.wantJwtError)

			u := &Usecase{
				repo:    &mockRepo,
				hasher:  &mockHasher,
				jwtImpl: &mockJwt,
			}

			_, err := u.Register(context.Background(), tt.payload)
			if tt.isErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestLogin(t *testing.T) {
	testCase := []testCase{
		{
			name: "Testcase #1: Positive", wantError: nil, wantJwtError: nil, wantRedisError: nil, wantHashError: nil, isErr: false,
		},
		{
			name: "Testcase #2: Negative", wantError: errFoo, wantJwtError: nil, wantRedisError: nil, wantHashError: nil, isErr: true,
		},
		{
			name: "Testcase #3: Negative", wantError: nil, wantJwtError: errFoo, wantRedisError: nil, wantHashError: nil, isErr: true,
		},
		{
			name: "Testcase #4: Negative", wantError: nil, wantJwtError: nil, wantRedisError: errFoo, wantHashError: nil, isErr: true,
		},
		{
			name: "Testcase #5: Negative", wantError: nil, wantJwtError: nil, wantRedisError: nil, wantHashError: bcrypt.ErrMismatchedHashAndPassword, isErr: true,
		},
	}
	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := mockRepo.IRepository{}
			mockHasher := mockHasher.HashPassword{}
			mockJwt := mockJwt.JWTInterface{}

			mockRepo.On("Login", mock.Anything, mock.Anything).Return(model.Register{}, tt.wantError)
			mockRepo.On("Set", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.wantRedisError)
			mockHasher.On("VerifyPassword", mock.Anything, mock.Anything).Return(tt.wantHashError)
			mockJwt.On("Generate", mock.Anything, mock.Anything, mock.Anything).Return(token, tt.wantJwtError)

			u := &Usecase{
				repo:    &mockRepo,
				hasher:  &mockHasher,
				jwtImpl: &mockJwt,
			}

			_, err := u.Login(context.Background(), login)
			if tt.isErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestLogout(t *testing.T) {
	token := "thisistoken"
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
			mockRepo.On("Del", mock.Anything, mock.Anything).Return(tt.wantError)

			mockJwt := mockJwt.JWTInterface{}
			mockHasher := mockHasher.HashPassword{}

			u := &Usecase{
				repo:    &mockRepo,
				jwtImpl: &mockJwt,
				hasher:  &mockHasher,
			}

			err := u.Logout(context.Background(), token)
			if tt.isErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetAll(t *testing.T) {
	expectedCount := int64(10)
	testCase := []testCase{
		{
			name: "Testcase #1: Positive", wantError: nil,
		},
		{
			name: "Testcase #2: Negative", wantError: errFoo,
		},
	}
	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := mockRepo.IRepository{}
			mockRepo.On("GetAll", mock.Anything, mock.Anything).Return([]model.User{}, tt.wantError)
			mockRepo.On("Count", mock.Anything).Return(expectedCount, tt.wantError)

			u := &Usecase{
				repo: &mockRepo,
			}

			_, _, err := u.GetAll(context.Background(), param.Param{})
			assert.EqualValues(t, err, tt.wantError)
		})
	}
}

func TestGetByID(t *testing.T) {
	testCase := []testCase{
		{
			name: "Testcase #1: Positive", wantError: nil,
		},
		{
			name: "Testcase #2: Negative", wantError: errFoo,
		},
	}
	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := mockRepo.IRepository{}
			mockRepo.On("GetByID", mock.Anything, mock.Anything).Return(model.User{}, tt.wantError)

			u := &Usecase{
				repo: &mockRepo,
			}

			_, err := u.GetByID(context.Background(), register.ID)
			assert.EqualValues(t, err, tt.wantError)
		})
	}
}
