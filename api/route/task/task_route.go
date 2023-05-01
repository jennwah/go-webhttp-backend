package task

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"

	taskController "github.com/jennwah/go-webhttp-backend/api/controller/task"
	"github.com/jennwah/go-webhttp-backend/config"
	taskRepository "github.com/jennwah/go-webhttp-backend/repository/task"
	taskUseCase "github.com/jennwah/go-webhttp-backend/usecase/task"
)

func NewTaskRouter(env *config.Env, db *sqlx.DB, kv *redis.Client, group *gin.RouterGroup, logger *log.Entry) {
	ctxTimeout := time.Duration(env.ContextTimeout) * time.Second

	taskRepo := taskRepository.NewTaskRepository(db, kv, logger)
	taskUsecase := taskUseCase.NewTaskUsecase(taskRepo, ctxTimeout)
	taskCtrl := taskController.New(taskUsecase)

	group.GET("/task/:taskID", taskCtrl.GetByID)
	group.POST("/task", taskCtrl.Create)
}
