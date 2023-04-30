package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"

	"github.com/jennwah/go-webhttp-backend/api/controller"
	"github.com/jennwah/go-webhttp-backend/config"
	"github.com/jennwah/go-webhttp-backend/repository"
	"github.com/jennwah/go-webhttp-backend/usecase"
)

func NewTaskRouter(env *config.Env, db *sqlx.DB, group *gin.RouterGroup) {
	ctxTimeout := time.Duration(env.ContextTimeout) * time.Second

	taskRepo := repository.NewTaskRepository(db)
	taskController := &controller.TaskController{
		TaskUsecase: usecase.NewTaskUsecase(taskRepo, ctxTimeout),
	}

	group.GET("/task/:taskID", taskController.GetByID)
	group.POST("/task", taskController.Create)
}
