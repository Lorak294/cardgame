package ws

import (
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	Connection 	*websocket.Conn
	Message 	chan *Message
	Id 			string `json:"id"`
	Username 	string `json:"username"`
	RoomId 		string `json:"roomId"`
}

type Message struct {
	Content string `json:"content"`
	RoomId 	string `json:"roomId"`
	UserId 	string `json:"userId"`
	Username 	string `json:"username"`

	// add more fields or just keep everythin as coded content
}


// writes messages incoming from the Message channel to the WebSocket connection in a loop
func (c *Client) WriteMessage() {
	defer func() {
		c.Connection.Close()
	}()

	for {
		message, ok := <-c.Message
		if !ok {
			return
		}

		c.Connection.WriteJSON(message)
	}
}

// reads messages incoming from the WebSocket connection in a loop and writes it to the hub's Broadcast channel
func (c *Client) ReadMessage(hub *Hub) {
	defer func() {
		hub.Unregister <- c
		c.Connection.Close()
	}()	

	for {

		_, m, err := c.Connection.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err,websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v",err)
			}
			break
		}

		msg := &Message{
			Content: string(m),
			RoomId: c.RoomId,	
			Username: c.Username,
			UserId: c.Id,
		}

		hub.Broadcast <- msg
	}
} 