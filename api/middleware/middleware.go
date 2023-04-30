package middleware

import "github.com/gin-gonic/gin"

func GlobalMiddlewareSetup(ginServer *gin.Engine) {
	ginServer.Use(CorsMiddleware())
}
