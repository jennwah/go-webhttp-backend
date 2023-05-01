package task

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"

	"github.com/jennwah/go-webhttp-backend/domain/task"
)

var kvTaskKey = "task-%v"

const keyTTL = 24 * time.Hour

type taskRepository struct {
	database *sqlx.DB
	kv       *redis.Client
	logger   *log.Entry
}

func NewTaskRepository(db *sqlx.DB, kv *redis.Client, logger *log.Entry) task.TaskRepository {
	return &taskRepository{
		database: db,
		kv:       kv,
		logger:   logger,
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

// GetByID does a read aside cache strategy
// Reads from cache first, if found, returns
// Else request from primary datasource and updates cache before returning
func (tr *taskRepository) GetByID(c context.Context, id int64) (task.Task, error) {
	taskRes := task.Task{}
	kvKey := fmt.Sprintf(kvTaskKey, id)

	err := tr.kv.HGetAll(c, kvKey).Scan(&taskRes)
	if err != nil {
		tr.logger.Errorf("Redis HGetAll key: %v, err: %v", kvKey, err)
	}
	if taskRes.Title != "" {
		tr.logger.Info("Task returned from cache")
		return taskRes, nil
	}

	query := `SELECT t.id, t.title, t.created, t.updated FROM tasks t where t.id = ?;`
	err = tr.database.Get(&taskRes, query, id)

	if err == nil {
		tr.logger.Info("Setting task result into cache")
		tr.kv.HSet(c, kvKey, &taskRes)
		tr.kv.Expire(c, kvKey, keyTTL)
	}

	return taskRes, err
}
