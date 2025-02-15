package ws

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Hub *Hub
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

func NewHandler(hub *Hub) *Handler {
	return &Handler{
		Hub: hub,
	}
}
