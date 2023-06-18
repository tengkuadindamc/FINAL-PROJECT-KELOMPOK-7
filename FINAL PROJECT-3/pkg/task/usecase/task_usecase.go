package usecase

import (
	"errors"
	"final-project3/pkg/task/dto"
	"final-project3/pkg/task/model"
	"final-project3/pkg/task/repository"
)

type UsecaseInterfaceTask interface {
	CreateNewTask(req dto.TaskRequest) (model.Task, error)
	GetAllTask() ([]model.Task, error)
	UpdateTaskById(taskId int, input dto.TaskRequest) (model.Task, error)
	UpdateStatusByTaskId(taskId int, input dto.TaskRequest) (model.Task, error)
	UpdateCategoryByTaskId(taskId int, input dto.TaskRequest) (model.Task, error)
	DeleteTaskById(taskId int) error
}

type usecaseTask struct {
	repository repository.RepositoryInterfaceTask
}

func InitUsecaseTask(repository repository.RepositoryInterfaceTask) UsecaseInterfaceTask {
	return &usecaseTask{
		repository: repository,
	}
}

// CreateNewTask implements UsecaseInterfaceTask
func (u *usecaseTask) CreateNewTask(req dto.TaskRequest) (model.Task, error) {
	var task model.Task
	isTaskExist, _ := u.repository.GetTaskById(int(task.Id))
	if isTaskExist.Id != 0 {
		return task, errors.New("tasks already exist")
	}

	payload := model.Task{
		Title:       req.Title,
		Description: req.Description,
		Status:      false,
		UserId:      req.UserId,
		CategoryId:  req.CategoryId,
	}
	newTask, err := u.repository.CreateNewTask(payload)
	if err != nil {
		return newTask, err
	}

	return newTask, nil
}

// GetAllTask implements UsecaseInterfaceTask
func (u *usecaseTask) GetAllTask() ([]model.Task, error) {
	tasks, err := u.repository.GetAllTask()
	if err != nil {
		return tasks, err
	}

	return tasks, nil
}

// UpdateTaskById implements UsecaseInterfaceTask
func (u *usecaseTask) UpdateTaskById(taskId int, input dto.TaskRequest) (model.Task, error) {
	payload := model.Task{
		Title:       input.Title,
		Description: input.Description,
	}
	Task, err := u.repository.UpdateTaskById(taskId, payload)
	if err != nil {
		return Task, err
	}

	return Task, nil
}

// UpdateStatusByTaskId implements UsecaseInterfaceTask
func (u *usecaseTask) UpdateStatusByTaskId(taskId int, input dto.TaskRequest) (model.Task, error) {
	payload := model.Task{
		Status: input.Status,
	}
	task, err := u.repository.UpdateTaskById(taskId, payload)
	if err != nil {
		return task, err
	}

	return task, nil
}

// UpdateCategoryByTaskId implements UsecaseInterfaceTask
func (u *usecaseTask) UpdateCategoryByTaskId(taskId int, input dto.TaskRequest) (model.Task, error) {
	payload := model.Task{
		CategoryId: input.CategoryId,
	}
	Task, err := u.repository.UpdateTaskById(taskId, payload)
	if err != nil {
		return Task, err
	}

	return Task, nil
}

// DeleteTaskById implements UsecaseInterfaceTask
func (u *usecaseTask) DeleteTaskById(taskId int) error {
	err := u.repository.DeleteTaskById(taskId)
	if err != nil {
		return err
	}

	return nil
}
