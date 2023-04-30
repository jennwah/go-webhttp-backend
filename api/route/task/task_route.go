package task

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"

	taskController "github.com/jennwah/go-webhttp-backend/api/controller/task"
	"github.com/jennwah/go-webhttp-backend/config"
	taskRepository "github.com/jennwah/go-webhttp-backend/repository/task"
	taskUseCase "github.com/jennwah/go-webhttp-backend/usecase/task"
)

func NewTaskRouter(env *config.Env, db *sqlx.DB, group *gin.RouterGroup) {
	ctxTimeout := time.Duration(env.ContextTimeout) * time.Second

	taskRepo := taskRepository.NewTaskRepository(db)
	taskUsecase := taskUseCase.NewTaskUsecase(taskRepo, ctxTimeout)
	taskCtrl := taskController.New(taskUsecase)

	group.GET("/task/:taskID", taskCtrl.GetByID)
	group.POST("/task", taskCtrl.Create)
}
