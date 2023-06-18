package entities

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	UserId  uint `gorm:"not null"`
	User    User
	PhotoId uint `gorm:"not null"`
	Photo   Photo
	Message string `gorm:"not null"`
}
