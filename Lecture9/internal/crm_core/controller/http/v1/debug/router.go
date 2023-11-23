package debug

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"hw8/pkg/crm_core/logger"
	"net/http"
)

func NewDebugRouter(handler *gin.Engine, l *logger.Logger) {

	// Health Check
	handler.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello to auth service",
		})
	})

	pprof.Register(handler)

}
