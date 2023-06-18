package repository

import (
	"finalproject4/model"

	"gorm.io/gorm"
)

type ProductRepository interface {
	AddProduct(product model.Product) (model.Product, error)
	EditProduct(product model.Product) (model.Product, error)
	DeleteProduct(product model.Product) error
	ViewProduct() ([]model.Product, error)
	FindById(id int) (model.Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *productRepository {
	return &productRepository{
		db: db,
	}
}

func (p *productRepository) AddProduct(product model.Product) (model.Product, error) {
	err := p.db.Create(&product).Error
	return product, err
}

func (p *productRepository) EditProduct(product model.Product) (model.Product, error) {
	err := p.db.Save(&product).Error
	return product, err
}

func (p *productRepository) DeleteProduct(product model.Product) error {
	err := p.db.Debug().Delete(&product).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *productRepository) ViewProduct() ([]model.Product, error) {
	var products []model.Product
	err := p.db.Debug().Find(&products).Error
	if err != nil {
		return products, err
	}
	return products, nil
}

func (p *productRepository) FindById(id int) (model.Product, error) {
	var product model.Product
	err := p.db.Where("id = ?", id).First(&product).Error
	if err != nil {
		return product, err
	}
	return product, nil
}
