package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rzfhlv/gin-example/internal/modules/health-check/usecase"
	"github.com/rzfhlv/gin-example/pkg/message"
	"github.com/rzfhlv/gin-example/pkg/response"
)

type IHandler interface {
	Ping(g *gin.Context)
}

type Handler struct {
	usecase usecase.IUsecase
}

func New(usecase usecase.IUsecase) IHandler {
	return &Handler{
		usecase: usecase,
	}
}

func (h *Handler) Ping(g *gin.Context) {
	ctx := g.Request.Context()
	err := h.usecase.Ping(ctx)
	if err != nil {
		log.Printf("Error Ping %v", err.Error())
		g.JSON(http.StatusInternalServerError, response.Set(message.ERROR, message.SOMETHINGWENTWRONG, nil, nil))
		return
	}
	g.JSON(http.StatusOK, response.Set(message.SUCCESS, message.HEALTHCHECK, nil, nil))
}
