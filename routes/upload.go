package routes

import (
	"github.com/gin-gonic/gin"
	uploadcontroller "github.com/kritsanapr/gin-backend-api/controllers/upload"
)

func InitUploadRoutes(rg *gin.RouterGroup) {

	routerGroup := rg.Group("/upload")

	//{{domain_url}}/api/v1/upload
	routerGroup.POST("/", uploadcontroller.UploadFile)

}
