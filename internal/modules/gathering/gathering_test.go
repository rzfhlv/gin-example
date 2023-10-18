package gathering

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rzfhlv/gin-example/config"
	mockHandler "github.com/rzfhlv/gin-example/shared/mocks/modules/gathering/handler"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	cfg := config.Config{
		MySQL: nil,
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
