package main

import (
	"github.com/gin-gonic/gin"

	"github.com/jennwah/go-webhttp-backend/api/middleware"
	"github.com/jennwah/go-webhttp-backend/api/route"
	"github.com/jennwah/go-webhttp-backend/config"
	"github.com/jennwah/go-webhttp-backend/packages/logger"
	"github.com/jennwah/go-webhttp-backend/packages/mysql"
	"github.com/jennwah/go-webhttp-backend/packages/redis"
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

	redisStore, err := redis.NewClient(redis.Config{
		Host: env.RedisHost,
		Port: env.RedisPort,
		Pass: env.RedisPass,
	})
	if err != nil {
		logger.Errorf("fail to init redis store. Err: %v", err)
	}
	defer redisStore.Close()

	ginServer := gin.Default()
	middleware.GlobalMiddlewareSetup(ginServer)
	route.Setup(env, db, redisStore, ginServer, logger)
	ginServer.Run(env.ServerAddress)
}
