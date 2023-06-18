package category

import (
	"final-project3/pkg/category/controller"
	"final-project3/pkg/category/repository"
	"final-project3/pkg/category/usecase"

	"gorm.io/gorm"
)

func InitHttpCategoryController(db *gorm.DB) *controller.CategoryHTTPController {
	repo := repository.InitRepositoryCategory(db)
	uc := usecase.InitUsecaseCategory(repo)
	controller := controller.InitControllerCategory(uc)

	return controller
}
