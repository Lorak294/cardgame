package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service IService
}

func (h *Handler) CreateUser(ctx *gin.Context) {

	// parse request
	var u_req CreateUserRequest
	if err := ctx.ShouldBindJSON(&u_req); err != nil  {
		ctx.JSON(http.StatusBadRequest,gin.H{"error": err.Error()})
		return
	}

	// create user 
	res, err := h.service.CreateUser(ctx.Request.Context(),&u_req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 
	}

	ctx.JSON(http.StatusOK,res)
}

func NewHandler(s IService) *Handler {
	return &Handler{
		service: s,
	}
}
