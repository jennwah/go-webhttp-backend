package route

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"

	"github.com/jennwah/go-webhttp-backend/api/route/healthcheck"
	"github.com/jennwah/go-webhttp-backend/api/route/task"
	"github.com/jennwah/go-webhttp-backend/config"
)

func Setup(env *config.Env, db *sqlx.DB, kv *redis.Client, gin *gin.Engine, logger *log.Entry) {
	v1Router := gin.Group("api/v1")

	// healthcheck routes
	healthcheck.NewHealthCheckRouter(v1Router)
	// task routes
	task.NewTaskRouter(env, db, kv, v1Router, logger)
}
