package member

import (
	"github.com/gin-gonic/gin"
	"github.com/rzfhlv/gin-example/config"
	"github.com/rzfhlv/gin-example/internal/modules/member/handler"
	"github.com/rzfhlv/gin-example/internal/modules/member/repository"
	"github.com/rzfhlv/gin-example/internal/modules/member/usecase"
)

func Mount(route *gin.RouterGroup, h handler.IHandler) (g *gin.RouterGroup) {
	g = route.Group("/members")
	g.GET("", h.Get)
	g.GET("/:id", h.GetByID)
	g.POST("", h.Create)
	return
}

type Member struct {
	Handler handler.IHandler
}

func New(cfg *config.Config) *Member {
	Repo := repository.New(cfg.MySQL)
	Usecase := usecase.New(Repo, cfg.Pkg.Hasher)
	Handler := handler.New(Usecase)

	return &Member{
		Handler: Handler,
	}
}
