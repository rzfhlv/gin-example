package handler

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rzfhlv/gin-example/internal/modules/user/model"
	"github.com/rzfhlv/gin-example/internal/modules/user/usecase"
	"github.com/rzfhlv/gin-example/pkg/message"
	"github.com/rzfhlv/gin-example/pkg/param"
	"github.com/rzfhlv/gin-example/pkg/response"
)

type IHandler interface {
	Register(g *gin.Context)
	Login(g *gin.Context)
	Logout(g *gin.Context)
	GetAll(g *gin.Context)
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

func (h *Handler) Register(g *gin.Context) {
	ctx := g.Request.Context()

	register := model.Register{}
	err := g.ShouldBindJSON(&register)
	if err != nil {
		log.Printf("Error Binding and Validation Register, %v", err.Error())
		g.JSON(http.StatusUnprocessableEntity, response.Set(message.ERROR, err.Error(), nil, nil))
		return
	}
	register.CreatedAt = time.Now()

	jwt, err := h.usecase.Register(ctx, register)
	if err != nil {
		log.Printf("Error Register User, %v", err.Error())
		g.JSON(http.StatusInternalServerError, response.Set(message.ERROR, message.SOMETHINGWENTWRONG, nil, nil))
		return
	}

	g.JSON(http.StatusOK, response.Set(message.SUCCESS, message.OK, nil, jwt))
}

func (h *Handler) Login(g *gin.Context) {
	ctx := g.Request.Context()

	login := model.Login{}
	err := g.ShouldBindJSON(&login)
	if err != nil {
		log.Printf("Error Binding and Validation Login, %v", err.Error())
		g.JSON(http.StatusUnprocessableEntity, response.Set(message.ERROR, err.Error(), nil, nil))
		return
	}

	jwt, err := h.usecase.Login(ctx, login)
	if err != nil {
		log.Printf("Error Login User, %v", err.Error())
		g.JSON(http.StatusInternalServerError, response.Set(message.ERROR, message.SOMETHINGWENTWRONG, nil, nil))
		return
	}

	g.JSON(http.StatusOK, response.Set(message.SUCCESS, message.OK, nil, jwt))
}

func (h *Handler) Logout(g *gin.Context) {
	ctx := g.Request.Context()

	auth := strings.Split(g.Request.Header.Get("Authorization"), " ")
	if len(auth) < 2 {
		log.Printf("Error Logout User, %v", auth)
		g.JSON(http.StatusUnauthorized, response.Set(message.ERROR, message.UNAUTHORIZED, nil, nil))
		return
	}

	value, ok := g.Get("username")
	if !ok {
		log.Printf("Error Get Context Username %v and %v", value, ok)
		g.JSON(http.StatusUnauthorized, response.Set(message.ERROR, message.UNAUTHORIZED, nil, nil))
		return
	}
	username := fmt.Sprintf("%v", value)
	err := h.usecase.Logout(ctx, username)
	if err != nil {
		log.Printf("Error Logout User, %v", err.Error())
		g.JSON(http.StatusInternalServerError, response.Set(message.ERROR, message.SOMETHINGWENTWRONG, nil, nil))
		return
	}

	g.JSON(http.StatusOK, response.Set(message.SUCCESS, message.OK, nil, nil))
}

func (h *Handler) GetAll(g *gin.Context) {
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

	users, total, err := h.usecase.GetAll(ctx, queryParam)
	if err != nil {
		log.Printf("Error Get Member, %v", err.Error())
		g.JSON(http.StatusInternalServerError, response.Set(message.ERROR, message.SOMETHINGWENTWRONG, nil, nil))
		return
	}
	queryParam.Total = total
	meta := response.BuildMeta(queryParam, len(users))

	g.JSON(http.StatusOK, response.Set(message.SUCCESS, message.OK, meta, users))
}

func (h *Handler) GetByID(g *gin.Context) {
	ctx := g.Request.Context()

	id := g.Param("id")
	userID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Printf("Error Parse Member ID, %v", err.Error())
		g.JSON(http.StatusUnprocessableEntity, response.Set(message.ERROR, message.UNPROCESSABLEENTITY, nil, nil))
		return
	}

	user, err := h.usecase.GetByID(ctx, userID)
	if err != nil {
		log.Printf("Error Get By ID Member, %v", err.Error())
		if err == sql.ErrNoRows {
			g.JSON(http.StatusNotFound, response.Set(message.ERROR, message.NOTFOUND, nil, nil))
			return
		}
		g.JSON(http.StatusInternalServerError, response.Set(message.SUCCESS, message.SOMETHINGWENTWRONG, nil, nil))
		return
	}

	g.JSON(http.StatusOK, response.Set(message.SUCCESS, message.OK, nil, user))
}
