package v1

import (
	"github.com/gin-gonic/gin"
	"hw8/internal/crm_core/controller/http/middleware"
	"hw8/internal/crm_core/entity"
	"hw8/internal/crm_core/service"
	"hw8/pkg/crm_core/logger"
	"net/http"
)

type companyRoutes struct {
	s *service.Service
	l *logger.Logger
}

func newCompanyRoutes(handler *gin.RouterGroup, s *service.Service, l *logger.Logger, MW *middleware.Middleware) {
	r := &companyRoutes{s, l}

	companyHandler := handler.Group("/company")
	{
		//middleware for users
		companyHandler.GET("/", r.getCompanies)
		companyHandler.GET("/:id", r.getCompany)
		companyHandler.POST("/", r.createCompany)
		companyHandler.PUT("/:id", r.updateCompany)
		companyHandler.DELETE("/:id", r.deleteCompany)
	}
}

func (cr *companyRoutes) getCompanies(ctx *gin.Context) {
	companies, err := cr.s.GetCompanies(ctx)

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
		Data:    companies,
	})
}

func (cr *companyRoutes) getCompany(ctx *gin.Context) {
	id := ctx.Param("id")

	company, err := cr.s.GetCompany(ctx, id)

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
		Data:    company,
	})
}

func (cr *companyRoutes) createCompany(ctx *gin.Context) {
	var company entity.Company

	if err := ctx.ShouldBindJSON(&company); err != nil {
		ctx.JSON(http.StatusBadRequest, &entity.CustomResponse{
			Status:  -1,
			Message: err.Error(),
		})
		return
	}

	if err := cr.s.CreateCompany(ctx, company); err != nil {
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

func (cr *companyRoutes) updateCompany(ctx *gin.Context) {
	id := ctx.Param("id")

	var newCompany entity.NewCompany

	if err := ctx.ShouldBindJSON(&newCompany); err != nil {
		ctx.JSON(http.StatusBadRequest, &entity.CustomResponse{
			Status:  -1,
			Message: err.Error(),
		})
		return
	}

	if err := cr.s.UpdateCompany(ctx, newCompany, id); err != nil {
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

func (cr *companyRoutes) deleteCompany(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := cr.s.DeleteCompany(ctx, id); err != nil {
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
