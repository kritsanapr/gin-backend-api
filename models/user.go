package models

import (
	"time"

	"github.com/matthewhartstonge/argon2"
	"gorm.io/gorm"
)

type User struct {
	ID        uint   `gorm:"primary_key" json:"id"`
	Fullname  string `gorm:"type:varchar(255); not null"`
	Email     string `gorm:"type:varchar(255); not null; unique"`
	Password  string `gorm:"type:varchar(255); not null"`
	IsAdmin   bool   `gorm:"type:bool; default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (user *User) BeforeCreate(db *gorm.DB) error {
	user.Password = hashPassword(user.Password)
	return nil
}

func hashPassword(password string) string {
	argon := argon2.DefaultConfig()
	encoded, _ := argon.HashEncoded([]byte(password))
	return string(encoded)
}
