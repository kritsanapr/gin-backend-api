package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kritsanapr/gin-backend-api/configs"
	"github.com/kritsanapr/gin-backend-api/routes"
)

func main() {
	configs.ConnectionDB()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := SetupRouter()

	router.Run(":" + os.Getenv("PORT"))
}

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"API VERSION": "1.0.0",
		})
	})

	apiV1 := router.Group("/api/v1")
	routes.InitHomeRoutes(apiV1)
	routes.InitProductRoutes(apiV1)
	routes.InitUserRoutes(apiV1)

	return router
}
