package task

import (
	"context"
	"github.com/jennwah/go-webhttp-backend/domain/task"
	"time"
)

type taskUsecase struct {
	taskRepository task.TaskRepository
	contextTimeout time.Duration
}

func NewTaskUsecase(taskRepository task.TaskRepository, timeout time.Duration) task.TaskUsecase {
	return &taskUsecase{
		taskRepository: taskRepository,
		contextTimeout: timeout,
	}
}

func (tu *taskUsecase) Create(c context.Context, task *task.Task) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.taskRepository.Create(ctx, task)
}

func (tu *taskUsecase) GetByID(c context.Context, id int64) (task.Task, error) {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.taskRepository.GetByID(ctx, id)
}
