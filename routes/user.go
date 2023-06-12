package routes

import (
	"github.com/gin-gonic/gin"
	usercontroller "github.com/kritsanapr/gin-backend-api/controllers/user"
)

func InitUserRoutes(rg *gin.RouterGroup) {
	routerGroup := rg.Group("/user")

	routerGroup.GET("/", usercontroller.GetAll)
	routerGroup.GET("/:id", usercontroller.GetById)
	routerGroup.POST("/register", usercontroller.Register)

}
