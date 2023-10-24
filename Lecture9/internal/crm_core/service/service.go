package service

import (
	"hw8/config/crm_core"
	"hw8/internal/crm_core/repository"
	"hw8/pkg/crm_core/logger"
)

type Service struct {
	Repo   *repository.CRMSystemRepo
	Config *crm_core.Configuration
}

func New(config *crm_core.Configuration, repo *repository.CRMSystemRepo, l *logger.Logger) *Service {
	return &Service{
		Repo:   repo,
		Config: config,
	}
}
