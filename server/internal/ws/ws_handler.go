package ws

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	// TODO: Change this later to accept the frontend origin only
	CheckOrigin: func(r * http.Request) bool {
		// origin := r.Header.Get("Origin")
		// return origin == "localhost:3000"
		return true
	},
}

type Handler struct {
	Hub *Hub
}

func NewHandler(hub *Hub) *Handler {
	return &Handler{
		Hub: hub,
	}
}

func (h *Handler) CreateRoom(ctx *gin.Context) {
	var req CreateRoomRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest,gin.H{"error": err.Error()})
		return
	}

	h.Hub.Rooms[req.Id] = &Room{
		Id: req.Id,
		Name: req.Name,
		Clients: make(map[string]*Client),
	}
	ctx.JSON(http.StatusOK,req)
}


func (h *Handler) JoinRoom(ctx *gin.Context) {

	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request,nil)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,gin.H{"error": err.Error()})
		return
	}

	roomId := ctx.Param("roomId")
	clientId := ctx.Query("userId")
	username := ctx.Query("username")

	cl := &Client{
		Connection: conn,
		Id: clientId,
		RoomId: roomId,
		Username: username,
		Message: make(chan *Message,10),
	}

	msg := &Message{
		Content: "A new user has joined the room",
		RoomId: roomId,
		UserId: clientId,
		Username: username,
	}

	// register new client through the register channel
	h.Hub.Register <- cl
	// broadcast join message
	h.Hub.Broadcast <- msg

	// run writing on another go routine
	go cl.WriteMessage()
	// run reading messages on the same go routine that accepted the connection
	cl.ReadMessage(h.Hub)
}


