package healthcheck

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewHealthCheckRouter(group *gin.RouterGroup) {
	group.GET("/healthcheck", healthcheck)
}

func healthcheck(c *gin.Context) {
	c.JSON(http.StatusOK, "healthcheck successful")
}
