package handler

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rzfhlv/gin-example/internal/modules/invitation/model"
	"github.com/rzfhlv/gin-example/internal/modules/invitation/usecase"
	"github.com/rzfhlv/gin-example/pkg/message"
	"github.com/rzfhlv/gin-example/pkg/param"
	"github.com/rzfhlv/gin-example/pkg/response"
)

type IHandler interface {
	Create(g *gin.Context)
	Get(g *gin.Context)
	GetByID(g *gin.Context)
	Update(g *gin.Context)
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

	invitationPayload := model.Invitation{}
	err := g.ShouldBindJSON(&invitationPayload)
	if err != nil {
		log.Printf("Error Binding and Validation Invitation, %v", err.Error())
		g.JSON(http.StatusUnprocessableEntity, response.Set(message.ERROR, err.Error(), nil, nil))
		return
	}

	invitation, err := h.usecase.Create(ctx, invitationPayload)
	if err != nil {
		log.Printf("Error Create Invitation, %v", err.Error())
		g.JSON(http.StatusInternalServerError, response.Set(message.ERROR, message.SOMETHINGWENTWRONG, nil, nil))
		return
	}

	g.JSON(http.StatusOK, response.Set(message.SUCCESS, message.OK, nil, invitation))
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

	invitations, total, err := h.usecase.Get(ctx, queryParam)
	if err != nil {
		log.Printf("Error Get Invitation, %v", err.Error())
		g.JSON(http.StatusInternalServerError, response.Set(message.ERROR, message.SOMETHINGWENTWRONG, nil, nil))
	}
	queryParam.Total = total
	meta := response.BuildMeta(queryParam, len(invitations))

	g.JSON(http.StatusOK, response.Set(message.SUCCESS, message.OK, meta, invitations))
}

func (h *Handler) GetByID(g *gin.Context) {
	ctx := g.Request.Context()

	id := g.Param("id")
	invitationID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Printf("Error Parse Invitation ID, %v", err.Error())
		g.JSON(http.StatusUnprocessableEntity, response.Set(message.ERROR, message.UNPROCESSABLEENTITY, nil, nil))
		return
	}

	invitation, err := h.usecase.GetByID(ctx, invitationID)
	if err != nil {
		log.Printf("Error Get By ID Invitation, %v", err.Error())
		if err == sql.ErrNoRows {
			g.JSON(http.StatusNotFound, response.Set(message.ERROR, message.NOTFOUND, nil, nil))
			return
		}
		g.JSON(http.StatusInternalServerError, response.Set(message.SUCCESS, message.SOMETHINGWENTWRONG, nil, nil))
		return
	}

	g.JSON(http.StatusOK, response.Set(message.SUCCESS, message.OK, nil, invitation))
}

func (h *Handler) Update(g *gin.Context) {
	ctx := g.Request.Context()

	id := g.Param("id")
	invitationID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Printf("Error Parse Invitation ID, %v", err.Error())
		g.JSON(http.StatusUnprocessableEntity, response.Set(message.ERROR, message.UNPROCESSABLEENTITY, nil, nil))
		return
	}

	invitationPayload := model.Invitation{}
	err = g.ShouldBindJSON(&invitationPayload)
	if err != nil {
		log.Printf("Error Binding and Validation Invitation, %v", err.Error())
		g.JSON(http.StatusUnprocessableEntity, response.Set(message.ERROR, err.Error(), nil, nil))
		return
	}

	invitation, err := h.usecase.Update(ctx, invitationPayload, invitationID)
	if err != nil {
		log.Printf("Error Update Invitation, %v", err.Error())
		g.JSON(http.StatusInternalServerError, response.Set(message.SUCCESS, message.SOMETHINGWENTWRONG, nil, nil))
		return
	}

	g.JSON(http.StatusOK, response.Set(message.SUCCESS, message.OK, nil, invitation))
}
