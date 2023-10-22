package user

import (
	"github.com/gin-gonic/gin"
	"github.com/rzfhlv/gin-example/config"
	"github.com/rzfhlv/gin-example/internal/modules/user/handler"
	"github.com/rzfhlv/gin-example/internal/modules/user/repository"
	"github.com/rzfhlv/gin-example/internal/modules/user/usecase"
	"github.com/rzfhlv/gin-example/middleware/auth"
)

func Mount(route *gin.RouterGroup, h handler.IHandler, a auth.IAuth) (g *gin.RouterGroup) {
	g = route.Group("/users")
	g.POST("/register", h.Register)
	g.POST("/login", h.Login)
	g.POST("/logout", a.Bearer(), h.Logout)
	g.GET("", a.Bearer(), h.GetAll)
	g.GET("/:id", a.Bearer(), h.GetByID)
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
