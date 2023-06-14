package usercontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kritsanapr/gin-backend-api/configs"
	"github.com/kritsanapr/gin-backend-api/models"
)

func GetAll(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "user api v1",
	})
}

func GetById(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"message": "get user by id",
		"data": map[string]interface{}{
			"id": id,
		},
	})
}

func Register(c *gin.Context) {
	var json InputRegister
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Fullname: json.Fullname,
		Email:    json.Email,
		Password: json.Password,
	}

	result := configs.DB.Debug().Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error",
			"data":    result.Error,
		})
		return
	}

	c.JSON(200, gin.H{
		"data": "สมัครสมาชิกเรียบร้อยแล้ว",
	})
}

func Login(c *gin.Context) {
	var json InputLogin
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(200, gin.H{
		"message": json,
	})
}

func SearchByName(c *gin.Context) {
	fullname := c.Query("fullname")

	c.JSON(200, gin.H{
		"message": "search by name",
		"data": map[string]interface{}{
			"name": fullname,
		},
	})
}
