package auth

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/rzfhlv/gin-example/config"
	pJwt "github.com/rzfhlv/gin-example/pkg/jwt"
	"github.com/rzfhlv/gin-example/pkg/message"
	"github.com/rzfhlv/gin-example/pkg/response"
)

var (
	ID       = "id"
	EMAIL    = "email"
	USERNAME = "username"

	BEARER        = "Bearer"
	AUTHORIZATION = "Authorization"

	UNSUPPORTEDTOKENLOG  = "Auth Unsupported Token"
	EMPTYTOKENLOG        = "Auth Empty Token"
	VALIDATIONINVALIDLOG = "Auth Validation Invalid"
	REDISLOG             = "Auth Redis Key Deleted"
)

type IAuth interface {
	Bearer() gin.HandlerFunc
}

type Auth struct {
	redis   *redis.Client
	jwtImpl pJwt.JWTInterface
}

func New(cfg *config.Config) IAuth {
	return &Auth{
		redis:   cfg.Redis,
		jwtImpl: cfg.Pkg.JWTImpl,
	}
}

func (a *Auth) Bearer() gin.HandlerFunc {
	return func(c *gin.Context) {
		split := strings.Split(c.Request.Header.Get(AUTHORIZATION), " ")
		if len(split) < 2 {
			log.Printf(UNSUPPORTEDTOKENLOG+" %v", split)
			c.JSON(http.StatusUnauthorized, response.Set(message.ERROR, message.UNAUTHORIZED, nil, nil))
			c.Abort()
			return
		}

		if split[0] != BEARER {
			log.Printf(UNSUPPORTEDTOKENLOG+" %v", split[0])
			c.JSON(http.StatusUnauthorized, response.Set(message.ERROR, message.UNAUTHORIZED, nil, nil))
			c.Abort()
			return
		}

		if split[1] == "" {
			log.Printf(UNSUPPORTEDTOKENLOG+" %v", split[1])
			c.JSON(http.StatusUnauthorized, response.Set(message.ERROR, message.UNAUTHORIZED, nil, nil))
			c.Abort()
			return
		}

		claims, err := a.jwtImpl.ValidateToken(split[1])
		if err != nil {
			log.Printf(VALIDATIONINVALIDLOG+" %v", err.Error())
			c.JSON(http.StatusUnauthorized, response.Set(message.ERROR, message.UNAUTHORIZED, nil, nil))
			c.Abort()
			return
		}

		err = a.redis.Get(context.Background(), claims.Username).Err()
		if err != nil {
			log.Printf(REDISLOG+" %v", err.Error())
			c.JSON(http.StatusUnauthorized, response.Set(message.ERROR, message.UNAUTHORIZED, nil, nil))
			c.Abort()
			return
		}

		c.Set(ID, claims.ID)
		c.Set(EMAIL, claims.Email)
		c.Set(USERNAME, claims.Username)

		c.Next()
	}
}
