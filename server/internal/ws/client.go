package ws

import "github.com/gorilla/websocket"

type Client struct {
	Connection 	*websocket.Conn
	Message 	chan*Message
	Id 			string `json:"id"`
	RoomId 		string `json:"roomId"`
}

type Message struct {
	Content string `json:"content"`
	RoomId 	string `json:"roomId"`
	UserId 	string `json:"userId"`

	// add more fields or just keep everythin as coded content
}