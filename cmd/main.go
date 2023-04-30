package main

import (
	"github.com/gin-gonic/gin"

	"github.com/jennwah/go-webhttp-backend/api/middleware"
	"github.com/jennwah/go-webhttp-backend/api/route"
	"github.com/jennwah/go-webhttp-backend/config"
	"github.com/jennwah/go-webhttp-backend/packages/logger"
	"github.com/jennwah/go-webhttp-backend/packages/mysql"
)

func main() {
	env := config.NewEnv()
	logger := logger.New()
	logger.Info("started logger")

	db, err := mysql.NewClient(mysql.Config{
		Host: env.DBHost,
		Port: env.DBPort,
		User: env.DBUser,
		Pass: env.DBPass,
		Name: env.DBName,
	})
	if err != nil {
		logger.Fatalf("fail to init db. Err: %v", err)
	}
	defer db.Close()

	ginServer := gin.Default()
	middleware.GlobalMiddlewareSetup(ginServer)
	route.Setup(env, db, ginServer)
	ginServer.Run(env.ServerAddress)
}
