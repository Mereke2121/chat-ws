package main

import (
	app "go/src/github.com/chat-ws"
	"go/src/github.com/chat-ws/pkg/handlers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	server := new(app.Server)
	if err := server.Run(":8000", initRoutes()); err != nil {
		log.Fatal(err)
	}
}

func initRoutes() *gin.Engine {
	routes := gin.New()
	routes.LoadHTMLGlob("./html/*")

	routes.GET("/", handlers.HomePage)
	routes.GET("/ws", handlers.WsConnection)

	return routes
}
