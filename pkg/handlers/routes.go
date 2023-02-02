package handlers

import "github.com/gin-gonic/gin"

func InitRoutes() *gin.Engine {
	routes := gin.Default()

	routes.GET("/", HomePage)

	return routes
}
