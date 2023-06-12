package routes

import "github.com/gin-gonic/gin"

func InitUserRoutes(rg *gin.RouterGroup) {
	routerGroup := rg.Group("/user")
	routerGroup.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "user api v1",
		})
	})
}
