package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	r := gin.Default()

	r.POST("/user", func(c *gin.Context) {
		var data User

		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		println("Name:", data.Name)
		println("Email:", data.Email)

		c.JSON(http.StatusOK, gin.H{
			"message": "Data received!",
			"data":    data,
		})
	})

	r.Run(":8080")
}
