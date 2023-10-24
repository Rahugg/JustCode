package repository

import (
	"gorm.io/gorm"
	"hw8/config/crm_core"
	"hw8/pkg/crm_core/logger"
	"hw8/pkg/crm_core/postgres"
)

type CRMSystemRepo struct {
	DB *gorm.DB
}

func New(config *crm_core.Configuration, l *logger.Logger) *CRMSystemRepo {
	db := postgres.ConnectDB(config, l)
	return &CRMSystemRepo{
		DB: db,
	}
}
