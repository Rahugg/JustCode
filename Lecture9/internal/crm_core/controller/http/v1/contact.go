package v1

import (
	"github.com/gin-gonic/gin"
	"hw8/internal/crm_core/controller/http/middleware"
	"hw8/internal/crm_core/service"
	"hw8/pkg/crm_core/logger"
)

type contactRoutes struct {
	s *service.Service
	l *logger.Logger
}

func newContactRoutes(handler *gin.RouterGroup, s *service.Service, l *logger.Logger, MW *middleware.Middleware) {
	r := &contactRoutes{s, l}

	contactHandler := handler.Group("/contact")
	{
		//middleware for users
		contactHandler.GET("/", r.getContacts)
		contactHandler.GET("/:id", r.getContact)
		contactHandler.POST("/", r.createContact)
		contactHandler.PUT("/:id", r.updateContact)
		contactHandler.DELETE("/:id", r.deleteContact)
	}
}

func (tr *contactRoutes) getContacts(ctx *gin.Context) {

}
func (tr *contactRoutes) getContact(ctx *gin.Context) {

}
func (tr *contactRoutes) createContact(ctx *gin.Context) {

}
func (tr *contactRoutes) updateContact(ctx *gin.Context) {

}
func (tr *contactRoutes) deleteContact(ctx *gin.Context) {

}
