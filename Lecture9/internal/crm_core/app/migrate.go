package app

import (
	"fmt"
	entityRepo "hw8/internal/crm_core/entity"
	"hw8/internal/crm_core/repository"
	"hw8/pkg/crm_core/logger"
)

func Migrate(repo *repository.CRMSystemRepo, l *logger.Logger) {
	repo.DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	err := repo.DB.AutoMigrate(
		&entityRepo.Company{},
		&entityRepo.Role{},
	)
	if err != nil {
		l.Fatal("Automigration failed")
	}

	fmt.Println("👍 Migration complete")
}
