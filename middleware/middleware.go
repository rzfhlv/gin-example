package middleware

import (
	"github.com/rzfhlv/gin-example/config"
	"github.com/rzfhlv/gin-example/middleware/auth"
)

type Middleware struct {
	Auth auth.IAuth
}

func New(cfg *config.Config) *Middleware {
	auth := auth.New(cfg)

	return &Middleware{
		Auth: auth,
	}
}
