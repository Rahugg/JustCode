package repository

import (
	"github.com/gin-gonic/gin"
	"hw8/internal/crm_core/entity"
)

type (
	CompanyRepo interface {
		GetCompanies(ctx *gin.Context) (*[]entity.Company, error)
		GetCompany(ctx *gin.Context, id string) (*entity.Company, error)
		CreateCompany(ctx *gin.Context, company *entity.Company) error
		SaveCompany(ctx *gin.Context, newCompany entity.Company) error
		DeleteCompany(ctx *gin.Context, id string) error
	}
)
