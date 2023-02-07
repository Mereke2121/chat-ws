package handlers

import (
	"fmt"
	"net/http"

	"go/src/github.com/chat-ws/utils"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

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

	msg := utils.WsMessageResponse{
		Message: "Websocket connected",
	}
	conn.WriteJSON(msg)

	go ListenWS(c, conn)
}

func ListenWS(c *gin.Context, conn *websocket.Conn) {
	for {
		var payload utils.WsPayload
		err := conn.ReadJSON(&payload)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		switch payload.MessageType {
		case utils.UsernameType: // добавляем user-a
			users[conn] = payload.Username
			listUsers := getAllUsers(users)
			response := utils.WsUserResponse{
				Type:  utils.UsernameType,
				Users: listUsers,
			}
			if err := conn.WriteJSON(response); err != nil {
				c.AbortWithError(http.StatusInternalServerError, err)
				return
			}
		case utils.MessageType: // отправляем message в чат
			message := fmt.Sprintf("<strong>%s</strong>: %s", payload.Username, payload.Message)
			response := utils.WsMessageResponse{
				UserName: payload.Username,
				Type:     utils.MessageType,
				Message:  message,
			}
			for user := range users {
				if err := user.WriteJSON(response); err != nil {
					c.AbortWithError(http.StatusInternalServerError, err)
					return
				}
			}
		}
	}
}

func getAllUsers(users map[*websocket.Conn]string) []string {
	var result []string
	for _, user := range users {
		result = append(result, user)
	}
	return result
}
