package service

import (
	"errors"
	"finalproject4/model"
	"finalproject4/repository"
)

type CategoriesService interface {
	CreateCategory(role_user string, categories model.CategoryInput) (model.Category, error)
	GetAllCategory() ([]model.Category, error)
	PatchCategory(role_user string, id_category int, categories model.CategoryInput) (model.Category, error)
	DeleteCategory(role_user string, id_category int) error
}

type categoriesService struct {
	CategoriesRepository repository.CategoriesRepository
}

func NewCategoriesService(categoriesRepository repository.CategoriesRepository) *categoriesService {
	return &categoriesService{categoriesRepository}
}

func (s *categoriesService) CreateCategory(role_user string, categories model.CategoryInput) (model.Category, error) {
	if role_user != "admin" {
		return model.Category{}, errors.New("you are not allowed to create a category")
	}
	category := model.Category{
		Type:                categories.Type,
		Sold_product_amount: 0,
	}

	categoriesData, err := s.CategoriesRepository.CreateCategory(category)
	if err != nil {
		return model.Category{}, err
	}

	return categoriesData, nil
}

func (s *categoriesService) GetAllCategory() ([]model.Category, error) {
	return s.CategoriesRepository.GetAllCategory()
}

func (s *categoriesService) PatchCategory(role_user string, id_category int, categories model.CategoryInput) (model.Category, error) {
	if role_user != "admin" {
		return model.Category{}, errors.New("you are not allowed to patch a category")
	}
	category := model.Category{
		Type: categories.Type,
	}

	_, err := s.CategoriesRepository.UpdateCategory(id_category, category)
	if err != nil {
		return model.Category{}, err
	}

	categoryData, err := s.CategoriesRepository.GetCategoryByID(id_category)
	if categoryData.ID == 0 {
		return model.Category{}, errors.New("category not found")
	}

	return categoryData, err
}

func (s *categoriesService) DeleteCategory(role_user string, id_category int) error {
	if role_user != "admin" {
		return errors.New("you are not allowed to delete a category")
	}
	categoryData, err := s.CategoriesRepository.GetCategoryByID(id_category)
	if err != nil {
		return errors.New("category not found")
	}
	if categoryData.ID == 0 {
		return errors.New("category not found")
	}

	err = s.CategoriesRepository.DeleteCategory(categoryData)
	if err != nil {
		return err
	}

	return nil
}
