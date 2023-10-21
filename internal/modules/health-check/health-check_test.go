package healthcheck

import (
	"testing"

	"github.com/gin-gonic/gin"
	aMySQL "github.com/rzfhlv/gin-example/adapter/mysql"
	aRedis "github.com/rzfhlv/gin-example/adapter/redis"
	"github.com/rzfhlv/gin-example/config"
	mockHandler "github.com/rzfhlv/gin-example/shared/mocks/modules/health-check/handler"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	mySql := aMySQL.MySQL{}
	redis := aRedis.Redis{}
	cfg := config.Config{
		MySQL: &mySql,
		Redis: &redis,
	}

	hc := New(&cfg)
	assert.NotNil(t, hc)
}

func TestMount(t *testing.T) {
	mockHandler := mockHandler.IHandler{}

	g := gin.Default()
	route := g.Group("/v1")
	m := Mount(route, &mockHandler)
	assert.NotNil(t, m)
}
