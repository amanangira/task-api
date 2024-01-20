package service

import (
	"context"
	"task/persistence/models"
	"task/repository"
)

type ITaskService interface {
	Create(ctx context.Context, input models.Task) (string, error)
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
