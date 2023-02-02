package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context) {
	c.String(http.StatusOK, "this is the home page")
}
