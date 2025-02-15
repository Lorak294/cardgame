package ws

type Room struct {
	Id      string             `json:"id"`
	Name    string             `json:"name"`
	Clients map[string]*Client `json:"clients"`
}

type CreateRoomRequest struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Hub struct {
	Rooms      map[string]*Room
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan *Message
}

func NewHub() *Hub {
	return &Hub{
		Rooms:      make(map[string]*Room),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan *Message, 5),
	}
}

// Run function that runs on a separate Go routine
func (h *Hub) Run() {
	for {
		select {
		case cl := <-h.Register:
			// check if room exists
			if _, ok := h.Rooms[cl.RoomId]; ok {
				room := h.Rooms[cl.RoomId]
				// add client too the room if not already added
				if _, exists := room.Clients[cl.Id]; !exists {
					room.Clients[cl.Id] = cl
				}
			}
		case cl := <-h.Unregister:
			// check if room exists
			if _, ok := h.Rooms[cl.RoomId]; ok {
				room := h.Rooms[cl.RoomId]
				// remove the client and close its message channel
				if _, exists := room.Clients[cl.Id]; exists {
					// broadcast msg about client leaving the room
					if len(h.Rooms[cl.RoomId].Clients) != 0 {
						h.Broadcast <- &Message{
							Content:  "User left the room",
							RoomId:   cl.RoomId,
							UserId:   cl.Id,
							Username: cl.Username,
						}
					}
					delete(h.Rooms[cl.RoomId].Clients, cl.Id)
					close(cl.Message)
				}
			}
		case msg := <-h.Broadcast:
			// check if room exists
			if _, ok := h.Rooms[msg.RoomId]; ok {
				// send msg to each client in the room
				for _, cl := range h.Rooms[msg.RoomId].Clients {
					cl.Message <- msg
				}
			}
		}
	}
}