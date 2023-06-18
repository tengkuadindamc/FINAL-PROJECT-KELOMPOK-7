package repository

import (
	"final-project3/pkg/category/model"

	"gorm.io/gorm"
)

type RepositoryInterfaceCategory interface {
	CreateNewCategory(category model.Category) (model.Category, error)
	GetAllCategory() ([]model.Category, error)
	GetCategoryById(CategoryId int) (model.Category, error)
	UpdateCategoryById(categoryId int, category model.Category) (model.Category, error)
	DeleteCategoryById(categoryId int) error
}

type repositoryCategory struct {
	db *gorm.DB
}

func InitRepositoryCategory(db *gorm.DB) RepositoryInterfaceCategory {
	db.AutoMigrate(model.Category{})
	return &repositoryCategory{
		db: db,
	}
}

func (r *repositoryCategory) CreateNewCategory(category model.Category) (model.Category, error) {
	if err := r.db.Table("categories").Create(&category).Error; err != nil {
		return category, err
	}

	return category, nil
}

// GetAllCategory implements RepositoryInterfaceCategory
func (r *repositoryCategory) GetAllCategory() ([]model.Category, error) {
	var categories []model.Category
	if err := r.db.Preload("Tasks").Find(&categories).Error; err != nil {
		return categories, err
	}

	return categories, nil
}

// GetCategoryById implements RepositoryInterfaceCategory
func (r *repositoryCategory) GetCategoryById(categoryId int) (model.Category, error) {
	var category model.Category
	if err := r.db.Table("categories").Where("id = ?", categoryId).First(&category).Error; err != nil {
		return category, err
	}

	return category, nil
}

// UpdateCategoryById implements RepositoryInterfaceCategory
func (r *repositoryCategory) UpdateCategoryById(categoryId int, category model.Category) (model.Category, error) {
	if err := r.db.Table("categories").Where("id = ?", categoryId).Updates(&category).Error; err != nil {
		return category, err
	}

	return category, nil
}

// DeleteCategoryById implements RepositoryInterfaceCategory
func (r *repositoryCategory) DeleteCategoryById(CategoryId int) error {
	if err := r.db.Table("categories").Where("id = ?", CategoryId).Delete(&model.Category{}).Error; err != nil {
		return err
	}

	return nil
}