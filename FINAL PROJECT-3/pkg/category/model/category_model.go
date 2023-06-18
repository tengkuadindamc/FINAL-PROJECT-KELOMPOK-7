package model

import (
	"final-project3/pkg/task/model"
	"time"
)

type Category struct {
	Id        int64        `json:"id" gorm:"primaryKey;UNIQUE"`
	Type      string       `json:"type" binding:"required"`
	Tasks     []model.Task `json:"tasks" gorm:"foreignKey:CategoryId; constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt time.Time    `json:"created_at" gorm:"autoCreateTime:true"`
	UpdatedAt time.Time    `json:"updated_at" gorm:"autoUpdateTime:true"`
}
