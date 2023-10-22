package middleware

import (
	"testing"

	"github.com/rzfhlv/gin-example/config"
	pJwt "github.com/rzfhlv/gin-example/pkg/jwt"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	jwtImpl := pJwt.JWTImpl{}

	cfg := config.Config{
		Redis: nil,
		Pkg: config.Pkg{
			JWTImpl: &jwtImpl,
		},
	}

	m := New(&cfg)
	assert.NotNil(t, m)
}
