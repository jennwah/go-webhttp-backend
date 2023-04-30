package route

import (
	"github.com/jmoiron/sqlx"

	"github.com/gin-gonic/gin"

	"github.com/jennwah/go-webhttp-backend/config"
)

func Setup(env *config.Env, db *sqlx.DB, gin *gin.Engine) {
	router := gin.Group("")

	NewTaskRouter(env, db, router)
}
