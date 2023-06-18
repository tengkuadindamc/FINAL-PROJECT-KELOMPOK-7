package dto

import (
	"time"
)

type UserRequest struct {
	Fullname string `json:"full_name" binding:"required"`
	Email    string `json:"email" gorm:"unique" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Role     string `json:"role"`
}

type LoginRequest struct {
	Email    string `json:"email" gorm:"unique" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type UserResponse struct {
	Id        int       `json:"id" gorm:"primaryKey;UNIQUE"`
	Fullname  string    `json:"full_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime:true"`
}

type UserResponses struct {
	Id        int       `json:"id" gorm:"primaryKey;UNIQUE"`
	Fullname  string    `json:"full_name"`
	Email     string    `json:"email"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime:true"`
}