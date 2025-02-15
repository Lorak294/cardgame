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

func (h *Handler) Login(ctx *gin.Context) {

	// parse request
	var u_req LoginUserRequest
	if err := ctx.ShouldBindJSON(&u_req); err != nil  {
		ctx.JSON(http.StatusBadRequest,gin.H{"error": err.Error()})
		return
	}

	// login
	res, err := h.service.Login(ctx.Request.Context(),&u_req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 
	}

	// set jwt in cookies and clear it from the response body
	ctx.SetCookie("jwt",res.AccessToken,3600,"/","localhost",false,true)
	res = &LoginUserResponse{
		Id: res.Id,
		Username: res.Username,
	}

	ctx.JSON(http.StatusOK,res)
}

func (h *Handler) Logout(ctx *gin.Context) {
	ctx.SetCookie("jwt","",-1,"","",false,true)
	ctx.JSON(http.StatusOK,gin.H{"message": "logout successful"})
}

func NewHandler(s IService) *Handler {
	return &Handler{
		service: s,
	}
}
