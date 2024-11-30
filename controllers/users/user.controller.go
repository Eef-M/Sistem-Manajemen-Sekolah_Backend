package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Testing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "its work",
	})
}
