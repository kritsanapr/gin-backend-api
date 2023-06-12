package routes

import "github.com/gin-gonic/gin"

func InitHomeRoutes(rg *gin.RouterGroup) {
	routerGroup := rg.Group("/home")
	routerGroup.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "home api v1",
		})
	})
}
