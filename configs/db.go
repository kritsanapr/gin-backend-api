package configs

import (
	"fmt"
	"os"

	"github.com/kritsanapr/gin-backend-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectionDB() {
	dsn := os.Getenv("DATABASE_DNS")
	dsn = os.ExpandEnv(dsn)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		fmt.Println("เชื่อมต่อฐานข้อมูลไม่สำเร็จ")
		fmt.Println(dsn)
		panic(err)
	}

	// Migration
	db.AutoMigrate(&models.User{}, &models.Blog{})

	fmt.Println("เชื่อมต่อฐานข้อมูลสำเร็จ")
	DB = db
}
