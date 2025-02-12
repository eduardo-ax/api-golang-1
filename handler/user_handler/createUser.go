package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUsersHandle(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "CREATE USERS",
	})
}
