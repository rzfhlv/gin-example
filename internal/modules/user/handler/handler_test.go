package handler

import (
	"database/sql"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rzfhlv/gin-example/internal/modules/user/model"
	mockUsecase "github.com/rzfhlv/gin-example/shared/mocks/modules/user/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type testCase struct {
	name, body, param, queryParam, setContext string
	wantError                                 error
	code                                      int
}

var (
	errFoo         = errors.New("error")
	payloadSuccess = `{"id":1,"username":"johndoe","email":"john@test.com","password":"password"}`
	payloadFail    = `{"id":1,"username":"","email":"john@test.com","password":"password"}`
	loginSuccess   = `{"username":"johndoe","password":"password"}`
	loginFail      = `{"username":"","password":"password"}`
)

func TestNew(t *testing.T) {
	mockUsecase := mockUsecase.IUsecase{}

	h := New(&mockUsecase)
	assert.NotNil(t, h)
}

func TestRegister(t *testing.T) {
	gin.SetMode(gin.TestMode)

	testCase := []testCase{
		{
			name: "Testcase #1: Positive", body: payloadSuccess, wantError: nil, code: http.StatusOK,
		},
		{
			name: "Testcase #2: Negative", body: payloadSuccess, wantError: errFoo, code: http.StatusInternalServerError,
		},
		{
			name: "Testcase #3: Negative", body: payloadFail, wantError: errFoo, code: http.StatusUnprocessableEntity,
		},
	}
	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			mockUsecase := mockUsecase.IUsecase{}
			mockUsecase.On("Register", mock.Anything, mock.Anything).Return(model.JWT{}, tt.wantError)

			h := &Handler{
				usecase: &mockUsecase,
			}

			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = httptest.NewRequest(http.MethodPost, "/v1/users/register", strings.NewReader(tt.body))
			ctx.Request.Header.Set("Content-Type", "application/json")

			h.Register(ctx)
			assert.EqualValues(t, tt.code, w.Code)
		})
	}
}

func TestLogin(t *testing.T) {
	gin.SetMode(gin.TestMode)

	testCase := []testCase{
		{
			name: "Testcase #1: Positive", body: loginSuccess, wantError: nil, code: http.StatusOK,
		},
		{
			name: "Testcase #2: Negative", body: loginSuccess, wantError: errFoo, code: http.StatusInternalServerError,
		},
		{
			name: "Testcase #3: Negative", body: loginFail, wantError: errFoo, code: http.StatusUnprocessableEntity,
		},
	}
	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			mockUsecase := mockUsecase.IUsecase{}
			mockUsecase.On("Login", mock.Anything, mock.Anything).Return(model.JWT{}, tt.wantError)

			h := &Handler{
				usecase: &mockUsecase,
			}

			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = httptest.NewRequest(http.MethodPost, "/v1/users/login", strings.NewReader(tt.body))
			ctx.Request.Header.Set("Content-Type", "application/json")

			h.Login(ctx)
			assert.EqualValues(t, tt.code, w.Code)
		})
	}
}

func TestLogout(t *testing.T) {
	gin.SetMode(gin.TestMode)

	testCase := []testCase{
		{
			name: "Testcase #1: Positive", body: "Bearer thisistoken", wantError: nil, setContext: "username", code: http.StatusOK,
		},
		{
			name: "Testcase #2: Negative", body: "Bearer", wantError: nil, setContext: "username", code: http.StatusUnauthorized,
		},
		{
			name: "Testcase #3: Negative", body: "Bearer thisistoken", wantError: nil, setContext: "any", code: http.StatusUnauthorized,
		},
		{
			name: "Testcase #4: Negative", body: "Bearer thisistoken", wantError: errFoo, setContext: "username", code: http.StatusInternalServerError,
		},
	}
	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			mockUsecase := mockUsecase.IUsecase{}
			mockUsecase.On("Logout", mock.Anything, mock.Anything).Return(tt.wantError)

			h := &Handler{
				usecase: &mockUsecase,
			}

			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = httptest.NewRequest(http.MethodPost, "/v1/users/logout", nil)
			ctx.Request.Header.Set("Content-Type", "application/json")
			ctx.Request.Header.Set("Authorization", tt.body)
			ctx.Set(tt.setContext, "johndoe")

			h.Logout(ctx)
			assert.EqualValues(t, tt.code, w.Code)
		})
	}
}

func TestGetAll(t *testing.T) {
	gin.SetMode(gin.TestMode)

	expectedCount := int64(10)
	testCase := []testCase{
		{
			name: "Testcase #1: Positive", queryParam: "?page=1", wantError: nil, code: http.StatusOK,
		},
		{
			name: "Testcase #2: Negative", queryParam: "?page=1", wantError: errFoo, code: http.StatusInternalServerError,
		},
		{
			name: "Testcase #3: Negative", queryParam: "?page=one", wantError: errFoo, code: http.StatusUnprocessableEntity,
		},
	}
	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			mockUsecase := mockUsecase.IUsecase{}
			mockUsecase.On("GetAll", mock.Anything, mock.Anything).Return([]model.User{}, expectedCount, tt.wantError)

			h := &Handler{
				usecase: &mockUsecase,
			}

			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = httptest.NewRequest(http.MethodGet, "/v1/users"+tt.queryParam, nil)
			ctx.Request.Header.Set("Content-Type", "application/json")

			h.GetAll(ctx)
			assert.Equal(t, tt.code, w.Code)
		})
	}
}

func TestGetByID(t *testing.T) {
	gin.SetMode(gin.TestMode)

	testCase := []testCase{
		{
			name: "Testcase #1: Positive", param: "1", wantError: nil, code: http.StatusOK,
		},
		{
			name: "Testcase #2: Negative", param: "1", wantError: errFoo, code: http.StatusInternalServerError,
		},
		{
			name: "Testcase #3: Negative", param: "one", wantError: errFoo, code: http.StatusUnprocessableEntity,
		},
		{
			name: "Testcase #3: Negative", param: "0", wantError: sql.ErrNoRows, code: http.StatusNotFound,
		},
	}
	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			mockUsecase := mockUsecase.IUsecase{}
			mockUsecase.On("GetByID", mock.Anything, mock.Anything).Return(model.User{}, tt.wantError)

			h := &Handler{
				usecase: &mockUsecase,
			}

			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = httptest.NewRequest(http.MethodGet, "/v1/users/"+tt.param, nil)
			ctx.Request.Header.Set("Content-Type", "application/json")
			ctx.Params = gin.Params{{Key: "id", Value: tt.param}}

			h.GetByID(ctx)
			assert.Equal(t, tt.code, w.Code)
		})
	}
}
