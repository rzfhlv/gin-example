package user

import (
	"testing"

	"github.com/gin-gonic/gin"
	aMySQL "github.com/rzfhlv/gin-example/adapter/mysql"
	aRedis "github.com/rzfhlv/gin-example/adapter/redis"
	"github.com/rzfhlv/gin-example/config"
	mockAuth "github.com/rzfhlv/gin-example/shared/mocks/middleware/auth"
	mockHandler "github.com/rzfhlv/gin-example/shared/mocks/modules/user/handler"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	mySql := aMySQL.MySQL{}
	redis := aRedis.Redis{}
	cfg := config.Config{
		MySQL: &mySql,
		Redis: &redis,
	}

	m := New(&cfg)
	assert.NotNil(t, m)
}

func TestMount(t *testing.T) {
	mockHandler := mockHandler.IHandler{}
	mockAuth := mockAuth.IAuth{}

	g := gin.Default()
	route := g.Group("/v1")
	u := Mount(route, &mockHandler, &mockAuth)
	assert.NotNil(t, u)
}
