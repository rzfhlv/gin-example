package handler

import (
	"database/sql"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rzfhlv/gin-example/internal/modules/gathering/model"
	mockUsecase "github.com/rzfhlv/gin-example/shared/mocks/modules/gathering/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type testCase struct {
	name, body string
	wantError  error
	code       int
}

var (
	errFoo         = errors.New("error")
	payloadSuccess = `{"creator":"john doe","type":"family","name":"family gathering","location":"puncak","schedule_at":"2023-11-10 12:00:00"}`
	payloadFail    = `{"creator":"john doe","type":"","name":"family gathering","location":"puncak","schedule_at":"2023-11-10 12:00:00"}`
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
			mockUsecase.On("Create", mock.Anything, mock.Anything).Return(model.Gathering{}, tt.wantError)

			h := &Handler{
				usecase: &mockUsecase,
			}

			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = httptest.NewRequest(http.MethodPost, "/v1/gatherings", strings.NewReader(tt.body))
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
			name: "Testcase #1: Positive", body: "?page=1", wantError: nil, code: http.StatusOK,
		},
		{
			name: "Testcase #2: Negative", body: "?page=1", wantError: errFoo, code: http.StatusInternalServerError,
		},
		{
			name: "Testcase #3: Negative", body: "?page=one", wantError: errFoo, code: http.StatusUnprocessableEntity,
		},
	}
	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			mockUsecase := mockUsecase.IUsecase{}
			mockUsecase.On("Get", mock.Anything, mock.Anything).Return([]model.Gathering{}, expectedCount, tt.wantError)

			h := &Handler{
				usecase: &mockUsecase,
			}

			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = httptest.NewRequest(http.MethodGet, "/v1/gatherings"+tt.body, nil)
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
			name: "Testcase #1: Positive", body: "1", wantError: nil, code: http.StatusOK,
		},
		{
			name: "Testcase #2: Negative", body: "1", wantError: errFoo, code: http.StatusInternalServerError,
		},
		{
			name: "Testcase #3: Negative", body: "one", wantError: errFoo, code: http.StatusUnprocessableEntity,
		},
		{
			name: "Testcase #3: Negative", body: "0", wantError: sql.ErrNoRows, code: http.StatusNotFound,
		},
	}
	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			mockUsecase := mockUsecase.IUsecase{}
			mockUsecase.On("GetByID", mock.Anything, mock.Anything).Return(model.Gathering{}, tt.wantError)

			h := &Handler{
				usecase: &mockUsecase,
			}

			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = httptest.NewRequest(http.MethodGet, "/v1/gatherings/"+tt.body, nil)
			ctx.Request.Header.Set("Content-Type", "application/json")
			ctx.Params = gin.Params{{Key: "id", Value: tt.body}}

			h.GetByID(ctx)
			assert.Equal(t, tt.code, w.Code)
		})
	}
}

func TestGetDetailByID(t *testing.T) {
	gin.SetMode(gin.TestMode)

	testCase := []testCase{
		{
			name: "Testcase #1: Positive", body: "1", wantError: nil, code: http.StatusOK,
		},
		{
			name: "Testcase #2: Negative", body: "1", wantError: errFoo, code: http.StatusInternalServerError,
		},
		{
			name: "Testcase #3: Negative", body: "one", wantError: errFoo, code: http.StatusUnprocessableEntity,
		},
		{
			name: "Testcase #3: Negative", body: "0", wantError: sql.ErrNoRows, code: http.StatusNotFound,
		},
	}
	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			mockUsecase := mockUsecase.IUsecase{}
			mockUsecase.On("GetDetailByID", mock.Anything, mock.Anything).Return(model.GatheringDetail{}, tt.wantError)

			h := &Handler{
				usecase: &mockUsecase,
			}

			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = httptest.NewRequest(http.MethodGet, "/v1/gatherings/"+tt.body, nil)
			ctx.Request.Header.Set("Content-Type", "application/json")
			ctx.Params = gin.Params{{Key: "id", Value: tt.body}}

			h.GetDetailByID(ctx)
			assert.Equal(t, tt.code, w.Code)
		})
	}
}
