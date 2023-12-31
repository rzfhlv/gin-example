package handler

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rzfhlv/gin-example/internal/modules/member/model"
	"github.com/rzfhlv/gin-example/internal/modules/member/usecase"
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

	memberPayload := model.Member{}
	err := g.ShouldBindJSON(&memberPayload)
	if err != nil {
		log.Printf("Error Binding and Validation Member, %v", err.Error())
		g.JSON(http.StatusUnprocessableEntity, response.Set(message.ERROR, err.Error(), nil, nil))
		return
	}
	memberPayload.CreatedAt = time.Now()

	member, err := h.usecase.Create(ctx, memberPayload)
	if err != nil {
		log.Printf("Error Create Member, %v", err.Error())
		g.JSON(http.StatusInternalServerError, response.Set(message.ERROR, message.SOMETHINGWENTWRONG, nil, nil))
		return
	}

	g.JSON(http.StatusOK, response.Set(message.SUCCESS, message.OK, nil, member))
}

func (h *Handler) Get(g *gin.Context) {
	ctx := g.Request.Context()
	queryParam := param.Param{}
	queryParam.Limit = param.DEFAULTLIMIT
	queryParam.Page = param.DEFAULTPAGE

	err := g.ShouldBind(&queryParam)
	if err != nil {
		log.Printf("Error Binding Query Param Gathering, %v", err.Error())
		g.JSON(http.StatusUnprocessableEntity, response.Set(message.ERROR, message.UNPROCESSABLEENTITY, nil, nil))
		return
	}

	members, total, err := h.usecase.Get(ctx, queryParam)
	if err != nil {
		log.Printf("Error Get Member, %v", err.Error())
		g.JSON(http.StatusInternalServerError, response.Set(message.ERROR, message.SOMETHINGWENTWRONG, nil, nil))
		return
	}
	queryParam.Total = total
	meta := response.BuildMeta(queryParam, len(members))

	g.JSON(http.StatusOK, response.Set(message.SUCCESS, message.OK, meta, members))
}

func (h *Handler) GetByID(g *gin.Context) {
	ctx := g.Request.Context()

	id := g.Param("id")
	memberID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Printf("Error Parse Member ID, %v", err.Error())
		g.JSON(http.StatusUnprocessableEntity, response.Set(message.ERROR, message.UNPROCESSABLEENTITY, nil, nil))
		return
	}

	member, err := h.usecase.GetByID(ctx, memberID)
	if err != nil {
		log.Printf("Error Get By ID Member, %v", err.Error())
		if err == sql.ErrNoRows {
			g.JSON(http.StatusNotFound, response.Set(message.ERROR, message.NOTFOUND, nil, nil))
			return
		}
		g.JSON(http.StatusInternalServerError, response.Set(message.SUCCESS, message.SOMETHINGWENTWRONG, nil, nil))
		return
	}

	g.JSON(http.StatusOK, response.Set(message.SUCCESS, message.OK, nil, member))
}
