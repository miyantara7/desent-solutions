package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GenerateToken(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"token": "secret-token",
	})
}
