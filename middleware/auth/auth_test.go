package auth

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redismock/v9"
	"github.com/rzfhlv/gin-example/config"
	pJwt "github.com/rzfhlv/gin-example/pkg/jwt"
	"github.com/stretchr/testify/assert"
)

var (
	_ = os.Setenv("JWT_SECRET", "verysecret")
	_ = os.Setenv("JWT_EXPIRED", "1")
)

func TestAuthSuccess(t *testing.T) {
	gin.SetMode(gin.TestMode)

	jwtImpl := pJwt.JWTImpl{}
	token, _ := jwtImpl.Generate(int64(1), "johndoe", "johndoe@test.com")

	client, mock := redismock.NewClientMock()
	mock.ExpectGet("johndoe").SetVal(token)

	cfg := config.Config{
		Redis: client,
		Pkg: config.Pkg{
			JWTImpl: &jwtImpl,
		},
	}

	g := gin.Default()
	auth := New(&cfg)
	g.Use(auth.Bearer())
	g.GET("/v1/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, nil)
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/v1/users", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	g.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestAuthFail(t *testing.T) {
	gin.SetMode(gin.TestMode)

	jwtImpl := pJwt.JWTImpl{}
	token := "invalidtoken"

	client, mock := redismock.NewClientMock()
	mock.ExpectGet("johndoe").SetErr(errors.New("error"))

	cfg := config.Config{
		Redis: client,
		Pkg: config.Pkg{
			JWTImpl: &jwtImpl,
		},
	}

	g := gin.Default()
	auth := New(&cfg)
	g.Use(auth.Bearer())
	g.GET("/v1/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, nil)
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/v1/users", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	g.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestAuthFailHeader(t *testing.T) {
	gin.SetMode(gin.TestMode)

	jwtImpl := pJwt.JWTImpl{}

	client, mock := redismock.NewClientMock()
	mock.ExpectGet("johndoe").SetErr(errors.New("error"))

	cfg := config.Config{
		Redis: client,
		Pkg: config.Pkg{
			JWTImpl: &jwtImpl,
		},
	}

	g := gin.Default()
	auth := New(&cfg)
	g.Use(auth.Bearer())
	g.GET("/v1/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, nil)
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/v1/users", nil)
	req.Header.Set("Authorization", "")
	g.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestAuthFailHeaderBearer(t *testing.T) {
	gin.SetMode(gin.TestMode)

	jwtImpl := pJwt.JWTImpl{}

	client, mock := redismock.NewClientMock()
	mock.ExpectGet("johndoe").SetErr(errors.New("error"))

	cfg := config.Config{
		Redis: client,
		Pkg: config.Pkg{
			JWTImpl: &jwtImpl,
		},
	}

	g := gin.Default()
	auth := New(&cfg)
	g.Use(auth.Bearer())
	g.GET("/v1/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, nil)
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/v1/users", nil)
	req.Header.Set("Authorization", "Basic Abc")
	g.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestAuthFailEmpty(t *testing.T) {
	gin.SetMode(gin.TestMode)

	jwtImpl := pJwt.JWTImpl{}

	client, mock := redismock.NewClientMock()
	mock.ExpectGet("johndoe").SetErr(errors.New("error"))

	cfg := config.Config{
		Redis: client,
		Pkg: config.Pkg{
			JWTImpl: &jwtImpl,
		},
	}

	g := gin.Default()
	auth := New(&cfg)
	g.Use(auth.Bearer())
	g.GET("/v1/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, nil)
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/v1/users", nil)
	req.Header.Set("Authorization", "Bearer ")
	g.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestAuthFailRedis(t *testing.T) {
	gin.SetMode(gin.TestMode)

	jwtImpl := pJwt.JWTImpl{}
	token, _ := jwtImpl.Generate(int64(1), "johndoe", "johndoe@test.com")

	client, mock := redismock.NewClientMock()
	mock.ExpectGet("johndoe").SetErr(errors.New("error"))

	cfg := config.Config{
		Redis: client,
		Pkg: config.Pkg{
			JWTImpl: &jwtImpl,
		},
	}

	g := gin.Default()
	auth := New(&cfg)
	g.Use(auth.Bearer())
	g.GET("/v1/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, nil)
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/v1/users", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	g.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}
