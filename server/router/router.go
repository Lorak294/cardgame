package router

import (
	"server/internal/user"

	"github.com/gin-gonic/gin"
)

var r_eng *gin.Engine

func InitRouter(userHandler *user.Handler) {
	r_eng = gin.Default()

	r_eng.POST("/signup", userHandler.CreateUser)
}

func StartRouter(addr string) error {
	return r_eng.Run(addr)
}