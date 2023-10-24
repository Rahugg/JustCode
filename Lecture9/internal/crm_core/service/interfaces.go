package service

import (
	"github.com/gin-gonic/gin"
	"hw8/internal/crm_core/entity"
)

type (
	CompanyService interface {
		GetCompanies(ctx *gin.Context) (*[]entity.Company, error)
		GetCompany(ctx *gin.Context, id string) (*entity.Company, error)
		CreateCompany(ctx *gin.Context, company entity.Company) error
		UpdateCompany(ctx *gin.Context, newCompany entity.NewCompany, id string) error
		DeleteCompany(ctx *gin.Context, id string) error
	}
)
