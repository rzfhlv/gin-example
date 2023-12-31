package invitation

import (
	"github.com/gin-gonic/gin"
	"github.com/rzfhlv/gin-example/config"
	"github.com/rzfhlv/gin-example/internal/modules/invitation/handler"
	"github.com/rzfhlv/gin-example/internal/modules/invitation/repository"
	"github.com/rzfhlv/gin-example/internal/modules/invitation/usecase"
	"github.com/rzfhlv/gin-example/middleware/auth"
)

func Mount(route *gin.RouterGroup, h handler.IHandler, a auth.IAuth) (g *gin.RouterGroup) {
	g = route.Group("/invitations")
	g.Use(a.Bearer())
	g.GET("", h.Get)
	g.GET("/:id", h.GetByID)
	g.POST("", h.Create)
	g.PATCH("/:id", h.Update)
	g.GET("/me/:id", h.GetByMemberID)
	return
}

type Invitation struct {
	Handler handler.IHandler
}

func New(cfg *config.Config) *Invitation {
	Repo := repository.New(cfg.MySQL)
	Usecase := usecase.New(Repo)
	Handler := handler.New(Usecase)

	return &Invitation{
		Handler: Handler,
	}
}
