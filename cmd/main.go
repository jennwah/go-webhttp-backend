package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/jennwah/go-webhttp-backend/api/route"
	"github.com/jennwah/go-webhttp-backend/config"
	"github.com/jennwah/go-webhttp-backend/packages/mysql"
)

func main() {
	env := config.NewEnv()

	db, err := mysql.NewClient(mysql.Config{
		Host: env.DBHost,
		Port: env.DBPort,
		User: env.DBUser,
		Pass: env.DBPass,
		Name: env.DBName,
	})
	if err != nil {
		log.Fatalf("fail to init db. Err: %v", err)
	}
	defer db.Close()

	ginServer := gin.Default()
	route.Setup(env, db, ginServer)
	ginServer.Run(env.ServerAddress)
}
