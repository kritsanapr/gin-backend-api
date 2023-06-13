package configs

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectionDB() {
	dsn := os.Getenv("DATABASE_DNS")

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		fmt.Println("เชื่อมต่อฐานข้อมูลไม่สำเร็จ")
		panic(err)
	}

	fmt.Println("เชื่อมต่อฐานข้อมูลสำเร็จ")
	DB = db
}
