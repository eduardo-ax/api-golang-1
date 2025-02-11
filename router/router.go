package router

import "github.com/gin-gonic/gin"

func Initialize() {

	router := gin.Default()

	router.GET("/user", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Added",
		})
	})

	router.Run(":8080")
}
