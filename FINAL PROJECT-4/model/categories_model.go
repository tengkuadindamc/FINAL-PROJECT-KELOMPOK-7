package model

import "time"

type Category struct {
	GormModel
	Type                string    `json:"type" gorm:"not null"`
	Sold_product_amount int       `json:"sold_product_amount"`
	Products            []Product `json:"products" gorm:"constraint:OnDelete:SET NULL;"`
}

type CategoryInput struct {
	Type string `json:"type" gorm:"not null" binding:"required"`
}

type CategoryPostResponse struct {
	ID                  uint      `json:"id"`
	Type                string    `json:"type"`
	Sold_product_amount int       `json:"sold_product_amount"`
	CreatedAt           time.Time `json:"created_at"`
}

type CategoryGetResponse struct {
	Id                  uint              `json:"id"`
	Type                string            `json:"type"`
	Sold_product_amount int               `json:"sold_product_amount"`
	CreatedAt           time.Time         `json:"created_at"`
	UpdatedAt           time.Time         `json:"updated_at"`
	Products            []CategoryProduct `json:"products"`
}

type CategoryProduct struct {
	GormModel
	Title string `gorm:"not null" json:"title"`
	Price int    `gorm:"not null" json:"price"`
	Stock int    `gorm:"not null" json:"stock"`
}

type CategoryPatchResponse struct {
	ID                  uint      `json:"id"`
	Type                string    `json:"type"`
	Sold_product_amount int       `json:"sold_product_amount"`
	UpdatedAt           time.Time `json:"updated_at"`
}
