package user

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rzfhlv/gin-example/config"
	mockAuth "github.com/rzfhlv/gin-example/shared/mocks/middleware/auth"
	mockHandler "github.com/rzfhlv/gin-example/shared/mocks/modules/user/handler"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	cfg := config.Config{
		MySQL: nil,
		Redis: nil,
	}

	m := New(&cfg)
	assert.NotNil(t, m)
}

func TestMount(t *testing.T) {
	mockHandler := mockHandler.IHandler{}
	mockAuth := mockAuth.IAuth{}
	mockAuth.On("Bearer").Return(func() gin.HandlerFunc {
		return func(c *gin.Context) {
			c.Next()
		}
	})

	g := gin.Default()
	route := g.Group("/v1")
	u := Mount(route, &mockHandler, &mockAuth)
	assert.NotNil(t, u)
}
