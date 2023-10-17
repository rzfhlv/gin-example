package usecase

import (
	"testing"

	mockRepo "github.com/rzfhlv/gin-example/shared/mocks/modules/health-check/repository"
	"github.com/stretchr/testify/assert"
)

func TestNewUsecase(t *testing.T) {
	mockRepo := mockRepo.IRepository{}
	u := New(&mockRepo)
	assert.NotNil(t, u)
}
