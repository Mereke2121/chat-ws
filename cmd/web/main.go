package main

import (
	app "go/src/github.com/chat-ws"
	"go/src/github.com/chat-ws/pkg/handlers"
	"log"
)

func main() {
	server := new(app.Server)
	if err := server.Run(":8000", handlers.InitRoutes()); err != nil {
		log.Fatal(err)
	}
}
