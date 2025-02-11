package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/user/:id", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "LIST USER",
			})
		})
		v1.GET("/user", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "LIST USERS",
			})
		})
		v1.POST("/user", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "CREATE USER",
			})
		})
		v1.DELETE("/user", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "DELETE USER",
			})
		})
		v1.PUT("/user", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "MODIFY USER",
			})
		})
	}
}
