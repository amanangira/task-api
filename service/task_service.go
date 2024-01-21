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
	ListTasks(ctx context.Context) ([]models.Task, error)
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

func (t TaskService) Create(ctx context.Context, input models.Task) (string, error) {
	err := t.taskRepository.Create(ctx, &input)

	return input.ID, err
}

func (t TaskService) GetTask(taskID string) (models.Task, error) {
	return t.taskRepository.GetTask(taskID)
}

func (t TaskService) UpdateTask(taskID string, task models.Task) error {
	return t.taskRepository.UpdateTask(taskID, task)
}

func (t TaskService) DeleteTask(taskID string) error {
	return t.taskRepository.DeleteTask(taskID)
}

func (t TaskService) ListTasks(ctx context.Context) ([]models.Task, error) {
	return t.taskRepository.ListTasks(ctx)
}
