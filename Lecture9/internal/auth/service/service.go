package service

import (
	"hw8/config/auth"
	"hw8/internal/auth/repository"
	"hw8/pkg/auth/logger"
)

type Service struct {
	Repo   *repository.AuthRepo
	Config *auth.Configuration
}

func New(config *auth.Configuration, repo *repository.AuthRepo, l *logger.Logger) *Service {
	return &Service{
		Repo:   repo,
		Config: config,
	}
}
