package v1

import (
	"github.com/gin-gonic/gin"
	"hw8/internal/crm_core/controller/http/middleware"
	"hw8/internal/crm_core/entity"
	"hw8/internal/crm_core/service"
	"hw8/pkg/crm_core/logger"
	"net/http"
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

func (cr *contactRoutes) getContacts(ctx *gin.Context) {
	contacts, err := cr.s.GetContacts(ctx)

	if err != nil {
		ctx.JSON(http.StatusNotFound, &entity.CustomResponse{
			Status:  -1,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &entity.CustomResponseWithData{
		Status:  0,
		Message: "OK",
		Data:    contacts,
	})
}
func (cr *contactRoutes) getContact(ctx *gin.Context) {
	id := ctx.Param("id")

	contact, err := cr.s.GetContact(ctx, id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, &entity.CustomResponse{
			Status:  -1,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &entity.CustomResponseWithData{
		Status:  0,
		Message: "OK",
		Data:    contact,
	})
}
func (cr *contactRoutes) createContact(ctx *gin.Context) {
	var contact entity.Contact

	if err := ctx.ShouldBindJSON(&contact); err != nil {
		ctx.JSON(http.StatusBadRequest, &entity.CustomResponse{
			Status:  -1,
			Message: err.Error(),
		})
		return
	}

	if err := cr.s.CreateContact(ctx, contact); err != nil {
		ctx.JSON(http.StatusInternalServerError, &entity.CustomResponse{
			Status:  -2,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &entity.CustomResponse{
		Status:  0,
		Message: "OK",
	})
}
func (cr *contactRoutes) updateContact(ctx *gin.Context) {
	id := ctx.Param("id")

	var newContact entity.Contact

	if err := ctx.ShouldBindJSON(&newContact); err != nil {
		ctx.JSON(http.StatusBadRequest, &entity.CustomResponse{
			Status:  -1,
			Message: err.Error(),
		})
		return
	}

	if err := cr.s.UpdateContact(ctx, newContact, id); err != nil {
		ctx.JSON(http.StatusInternalServerError, &entity.CustomResponse{
			Status:  -2,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &entity.CustomResponse{
		Status:  0,
		Message: "OK",
	})
}
func (cr *contactRoutes) deleteContact(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := cr.s.DeleteContact(ctx, id); err != nil {
		ctx.JSON(http.StatusInternalServerError, &entity.CustomResponse{
			Status:  -1,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &entity.CustomResponse{
		Status:  0,
		Message: "OK",
	})
}
