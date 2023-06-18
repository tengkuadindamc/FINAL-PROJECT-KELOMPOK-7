package repository

import (
	"final-project3/pkg/task/model"

	"gorm.io/gorm"
)

type RepositoryInterfaceTask interface {
	CreateNewTask(task model.Task) (model.Task, error)
	GetAllTask() ([]model.Task, error)
	GetTaskById(taskId int) (model.Task, error)
	UpdateTaskById(taskId int, task model.Task) (model.Task, error)
	DeleteTaskById(taskId int) error
}

type repositoryTask struct {
	db *gorm.DB
}

func InitRepositoryTask(db *gorm.DB) RepositoryInterfaceTask {
	db.AutoMigrate(&model.Task{})
	return &repositoryTask{
		db: db,
	}
}

// CreateNewTask implements RepositoryInterfaceTask
func (r *repositoryTask) CreateNewTask(task model.Task) (model.Task, error) {
	if err := r.db.Table("tasks").Create(&task).Error; err != nil {
		return task, err
	}

	return task, nil
}

// GetAllTask implements RepositoryInterfaceTask
func (r *repositoryTask) GetAllTask() ([]model.Task, error) {
	var tasks []model.Task
	if err := r.db.Preload("User").Table("tasks").Find(&tasks).Error; err != nil {
		return tasks, err
	}

	return tasks, nil
}

// GetTaskById implements RepositoryInterfaceTask
func (r *repositoryTask) GetTaskById(taskId int) (model.Task, error) {
	var task model.Task
	if err := r.db.Table("tasks").Where("id = ?", taskId).First(&task).Error; err != nil {
		return task, err
	}

	return task, nil
}

// UpdateTaskById implements RepositoryInterfaceTask
func (r *repositoryTask) UpdateTaskById(taskId int, task model.Task) (model.Task, error) {
	if err := r.db.Table("tasks").Where("id = ?", taskId).Updates(&task).Error; err != nil {
		return task, err
	}

	return task, nil
}

// DeleteTaskById implements RepositoryInterfaceTask
func (r *repositoryTask) DeleteTaskById(taskId int) error {
	if err := r.db.Table("tasks").Where("id = ?", taskId).Delete(&model.Task{}).Error; err != nil {
		return err
	}

	return nil
}
