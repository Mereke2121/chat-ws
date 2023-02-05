package handlers

import (
	"net/http"

	"go/src/github.com/chat-ws/utils"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type WsResponse struct {
	Type    string   `json:"type"`
	Message string   `json:"message"`
	Users   []string `json:"users"`
}

type WsPayload struct {
	MessageType string `json:"message_type"`
	Username    string `json:"username"`
	Message     string `json:"message"`
}

var users = make(map[*websocket.Conn]string)

func WsConnection(c *gin.Context) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	msg := WsResponse{
		Message: "Websocket connected",
	}
	conn.WriteJSON(msg)

	go ListenWS(c, conn)
}

func ListenWS(c *gin.Context, conn *websocket.Conn) {
	for {
		var payload WsPayload
		err := conn.ReadJSON(&payload)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		if payload.MessageType == utils.UsernameType {
			users[conn] = payload.Username
			listUsers := getAllUsers()
			response := WsResponse{
				Type:  utils.UsernameType,
				Users: listUsers,
			}
			conn.WriteJSON(response)
		}
	}
}

func getAllUsers() []string {
	var result []string
	for _, user := range users {
		result = append(result, user)
	}
	return result
}
