package routes

import (
	"testing"

	"github.com/rzfhlv/gin-example/config"
	"github.com/rzfhlv/gin-example/internal"
	"github.com/rzfhlv/gin-example/internal/modules/gathering"
	healthcheck "github.com/rzfhlv/gin-example/internal/modules/health-check"
	"github.com/rzfhlv/gin-example/internal/modules/invitation"
	"github.com/rzfhlv/gin-example/internal/modules/member"
	"github.com/stretchr/testify/assert"
)

func TestRoutes(t *testing.T) {
	cfg := config.Config{
		MySQL: nil,
	}
	service := internal.Service{
		HealthCheck: healthcheck.New(&cfg),
		Member:      member.New(&cfg),
		Gathering:   gathering.New(&cfg),
		Invitation:  invitation.New(&cfg),
	}
	r := ListRoutes(&service)
	assert.NotNil(t, r)
}