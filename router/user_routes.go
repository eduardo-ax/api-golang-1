package router

import (
	handler "github.com/eduardo-ax/api-golang-1/handle"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/user/:id", handler.ListUsersHandle)
		v1.GET("/user", handler.GetUserHandle)
		v1.POST("/user", handler.CreateUsersHandle)
		v1.DELETE("/user", handler.DeleteUserHandle)
		v1.PUT("/user", handler.UpdateUserHandle)
	}
}
