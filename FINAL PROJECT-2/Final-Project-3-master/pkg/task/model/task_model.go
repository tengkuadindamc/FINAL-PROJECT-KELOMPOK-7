package model

import (
	"final-project3/pkg/user/model"
	"time"
)

type Task struct {
	Id          int64      `json:"id" gorm:"primaryKey;UNIQUE"`
	Title       string     `json:"title" binding:"required"`
	Description string     `json:"description" binding:"required"`
	Status      bool       `json:"status" binding:"required"`
	UserId      int        `json:"user_id"`
	CategoryId  int        `json:"category_id"`
	CreatedAt   time.Time  `json:"created_at" gorm:"autoCreateTime:true"`
	UpdatedAt   time.Time  `json:"updated_at" gorm:"autoUpdateTime:true"`
	User        model.User `json:"User"  gorm:"foreignKey:Id; constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
