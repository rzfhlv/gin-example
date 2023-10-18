package healthcheck

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rzfhlv/gin-example/config"
	mockHandler "github.com/rzfhlv/gin-example/shared/mocks/modules/health-check/handler"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type testCase struct {
	name, target string
	wantError    error
	code         int
}

var (
	errFoo = errors.New("error")
)

func TestMount(t *testing.T) {
	testCase := []testCase{
		{
			name: "Testcase #1: Positive", target: "/v1/health-check", wantError: nil, code: http.StatusOK,
		},
		{
			name: "Testcase #2: Negative", target: "/health", wantError: errFoo, code: http.StatusNotFound,
		},
	}
	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			mockHandler := mockHandler.IHandler{}
			mockHandler.On("Ping", mock.Anything).Return(tt.wantError)

			g := gin.Default()
			route := g.Group("/v1")
			Mount(route, &mockHandler)

			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, tt.target, nil)

			g.ServeHTTP(w, req)

			assert.Equal(t, tt.code, w.Code)
		})
	}
}

func TestNew(t *testing.T) {
	cfg := config.Config{
		MySQL: nil,
	}

	hc := New(&cfg)
	assert.NotNil(t, hc)
}
