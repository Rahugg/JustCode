package v1

import (
	"github.com/gin-gonic/gin"
	"hw8/internal/auth/controller/http/middleware"
	"hw8/internal/auth/service"
	"hw8/pkg/auth/logger"
	"net/http"
)

func NewRouter(handler *gin.Engine, s *service.Service, l *logger.Logger, MW *middleware.Middleware) {

	// Health Check
	handler.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello to auth service",
		})
	})

	h := handler.Group("/v1")
	{
		newUserRoutes(h, s, l, MW)
	}
}
