package main

import (
	"log"
	"os"

	"time"

	"github.com/gin-contrib/cors"
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
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://foo.com"},
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))
	router.Run(":" + os.Getenv("PORT"))
}

func SetupRouter() *gin.Engine {
	//  Load environment variables
	godotenv.Load(".env")

	// Set the router as the default one provided by Gin
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"API VERSION": "1.0.0",
		})
	})

	// Setup route group for the API
	apiV1 := router.Group("/api/v1")
	routes.InitHomeRoutes(apiV1)
	routes.InitProductRoutes(apiV1)
	routes.InitUserRoutes(apiV1)

	return router
}
