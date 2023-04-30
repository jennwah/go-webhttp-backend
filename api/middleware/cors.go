package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const cacheMaxAge = 24 * time.Hour

func CorsMiddleware() gin.HandlerFunc {
	cfg := cors.Config{
		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"PATCH",
			"DELETE",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Origin",
			"Content-Length",
			"Content-Type",
			"User-Agent",
			"Referrer",
			"Host",
			"Token",
			"Authorization",
		},
		AllowCredentials: true,
		MaxAge:           cacheMaxAge,
		AllowOrigins:     []string{"*"},
	}

	return cors.New(cfg)
}
