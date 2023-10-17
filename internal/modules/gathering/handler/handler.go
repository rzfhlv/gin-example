package handler

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rzfhlv/gin-example/internal/modules/gathering/model"
	"github.com/rzfhlv/gin-example/internal/modules/gathering/usecase"
	"github.com/rzfhlv/gin-example/pkg/message"
	"github.com/rzfhlv/gin-example/pkg/param"
	"github.com/rzfhlv/gin-example/pkg/response"
)

type IHandler interface {
	Create(g *gin.Context)
	Get(g *gin.Context)
	GetByID(g *gin.Context)
}

type Handler struct {
	usecase usecase.IUsecase
}

func New(usecase usecase.IUsecase) IHandler {
	return &Handler{
		usecase: usecase,
	}
}

func (h *Handler) Create(g *gin.Context) {
	ctx := g.Request.Context()

	gatheringPayload := model.Gathering{}
	err := g.ShouldBindJSON(&gatheringPayload)
	if err != nil {
		log.Printf("Error Binding and Validation Gathering, %v", err.Error())
		g.JSON(http.StatusUnprocessableEntity, response.Set(message.ERROR, err.Error(), nil, nil))
		return
	}

	gathering, err := h.usecase.Create(ctx, gatheringPayload)
	if err != nil {
		log.Printf("Error Create Gathering, %v", err.Error())
		g.JSON(http.StatusInternalServerError, response.Set(message.ERROR, message.SOMETHINGWENTWRONG, nil, nil))
		return
	}

	g.JSON(http.StatusOK, response.Set(message.SUCCESS, message.OK, nil, gathering))
}

func (h *Handler) Get(g *gin.Context) {
	ctx := g.Request.Context()
	queryParam := param.Param{}
	queryParam.Limit = param.DEFAULTLIMIT
	queryParam.Offset = param.DEFAULTPAGE

	err := g.ShouldBind(&queryParam)
	if err != nil {
		log.Printf("Error Binding Query Param Gathering, %v", err.Error())
		g.JSON(http.StatusUnprocessableEntity, response.Set(message.ERROR, message.UNPROCESSABLEENTITY, nil, nil))
		return
	}

	gatherings, total, err := h.usecase.Get(ctx, queryParam)
	if err != nil {
		log.Printf("Error Get Gathering, %v", err.Error())
		g.JSON(http.StatusInternalServerError, response.Set(message.ERROR, message.SOMETHINGWENTWRONG, nil, nil))
	}
	queryParam.Total = total
	meta := response.BuildMeta(queryParam, len(gatherings))

	g.JSON(http.StatusOK, response.Set(message.SUCCESS, message.OK, meta, gatherings))
}

func (h *Handler) GetByID(g *gin.Context) {
	ctx := g.Request.Context()

	id := g.Param("id")
	gatheringID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Printf("Error Parse Gathering ID, %v", err.Error())
		g.JSON(http.StatusUnprocessableEntity, response.Set(message.ERROR, message.UNPROCESSABLEENTITY, nil, nil))
		return
	}

	gathering, err := h.usecase.GetByID(ctx, gatheringID)
	if err != nil {
		log.Printf("Error Get By ID Gathering, %v", err.Error())
		if err == sql.ErrNoRows {
			g.JSON(http.StatusNotFound, response.Set(message.ERROR, message.NOTFOUND, nil, nil))
			return
		}
		g.JSON(http.StatusInternalServerError, response.Set(message.SUCCESS, message.SOMETHINGWENTWRONG, nil, nil))
		return
	}

	g.JSON(http.StatusOK, response.Set(message.SUCCESS, message.OK, nil, gathering))
}
