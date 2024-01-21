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
	ListTasks(ctx context.Context) ([]models.Task, error)
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
func (t TaskRepository) Create(ctx context.Context, task *models.Task) error {
	row := t.dbClient.QueryRowxContext(
		ctx,
		`INSERT INTO 
    				tasks (title, description, priority, due_at, updated_at)
    				values ($1, $2, $3, $4, now())
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

func (t TaskRepository) GetTask(taskID string) (models.Task, error) {
	var task models.Task
	err := t.dbClient.
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

func (t TaskRepository) UpdateTask(taskID string, task models.Task) error {
	_, err := t.dbClient.
		Exec("UPDATE tasks SET title = $1, description = $2, priority = $3, updated_at = now(), due_at = $4 WHERE id = $5;",
			task.Title,
			task.Description,
			task.Priority,
			task.DueAt,
			taskID)

	return err
}

func (t TaskRepository) DeleteTask(taskID string) error {
	_, err := t.dbClient.Exec("DELETE FROM tasks WHERE id = $1;", taskID)

	return err
}

func (t TaskRepository) ListTasks(ctx context.Context) ([]models.Task, error) {
	var tasks []models.Task
	var tempTask models.Task
	rows, err := t.dbClient.
		QueryxContext(
			ctx,
			"SELECT id, title, description, priority, created_at, updated_at, due_at FROM tasks;")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		if scanErr := rows.StructScan(&tempTask); scanErr != nil {
			return nil, scanErr
		}

		tasks = append(tasks, tempTask)
	}

	return tasks, nil
}
