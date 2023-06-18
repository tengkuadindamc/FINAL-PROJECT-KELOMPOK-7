package usecase

import (
	"errors"
	"final-project3/pkg/category/dto"
	"final-project3/pkg/category/model"
	"final-project3/pkg/category/repository"
)

type UsecaseInterfaceCategory interface {
	CreateNewCategory(req dto.CategoryRequest) (model.Category, error)
	GetAllCategory() ([]model.Category, error)
	UpdateCategoryById(categoryId int, input dto.CategoryRequest) (model.Category, error)
	DeleteCategoryById(categoryId int) error
}

type usecaseCategory struct {
	repository repository.RepositoryInterfaceCategory
}

func InitUsecaseCategory(repository repository.RepositoryInterfaceCategory) UsecaseInterfaceCategory {
	return &usecaseCategory{
		repository: repository,
	}
}

// CreateNewCategory implements UsecaseInterfaceCategory
func (u *usecaseCategory) CreateNewCategory(req dto.CategoryRequest) (model.Category, error) {
	var category model.Category
	isCategoryExist, _ := u.repository.GetCategoryById(int(category.Id))
	if isCategoryExist.Id != 0 {
		return category, errors.New("categories already exist")
	}

	payload := model.Category{
		Type: req.Type,
	}
	newCategory, err := u.repository.CreateNewCategory(payload)
	if err != nil {
		return newCategory, err
	}

	return newCategory, nil
}

// GetAllCategory implements UsecaseInterfaceCategory
func (u *usecaseCategory) GetAllCategory() ([]model.Category, error) {
	categories, err := u.repository.GetAllCategory()
	if err != nil {
		return categories, err
	}

	return categories, nil
}

// UpdateCategoryById implements UsecaseInterfaceCategory
func (u *usecaseCategory) UpdateCategoryById(categoryId int, input dto.CategoryRequest) (model.Category, error) {
	payload := model.Category{
		Type: input.Type,
	}
	category, err := u.repository.UpdateCategoryById(categoryId, payload)
	if err != nil {
		return category, err
	}

	return category, nil
}

// DeleteCategoryById implements UsecaseInterfaceCategory
func (u *usecaseCategory) DeleteCategoryById(categoryId int) error {
	err := u.repository.DeleteCategoryById(categoryId)
	if err != nil {
		return err
	}

	return nil
}
