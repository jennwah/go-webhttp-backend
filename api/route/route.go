package route

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"

	"github.com/jennwah/go-webhttp-backend/api/route/healthcheck"
	"github.com/jennwah/go-webhttp-backend/api/route/task"
	"github.com/jennwah/go-webhttp-backend/config"
)

func Setup(env *config.Env, db *sqlx.DB, gin *gin.Engine) {
	v1Router := gin.Group("api/v1")

	// healthcheck routes
	healthcheck.NewHealthCheckRouter(v1Router)
	// task routes
	task.NewTaskRouter(env, db, v1Router)
}
