package task

import (
	"context"
	"time"
)

type Task struct {
	ID      int64     `json:"-" db:"id" redis:"id"`
	Title   string    `validate:"required,gt=2" json:"title" db:"title" redis:"title"`
	Created time.Time `json:"created" db:"created" redis:"created"`
	Updated time.Time `json:"updated" db:"updated" redis:"updated"`
}

type TaskRepository interface {
	Create(c context.Context, task *Task) error
	GetByID(c context.Context, id int64) (Task, error)
}

type TaskUsecase interface {
	Create(c context.Context, task *Task) error
	GetByID(c context.Context, id int64) (Task, error)
}
