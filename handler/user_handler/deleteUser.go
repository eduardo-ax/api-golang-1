package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteUserHandle(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "DELETE USER",
	})
}
