package invitation

import (
	"testing"

	"github.com/gin-gonic/gin"
	aMySQL "github.com/rzfhlv/gin-example/adapter/mysql"
	aRedis "github.com/rzfhlv/gin-example/adapter/redis"
	"github.com/rzfhlv/gin-example/config"
	mockAuth "github.com/rzfhlv/gin-example/shared/mocks/middleware/auth"
	mockHandler "github.com/rzfhlv/gin-example/shared/mocks/modules/invitation/handler"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	mySql := aMySQL.MySQL{}
	redis := aRedis.Redis{}
	cfg := config.Config{
		MySQL: &mySql,
		Redis: &redis,
	}

	i := New(&cfg)
	assert.NotNil(t, i)
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
	m := Mount(route, &mockHandler, &mockAuth)
	assert.NotNil(t, m)
}
