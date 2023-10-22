package internal

import (
	"testing"

	"github.com/rzfhlv/gin-example/config"
	"github.com/stretchr/testify/assert"
)

func TestService(t *testing.T) {
	cfg := config.Config{
		MySQL: nil,
		Redis: nil,
	}

	s := New(&cfg)
	assert.NotNil(t, s)
}
