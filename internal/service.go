package internal

import (
	"github.com/rzfhlv/gin-example/config"
	healthcheck "github.com/rzfhlv/gin-example/internal/modules/health-check"
	"github.com/rzfhlv/gin-example/internal/modules/member"
)

type Service struct {
	HealthCheck *healthcheck.HealthCheck
	Member      *member.Member
}

func New(cfg *config.Config) *Service {
	healthCheck := healthcheck.New(cfg)
	member := member.New(cfg)

	return &Service{
		HealthCheck: healthCheck,
		Member:      member,
	}
}
