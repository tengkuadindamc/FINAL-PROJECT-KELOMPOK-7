package entities

import (
	"gorm.io/gorm"
)

type SocialMedia struct {
	gorm.Model
	Name           string `gorm:"not null"`
	SocialMediaUrl string `gorm:"not null"`
	UserId         uint   `gorm:"not null"`
	User           User
}
