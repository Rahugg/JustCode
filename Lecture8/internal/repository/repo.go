package repository

import (
	"gorm.io/gorm"
	"hw8/config"
	"hw8/pkg/logger"
	"hw8/pkg/postgres"
)

type Repo struct {
	DB *gorm.DB
}

func New(config *config.Configuration, l *logger.Logger) *Repo {
	db := postgres.ConnectDB(config, l)
	return &Repo{
		DB: db,
	}
}
