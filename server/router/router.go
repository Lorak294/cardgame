package router

import (
	"server/internal/user"
	"server/internal/ws"

	"github.com/gin-gonic/gin"
)

var r_eng *gin.Engine

func InitRouter(userHandler *user.Handler, wsHandler *ws.Handler) {
	r_eng = gin.Default()

	// user endpoints
	r_eng.POST("/signup", userHandler.CreateUser)
	r_eng.POST("/login", userHandler.Login)
	r_eng.GET("/logout", userHandler.Logout)

	// ws endpoints
	r_eng.POST("/ws/createRoom",wsHandler.CreateRoom)

}

func StartRouter(addr string) error {
	return r_eng.Run(addr)
}