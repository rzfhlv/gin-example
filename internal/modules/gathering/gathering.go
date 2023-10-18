package gathering

import (
	"github.com/gin-gonic/gin"
	"github.com/rzfhlv/gin-example/config"
	"github.com/rzfhlv/gin-example/internal/modules/gathering/handler"
	"github.com/rzfhlv/gin-example/internal/modules/gathering/repository"
	"github.com/rzfhlv/gin-example/internal/modules/gathering/usecase"
)

func Mount(route *gin.RouterGroup, h handler.IHandler) (g *gin.RouterGroup) {
	g = route.Group("/gatherings")
	g.GET("", h.Get)
	g.GET("/:id", h.GetByID)
	g.POST("", h.Create)
	g.GET("/:id/detail", h.GetDetailByID)
	return
}

type Gathering struct {
	Handler handler.IHandler
}

func New(cfg *config.Config) *Gathering {
	Repo := repository.New(cfg.MySQL)
	Usecase := usecase.New(Repo)
	Handler := handler.New(Usecase)

	return &Gathering{
		Handler: Handler,
	}
}
