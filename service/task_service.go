package service

import (
	"context"

	"task/persistence/models"
	"task/repository"
)

type ITaskService interface {
	Create(ctx context.Context, input models.Task) (string, error)
	GetTask(taskID string) (models.Task, error)
	UpdateTask(taskID string, task models.Task) error
	DeleteTask(taskID string) error
}

type TaskService struct {
	taskRepository repository.ITaskRepository
}

// TODO - change implementation to return pointer to service and make receivers pointer based for lite container

func NewTaskService(userRepository repository.ITaskRepository) ITaskService {
	return TaskService{
		taskRepository: userRepository,
	}
}

func (u TaskService) Create(ctx context.Context, input models.Task) (string, error) {
	err := u.taskRepository.Create(ctx, &input)

	return input.ID, err
}

func (s TaskService) GetTask(taskID string) (models.Task, error) {
	return s.taskRepository.GetTask(taskID)
}

func (s TaskService) UpdateTask(taskID string, task models.Task) error {
	return s.taskRepository.UpdateTask(taskID, task)
}

func (s TaskService) DeleteTask(taskID string) error {
	return s.taskRepository.DeleteTask(taskID)
}
