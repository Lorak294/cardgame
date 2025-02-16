package router

import (
	"server/internal/user"
	"server/internal/ws"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var r_eng *gin.Engine

func InitRouter(userHandler *user.Handler, wsHandler *ws.Handler) {
	r_eng = gin.Default()

	// set up cors
	r_eng.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"POST","GET"},
		AllowHeaders: []string{"Content-Type"},
		ExposeHeaders: []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func (origin string) bool  {
			return origin == "http://localhost:3000"
		},
		MaxAge: 12 * time.Hour,
	}))

	// user endpoints
	r_eng.POST("/signup", userHandler.CreateUser)
	r_eng.POST("/login", userHandler.Login)
	r_eng.GET("/logout", userHandler.Logout)

	// ws endpoints
	r_eng.POST("/ws/createRoom",wsHandler.CreateRoom)
	r_eng.GET("/ws/joinRoom/:roomId",wsHandler.JoinRoom)
	r_eng.GET("/ws/getRooms",wsHandler.GetRooms)
	r_eng.GET("/ws/getClients/:roomId",wsHandler.GetClients)

}

func StartRouter(addr string) error {
	return r_eng.Run(addr)
}