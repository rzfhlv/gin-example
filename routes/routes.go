package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rzfhlv/gin-example/internal"
	"github.com/rzfhlv/gin-example/internal/modules/gathering"
	healthcheck "github.com/rzfhlv/gin-example/internal/modules/health-check"
	"github.com/rzfhlv/gin-example/internal/modules/invitation"
	"github.com/rzfhlv/gin-example/internal/modules/member"
)

func ListRoutes(svc *internal.Service) (g *gin.Engine) {
	g = gin.Default()

	route := g.Group("/v1")

	healthcheck.Mount(route, svc.HealthCheck.Handler)
	member.Mount(route, svc.Member.Handler)
	gathering.Mount(route, svc.Gathering.Handler)
	invitation.Mount(route, svc.Invitation.Handler)
	return
}
