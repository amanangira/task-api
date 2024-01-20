package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"task/persistence/models"
)

type ITaskRepository interface {
	Create(ctx context.Context, user *models.Task) error
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
