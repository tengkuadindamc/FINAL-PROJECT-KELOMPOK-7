package faker

import (
	"finalproject4/model"
	"time"

	"gorm.io/gorm"
)

func Admin(db *gorm.DB) *model.User {
	return &model.User{
		GormModel: model.GormModel{
			ID:        0,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		},
		Fullname: "Admin",
		Email:    "admin@admin.com",
		Password: "password123", // password123
		Role:     "admin",
	}
}
