package usercontroller

import "github.com/gin-gonic/gin"

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
	// Get value from form-date
	// userData := map[string]interface{}{
	// 	"fullname": c.PostForm("fullname"),
	// 	"email":    c.PostForm("email"),
	// 	"password": c.PostForm("password"),
	// }

	// Get value from json body
	var userData map[string]interface{}
	c.BindJSON(&userData)

	c.JSON(200, gin.H{
		"data": userData,
	})
}

func Login(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "login",
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
