package entities

import (
	"fp-2/helpers"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"not null;uniqueIndex"`
	Email    string `gorm:"not null;uniqueIndex"`
	Password string `gorm:"not null"`
	Age      int    `gorm:"not null"`
}

func (user *User) BeforeCreate(tx *gorm.DB) error {
	user.Password = helpers.HashPass(user.Password)

	return nil
}
