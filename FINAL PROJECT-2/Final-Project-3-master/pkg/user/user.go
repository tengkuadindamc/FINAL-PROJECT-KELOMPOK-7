package user

import (
	"final-project3/pkg/user/controller"
	"final-project3/pkg/user/repository"
	"final-project3/pkg/user/usecase"

	"gorm.io/gorm"
)

func InitHttpUserController(db *gorm.DB) *controller.UserHTTPController {
	repo := repository.InitRepositoryUser(db)
	uc := usecase.InitUsecaseUser(repo)
	controller := controller.InitControllerUser(uc)

	return controller
}
