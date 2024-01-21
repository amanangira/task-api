package repository

import (
	"context"

	"task/persistence/models"

	"github.com/jmoiron/sqlx"
)

type ITaskRepository interface {
	Create(ctx context.Context, user *models.Task) error
	GetTask(taskID string) (models.Task, error)
	UpdateTask(taskID string, task models.Task) error
	DeleteTask(taskID string) error
}

type TaskRepository struct {
	dbClient *sqlx.DB
}

func NewTaskRepository(dbClient *sqlx.DB) ITaskRepository {
	return TaskRepository{
		dbClient,
	}
}

// Create
// TODO - document function
func (u TaskRepository) Create(ctx context.Context, task *models.Task) error {
	row := u.dbClient.QueryRowxContext(
		ctx,
		`INSERT INTO 
    				tasks (title, description, priority, due_at)
    				values ($1, $2, $3, $4)
				returning id;`,
		task.Title,
		task.Description,
		task.Priority,
		task.DueAt,
	)

	if row.Err() != nil {
		return row.Err()
	}

	if err := row.Scan(&task.ID); err != nil {
		return err
	}

	return nil
}

func (u TaskRepository) GetTask(taskID string) (models.Task, error) {
	var task models.Task
	err := u.dbClient.
		QueryRow("SELECT id, title, description, priority, created_at, updated_at, due_at FROM tasks WHERE id = $1", taskID).
		Scan(&task.ID,
			&task.Title,
			&task.Description,
			&task.Priority,
			&task.CreatedAt,
			&task.UpdatedAt,
			&task.DueAt)
	if err != nil {
		return task, err
	}

	return task, nil
}

func (u TaskRepository) UpdateTask(taskID string, task models.Task) error {
	_, err := u.dbClient.
		Exec("UPDATE tasks SET title = $1, description = $2, priority = $3, updated_at = now(), due_at = $6 WHERE id = $7;",
			task.Title,
			task.Description,
			task.Priority,
			task.DueAt,
			taskID)

	return err
}

func (u TaskRepository) DeleteTask(taskID string) error {
	_, err := u.dbClient.Exec("DELETE FROM tasks WHERE id = $1;", taskID)

	return err
}
