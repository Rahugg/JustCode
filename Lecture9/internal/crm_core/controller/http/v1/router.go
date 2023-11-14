package v1

import (
	"github.com/gin-gonic/gin"
	"hw8/internal/crm_core/controller/http/middleware"
	"hw8/internal/crm_core/service"
	"hw8/pkg/crm_core/cache"
	"hw8/pkg/crm_core/logger"
	"net/http"
)

func NewRouter(handler *gin.Engine, s *service.Service, l *logger.Logger, MW *middleware.Middleware, cc cache.Contact) {

	// Health Check
	handler.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello to auth service",
		})
	})

	h := handler.Group("/v1")
	{
		newCompanyRoutes(h, s, l, MW)
		newContactRoutes(h, s, l, MW, cc)
		newDealRoutes(h, s, l, MW)
		newTaskRoutes(h, s, l, MW)
		newTicketRoutes(h, s, l, MW)
	}
}
