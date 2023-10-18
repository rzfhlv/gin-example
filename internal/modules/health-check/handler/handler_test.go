package handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	mockUsecase "github.com/rzfhlv/gin-example/shared/mocks/modules/health-check/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type testCase struct {
	name      string
	wantError error
	code      int
}

var (
	errFoo = errors.New("error")
)

func TestNew(t *testing.T) {
	mockUsecase := mockUsecase.IUsecase{}

	h := New(&mockUsecase)
	assert.NotNil(t, h)
}

func TestPing(t *testing.T) {
	gin.SetMode(gin.TestMode)

	testCase := []testCase{
		{
			name: "Testcase #1: Positive", wantError: nil, code: http.StatusOK,
		},
		{
			name: "Testcase #2: Negative", wantError: errFoo, code: http.StatusInternalServerError,
		},
	}
	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			mockUsecase := mockUsecase.IUsecase{}
			mockUsecase.On("Ping", mock.Anything).Return(tt.wantError)

			h := &Handler{
				usecase: &mockUsecase,
			}

			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = &http.Request{
				Header: make(http.Header),
			}

			h.Ping(ctx)
			assert.EqualValues(t, tt.code, w.Code)
		})
	}
}
