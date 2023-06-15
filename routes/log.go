package routes

import (
	"github.com/gin-gonic/gin"
	logcontroller "github.com/kritsanapr/gin-backend-api/controllers/log"
)

func InitLogRoutes(rg *gin.RouterGroup) {

	routerGroup := rg.Group("/logs")

	//{{domain_url}}/api/v1/logs
	routerGroup.GET("/", logcontroller.GetLog)

	//{{domain_url}}/api/v1/logs/new
	routerGroup.POST("/new", logcontroller.InsertLog)

}
