package task

import (
	"final-project3/pkg/task/controller"
	"final-project3/pkg/task/repository"
	"final-project3/pkg/task/usecase"

	"gorm.io/gorm"
)

func InitHttpTaskController(db *gorm.DB) *controller.TaskHTTPController {
	repo := repository.InitRepositoryTask(db)
	uc := usecase.InitUsecaseTask(repo)
	controller := controller.InitControllerTask(uc)

	return controller
}
