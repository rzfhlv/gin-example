package user

import (
	"github.com/gin-gonic/gin"
	"github.com/rzfhlv/gin-example/config"
	"github.com/rzfhlv/gin-example/internal/modules/user/handler"
	"github.com/rzfhlv/gin-example/internal/modules/user/repository"
	"github.com/rzfhlv/gin-example/internal/modules/user/usecase"
)

func Mount(route *gin.RouterGroup, h handler.IHandler) (g *gin.RouterGroup) {
	g = route.Group("/users")
	g.POST("/register", h.Register)
	g.POST("/login", h.Login)
	g.POST("/logout", h.Logout)
	g.GET("", h.GetAll)
	g.GET("/:id", h.GetByID)
	return
}

type User struct {
	Handler handler.IHandler
}

func New(cfg *config.Config) *User {
	Repo := repository.New(cfg.MySQL.GetDB(), cfg.Redis.GetClient())
	Usecase := usecase.New(Repo, cfg.Pkg.Hasher, cfg.Pkg.JWTImpl)
	Handler := handler.New(Usecase)

	return &User{
		Handler: Handler,
	}
}
