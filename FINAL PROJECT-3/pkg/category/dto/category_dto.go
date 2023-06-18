package dto

import (
	"final-project3/pkg/task/model"
	"time"
)

type CategoryRequest struct {
	Type string `json:"type" binding:"required"`
}

type CategoryResponse struct {
	Id        int64        `json:"id"`
	Type      string       `json:"type"`
	Tasks     []model.Task `json:"tasks"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
}
