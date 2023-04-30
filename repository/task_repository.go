package repository

import (
	"context"
	"github.com/pkg/errors"

	"github.com/jmoiron/sqlx"

	"github.com/jennwah/go-webhttp-backend/domain"
)

type taskRepository struct {
	database *sqlx.DB
}

func NewTaskRepository(db *sqlx.DB) domain.TaskRepository {
	return &taskRepository{
		database: db,
	}
}

func (tr *taskRepository) Create(c context.Context, task *domain.Task) error {
	query := `INSERT INTO tasks (title) VALUE (?);`
	_, err := tr.database.Exec(query, task.Title)
	if err != nil {
		return errors.Wrapf(err, "create new task")
	}
	return nil
}

func (tr *taskRepository) GetByID(c context.Context, id int64) (domain.Task, error) {
	task := domain.Task{}
	query := `SELECT t.id, t.title, t.created, t.updated FROM tasks t where t.id = ?;`
	err := tr.database.Get(&task, query, id)

	return task, err
}
