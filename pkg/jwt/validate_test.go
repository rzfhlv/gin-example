package jwt

import (
	"os"
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func TestValidateTokenExp(t *testing.T) {
	os.Setenv("JWT_SECRET", "verysecret")
	os.Setenv("JWT_EXPIRED", "1")

	jwtImpl := JWTImpl{}
	validToken, _ := jwtImpl.Generate(123, "testuser", "test@example.com")

	token, err := jwt.ParseWithClaims(
		validToken,
		&JWTClaim{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	t.Logf("%+v", token.Claims)
	t.Logf("%+v", err)
}

func TestValidateTokenValid(t *testing.T) {
	os.Setenv("JWT_SECRET", "verysecret")
	os.Setenv("JWT_EXPIRED", "1")

	jwtImpl := JWTImpl{}
	validToken, _ := jwtImpl.Generate(123, "testuser", "test@example.com")

	claims, err := jwtImpl.ValidateToken(validToken)
	assert.NoError(t, err)
	assert.NotNil(t, claims)
	assert.Equal(t, int64(123), claims.ID)
	assert.Equal(t, "testuser", claims.Username)
	assert.Equal(t, "test@example.com", claims.Email)
}

func TestValidateTokenInvalid(t *testing.T) {
	invalidToken := "invalidtoken"

	jwtImpl := JWTImpl{}
	claims, err := jwtImpl.ValidateToken(invalidToken)
	assert.Error(t, err)
	assert.Error(t, err)
	assert.Nil(t, claims)
	assert.Equal(t, "token is malformed: token contains an invalid number of segments", err.Error())
}
