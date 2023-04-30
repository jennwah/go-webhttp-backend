package task

import (
	"context"
	"github.com/jennwah/go-webhttp-backend/domain/task"
	"github.com/pkg/errors"

	"github.com/jmoiron/sqlx"
)

type taskRepository struct {
	database *sqlx.DB
}

func NewTaskRepository(db *sqlx.DB) task.TaskRepository {
	return &taskRepository{
		database: db,
	}
}

func (tr *taskRepository) Create(c context.Context, task *task.Task) error {
	query := `INSERT INTO tasks (title) VALUE (?);`
	_, err := tr.database.Exec(query, task.Title)
	if err != nil {
		return errors.Wrapf(err, "create new task")
	}
	return nil
}

func (tr *taskRepository) GetByID(c context.Context, id int64) (task.Task, error) {
	task := task.Task{}
	query := `SELECT t.id, t.title, t.created, t.updated FROM tasks t where t.id = ?;`
	err := tr.database.Get(&task, query, id)

	return task, err
}
