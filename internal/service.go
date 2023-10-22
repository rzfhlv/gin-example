package internal

import (
	"github.com/rzfhlv/gin-example/config"
	"github.com/rzfhlv/gin-example/internal/modules/gathering"
	healthcheck "github.com/rzfhlv/gin-example/internal/modules/health-check"
	"github.com/rzfhlv/gin-example/internal/modules/invitation"
	"github.com/rzfhlv/gin-example/internal/modules/member"
	"github.com/rzfhlv/gin-example/internal/modules/user"
	"github.com/rzfhlv/gin-example/middleware"
)

type Service struct {
	HealthCheck *healthcheck.HealthCheck
	Member      *member.Member
	Gathering   *gathering.Gathering
	Invitation  *invitation.Invitation
	User        *user.User
	Middleware  *middleware.Middleware
}

func New(cfg *config.Config) *Service {
	healthCheck := healthcheck.New(cfg)
	member := member.New(cfg)
	gathering := gathering.New(cfg)
	invitation := invitation.New(cfg)
	user := user.New(cfg)

	middleware := middleware.New(cfg)

	return &Service{
		HealthCheck: healthCheck,
		Member:      member,
		Gathering:   gathering,
		Invitation:  invitation,
		User:        user,
		Middleware:  middleware,
	}
}
