package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Product struct {
	GormModel
	Title      string `gorm:"not null" json:"title" form:"title" valid:"required~Product name is required"`
	Stock      int    `gorm:"not null" json:"stock" form:"stock" valid:"required~Stock is required, range(5|999999)~Minimum product stock is 5"`
	Price      int    `gorm:"not null" json:"price" form:"price" valid:"required~Product Price is required, range(0|50000000)~Product Price is out of range"`
	CategoryID int    `json:"category_id" form:"category_id" gorm:"foreignkey:category_id;embedded"`
}

type AddProduct struct {
	Title      string `gorm:"not null" json:"title" form:"title" valid:"required~Product name is required"`
	Stock      int    `gorm:"not null" json:"stock" form:"stock" valid:"required~Stock is required, range(5|999999)~Minimum product stock is 5"`
	Price      int    `gorm:"not null" json:"price" form:"price" valid:"required~Product Price is required, range(0|50000000)~Product Price is out of range"`
	CategoryID int    `json:"category_id" form:"category_id"`
}

type EditProduct struct {
	Title      string `gorm:"not null" json:"title" form:"title" valid:"required~Product name is required"`
	Stock      int    `gorm:"not null" json:"stock" form:"stock" valid:"required~Stock is required, range(5|999999)~Minimum product stock is 5"`
	Price      int    `gorm:"not null" json:"price" form:"price" valid:"required~Product Price is required, range(0|50000000)~Product Price is out of range"`
	CategoryID int    `json:"category_id" form:"category_id"`
}

type ResponseAddProduct struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Stock      int       `json:"stock"`
	Price      int       `json:"price"`
	CategoryID int       `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
}

type ResponseEditProduct struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Stock      int       `json:"stock"`
	Price      int       `json:"price"`
	CategoryID int       `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)
	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (p *Product) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errUpdate := govalidator.ValidateStruct(p)
	if errUpdate != nil {
		err = errUpdate
		return
	}

	err = nil
	return
}
