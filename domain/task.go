package domain

import (
	"context"
	"time"
)

type Task struct {
	ID      int64     `json:"-" db:"id"`
	Title   string    `binding:"required" json:"title" db:"title"`
	Created time.Time `json:"created" db:"created"`
	Updated time.Time `json:"updated" db:"updated"`
}

type TaskRepository interface {
	Create(c context.Context, task *Task) error
	GetByID(c context.Context, id int64) (Task, error)
}

type TaskUsecase interface {
	Create(c context.Context, task *Task) error
	GetByID(c context.Context, id int64) (Task, error)
}
