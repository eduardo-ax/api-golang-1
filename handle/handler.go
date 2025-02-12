package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserHandle(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "GET USER",
	})
}

func ListUsersHandle(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "LIST USERS",
	})
}

func CreateUsersHandle(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "CREATE USERS",
	})
}

func DeleteUserHandle(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "DELETE USER",
	})
}

func UpdateUserHandle(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "MODIFY USER",
	})
}
