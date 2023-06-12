package routes

import "github.com/gin-gonic/gin"

func InitProductRoutes(rg *gin.RouterGroup) {
	routerGroup := rg.Group("/product")
	routerGroup.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "product api v1",
		})
	})
}
