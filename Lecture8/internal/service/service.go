package service

import (
	"hw8/config"
	"hw8/internal/repository"
	"hw8/pkg/logger"
)

type Service struct {
	Repo   *repository.Repo
	Config *config.Configuration
}

func New(config *config.Configuration, repo *repository.Repo, l *logger.Logger) *Service {
	return &Service{
		Repo:   repo,
		Config: config,
	}
}
