package v1

import (
	"github.com/gin-gonic/gin"
	"hw8/internal/crm_core/controller/http/middleware"
	"hw8/internal/crm_core/service"
	"hw8/pkg/crm_core/logger"
)

type ticketRoutes struct {
	s *service.Service
	l *logger.Logger
}

func newTicketRoutes(handler *gin.RouterGroup, s *service.Service, l *logger.Logger, MW *middleware.Middleware) {
	r := &ticketRoutes{s, l}

	ticketHandler := handler.Group("/ticket")
	{
		//middleware for users
		ticketHandler.GET("/", r.getTickets)
		ticketHandler.GET("/:id", r.getTicket)
		ticketHandler.POST("/", r.createTicket)
		ticketHandler.PUT("/:id", r.updateTicket)
		ticketHandler.DELETE("/:id", r.deleteTicket)
	}
}

func (tr *ticketRoutes) getTickets(ctx *gin.Context) {

}
func (tr *ticketRoutes) getTicket(ctx *gin.Context) {

}
func (tr *ticketRoutes) createTicket(ctx *gin.Context) {

}
func (tr *ticketRoutes) updateTicket(ctx *gin.Context) {

}
func (tr *ticketRoutes) deleteTicket(ctx *gin.Context) {

}
