package models

import (
	"time"

	"github.com/matthewhartstonge/argon2"
	"gorm.io/gorm"
)

type User struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Fullname  string    `json:"fullname" gorm:"type:varchar(255); not null"`
	Email     string    `json:"email" gorm:"type:varchar(255); not null; unique"`
	Password  string    `json:"-" gorm:"type:varchar(255); not null"`
	IsAdmin   bool      `json:"is_admin" gorm:"type:bool; default:false"`
	Blogs     []Blog    `json:"blogs" gorm:"foreignKey:UserID"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp; not null; default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp; not null; default:CURRENT_TIMESTAMP"`
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
