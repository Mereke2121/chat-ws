package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context) {
	if c.Request.URL.Path != "/" {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Not Found")
		return
	}
	if c.Request.Method != http.MethodGet {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Method not allowed")
		return
	}

	c.HTML(http.StatusOK, "home.html", gin.H{
		"title": "Home page",
	})
}
