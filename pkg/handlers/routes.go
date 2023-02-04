package handlers

import "github.com/gin-gonic/gin"

func InitRoutes() *gin.Engine {
	routes := gin.New()
	routes.LoadHTMLGlob("html/*")

	routes.GET("/", HomePage)

	return routes
}
