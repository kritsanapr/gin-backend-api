package routes

import (
	"github.com/gin-gonic/gin"
	usercontroller "github.com/kritsanapr/gin-backend-api/controllers/user"
	"github.com/kritsanapr/gin-backend-api/middlewares"
)

func InitUserRoutes(rg *gin.RouterGroup) {
	routerGroup := rg.Group("/user")
	// {{url}}/api/v1/user
	routerGroup.GET("/", usercontroller.GetAll)

	// {{url}}/api/v1/user/:id
	routerGroup.GET("/:id", usercontroller.GetById)

	// {{url}}/api/v1/user/register
	routerGroup.POST("/register", usercontroller.Register)

	// {{url}}/api/v1/user/login
	routerGroup.POST("/login", usercontroller.Login)

	// {{url}}/api/v1/user/search?fullname={{fullname}}
	routerGroup.GET("/search", usercontroller.SearchByName)

	// {{url}}/api/v1/user/get-profile
	routerGroup.GET("/get-profile", middlewares.AuthJWT(), usercontroller.GetProfile)
}
