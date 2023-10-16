package internal

import (
	"github.com/rzfhlv/gin-example/config"
	"github.com/rzfhlv/gin-example/internal/modules/gathering"
	healthcheck "github.com/rzfhlv/gin-example/internal/modules/health-check"
	"github.com/rzfhlv/gin-example/internal/modules/invitation"
	"github.com/rzfhlv/gin-example/internal/modules/member"
)

type Service struct {
	HealthCheck *healthcheck.HealthCheck
	Member      *member.Member
	Gathering   *gathering.Gathering
	Invitation  *invitation.Invitation
}

func New(cfg *config.Config) *Service {
	healthCheck := healthcheck.New(cfg)
	member := member.New(cfg)
	gathering := gathering.New(cfg)
	invitation := invitation.New(cfg)

	return &Service{
		HealthCheck: healthCheck,
		Member:      member,
		Gathering:   gathering,
		Invitation:  invitation,
	}
}
