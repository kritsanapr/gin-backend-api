package usercontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kritsanapr/gin-backend-api/configs"
	"github.com/kritsanapr/gin-backend-api/models"
)

func GetAll(c *gin.Context) {
	var users []models.User
	configs.DB.Find(&users)

	c.JSON(200, gin.H{
		"data": users,
	})
}

func GetById(c *gin.Context) {
	id := c.Param("id")

	var user models.User
	result := configs.DB.First(&user, id)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "ไม่พบข้อมูล",
		})
		return
	}

	c.JSON(200, gin.H{
		"data": user,
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

	userExists := configs.DB.Where("email = ?", user.Email).First(&models.User{})
	if userExists.RowsAffected > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": "อีเมล์นี้มีผู้ใช้งานแล้ว",
		})
		return
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

	users := []models.User{}
	result := configs.DB.Where("fullname LIKE ?", "%"+fullname+"%").Find(&users)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "ไม่พบข้อมูล",
		})
		return
	}

	c.JSON(200, gin.H{
		"data": users,
	})
}
