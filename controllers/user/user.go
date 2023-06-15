package usercontroller

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/kritsanapr/gin-backend-api/configs"
	"github.com/kritsanapr/gin-backend-api/models"
	"github.com/kritsanapr/gin-backend-api/utils"
	"github.com/matthewhartstonge/argon2"
)

func GetAll(c *gin.Context) {
	var users []models.User
	// configs.DB.Find(&users)
	configs.DB.Preload("Blogs").First(&users)

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
	var input InputLogin
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Email:    input.Email,
		Password: input.Password,
	}

	userAccount := configs.DB.Where("email = ?", input.Email).First(&user)
	if userAccount.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "ไม่พบข้อมูลผู้ใช้งาน",
		})
		return
	}

	fmt.Println("Email : ", user.Email)
	fmt.Println("Password : ", user.Password)
	ok, err := argon2.VerifyEncoded([]byte(input.Password), []byte(user.Password))
	if err != nil {
		fmt.Println("Error verifying password")
		fmt.Println(err)
	}

	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "รหัสผ่านไม่ถูกต้อง",
		})
		return
	}

	// Create token from jwt
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(time.Hour).Unix(),
	})

	secert := os.Getenv("JWT_SECRET")
	accessToken, _ := token.SignedString([]byte(secert))

	fmt.Println("Token : ", accessToken)

	c.JSON(200, gin.H{
		"message": "เข้าสู่ระบบสำเร็จ",
		"token":   accessToken,
		"data": []map[string]interface{}{
			{
				"fullname": user.Fullname,
				"email":    user.Email,
			},
		},
	})

}

func SearchByName(c *gin.Context) {
	fullname := c.Query("fullname")

	users := []models.User{}
	result := configs.DB.Where("fullname LIKE ?", "%"+fullname+"%").Scopes(utils.Paginate(c)).Find(&users)
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

func GetProfile(c *gin.Context) {
	user := c.MustGet("user").(models.User)

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}
