package dto

import (
	"time"
)

type TaskRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Status      bool   `json:"status"`
	UserId      int    `json:"user_id"`
	CategoryId  int    `json:"category_id"`
}

type TaskResponse struct {
	Id          int64     `json:"id" gorm:"primaryKey;UNIQUE"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      bool      `json:"status"`
	UserId      int       `json:"user_id"`
	CategoryId  int       `json:"category_id"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime:true"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime:true"`
}
