package repository

import (
	"gorm.io/gorm"
	"hw8/config/auth"
	"hw8/pkg/auth/logger"
	"hw8/pkg/auth/postgres"
)

type AuthRepo struct {
	DB *gorm.DB
}

func New(config *auth.Configuration, l *logger.Logger) *AuthRepo {
	db := postgres.ConnectDB(config, l)
	return &AuthRepo{
		DB: db,
	}
}
