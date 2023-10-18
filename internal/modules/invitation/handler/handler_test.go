package handler

import (
	"database/sql"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rzfhlv/gin-example/internal/modules/invitation/model"
	mockUsecase "github.com/rzfhlv/gin-example/shared/mocks/modules/invitation/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type testCase struct {
	name, body, param, queryParam string
	wantError                     error
	code                          int
}

var (
	errFoo         = errors.New("error")
	payloadSuccess = `{"member_id":1,"gathering_id":1,"status":"accept"}`
	payloadFail    = `{"member_id":,"gathering_id":1,"status":"accept"}`
)

func TestNew(t *testing.T) {
	mockUsecase := mockUsecase.IUsecase{}

	h := New(&mockUsecase)
	assert.NotNil(t, h)
}

func TestCreate(t *testing.T) {
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
			mockUsecase.On("Create", mock.Anything, mock.Anything).Return(model.Invitation{}, tt.wantError)

			h := &Handler{
				usecase: &mockUsecase,
			}

			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = httptest.NewRequest(http.MethodPost, "/v1/invitations", strings.NewReader(tt.body))
			ctx.Request.Header.Set("Content-Type", "application/json")

			h.Create(ctx)
			assert.EqualValues(t, tt.code, w.Code)
		})
	}
}

func TestGet(t *testing.T) {
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
			mockUsecase.On("Get", mock.Anything, mock.Anything).Return([]model.Invitation{}, expectedCount, tt.wantError)

			h := &Handler{
				usecase: &mockUsecase,
			}

			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = httptest.NewRequest(http.MethodGet, "/v1/invitations"+tt.queryParam, nil)
			ctx.Request.Header.Set("Content-Type", "application/json")

			h.Get(ctx)
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
			mockUsecase.On("GetByID", mock.Anything, mock.Anything).Return(model.Invitation{}, tt.wantError)

			h := &Handler{
				usecase: &mockUsecase,
			}

			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = httptest.NewRequest(http.MethodGet, "/v1/invitations/"+tt.param, nil)
			ctx.Request.Header.Set("Content-Type", "application/json")
			ctx.Params = gin.Params{{Key: "id", Value: tt.param}}

			h.GetByID(ctx)
			assert.Equal(t, tt.code, w.Code)
		})
	}
}

func TestUpdate(t *testing.T) {
	gin.SetMode(gin.TestMode)

	testCase := []testCase{
		{
			name: "Testcase #1: Positive", body: payloadSuccess, param: "1", wantError: nil, code: http.StatusOK,
		},
		{
			name: "Testcase #2: Negative", body: payloadSuccess, param: "1", wantError: errFoo, code: http.StatusInternalServerError,
		},
		{
			name: "Testcase #3: Negative", body: payloadFail, param: "1", wantError: errFoo, code: http.StatusUnprocessableEntity,
		},
		{
			name: "Testcase #4: Negative", body: payloadSuccess, param: "one", wantError: errFoo, code: http.StatusUnprocessableEntity,
		},
	}
	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			mockUsecase := mockUsecase.IUsecase{}
			mockUsecase.On("Update", mock.Anything, mock.Anything, mock.Anything).Return(model.Invitation{}, tt.wantError)

			h := &Handler{
				usecase: &mockUsecase,
			}

			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = httptest.NewRequest(http.MethodPatch, "/v1/invitations/"+tt.param, strings.NewReader(tt.body))
			ctx.Request.Header.Set("Content-Type", "application/json")
			ctx.Params = gin.Params{{Key: "id", Value: tt.param}}

			h.Update(ctx)
			assert.EqualValues(t, tt.code, w.Code)
		})
	}
}

func TestGetByMemberID(t *testing.T) {
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
			mockUsecase.On("GetByMemberID", mock.Anything, mock.Anything).Return([]model.InvitationDetail{}, tt.wantError)

			h := &Handler{
				usecase: &mockUsecase,
			}

			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = httptest.NewRequest(http.MethodGet, "/v1/invitations/"+tt.param, nil)
			ctx.Request.Header.Set("Content-Type", "application/json")
			ctx.Params = gin.Params{{Key: "id", Value: tt.param}}

			h.GetByMemberID(ctx)
			assert.Equal(t, tt.code, w.Code)
		})
	}
}
