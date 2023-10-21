package healthcheck

import (
	"github.com/gin-gonic/gin"
	"github.com/rzfhlv/gin-example/config"
	"github.com/rzfhlv/gin-example/internal/modules/health-check/handler"
	"github.com/rzfhlv/gin-example/internal/modules/health-check/repository"
	"github.com/rzfhlv/gin-example/internal/modules/health-check/usecase"
)

func Mount(route *gin.RouterGroup, h handler.IHandler) (g *gin.RouterGroup) {
	g = route.Group("/health-check")
	g.GET("", h.Ping)
	return
}

type HealthCheck struct {
	Handler handler.IHandler
}

func New(cfg *config.Config) *HealthCheck {
	Repo := repository.New(cfg.MySQL.GetDB())
	Usecase := usecase.New(Repo)
	Handler := handler.New(Usecase)

	return &HealthCheck{
		Handler: Handler,
	}
}
