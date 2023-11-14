package postgres

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"hw8/config/auth"
	"hw8/pkg/auth/logger"
)

func ConnectDB(config *auth.Configuration, l *logger.Logger) *gorm.DB {
	connectionStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost,
		config.DBPort,
		config.DBUser,
		config.DBPassword,
		config.DBName,
	)
	db, err := gorm.Open(postgres.Open(connectionStr), &gorm.Config{})
	if err != nil {
		l.Fatal(err)
	}
	return db
}
