package repository

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"hw8/config/crm_core"
	entityRepo "hw8/internal/crm_core/entity"
	"hw8/pkg/crm_core/logger"
	"net/http/httptest"
	"testing"
)

func testMigrate(repo *CRMSystemRepo, logger *logger.Logger) {
	repo.DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")

	err := repo.DB.AutoMigrate(
		&entityRepo.Company{},
		&entityRepo.Contact{},
	)
	if err != nil {
		logger.Fatal("Automigration failed")
	}

	fmt.Println("👍 Migration complete")

}

func getTestConfig() *crm_core.Configuration {
	return &crm_core.Configuration{
		App: crm_core.App{
			Name:    "YourAppName",
			Version: "1.0.0",
		},
		HTTP: crm_core.HTTP{
			Port:                   "8080",
			DebugPort:              "8081",
			DefaultReadTimeout:     10,
			DefaultWriteTimeout:    10,
			DefaultShutdownTimeout: 10,
		},
		Log: crm_core.Log{
			Level: "debug",
		},
		Gin: crm_core.Gin{
			Mode: "release",
		},
		DB: crm_core.DB{
			PoolMax:  10,
			Host:     "localhost",
			User:     "postgres",
			Password: "12345",
			Name:     "testdb",
			Port:     5432,
		},
		Transport: crm_core.Transport{
			Validate: crm_core.ValidateTransport{
				Host:    "validatehost",
				Timeout: 5,
			},
		},
		Jwt: crm_core.Jwt{
			AccessPrivateKey:      "testaccessprivatekey",
			AccessPublicKey:       "testaccesspublickey",
			AccessTokenExpiredIn:  3600,
			AccessTokenMaxAge:     1800,
			RefreshPrivateKey:     "testrefreshprivatekey",
			RefreshPublicKey:      "testrefreshpublickey",
			RefreshTokenExpiredIn: 86400,
			RefreshTokenMaxAge:    2592000,
		},
	}
}

func TestGetCompanies(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	testConfig := getTestConfig()
	logger := logger.New("1")

	repo := New(testConfig, logger)
	testMigrate(repo, logger)

	companies, err := repo.GetCompanies(ctx)

	assert.Nil(t, err, "Expected no error")
	assert.NotNil(t, companies, "Expected companies to be not nil")

	assert.Len(t, *companies, 0, "Unexpected length of companies slice")
}

func TestCreateCompany(t *testing.T) {
	testConfig := getTestConfig()
	logger := logger.New("1")
	repo := New(testConfig, logger)
	testMigrate(repo, logger)

	tests := []struct {
		name        string
		mockCompany *entityRepo.Company
		wantError   bool
	}{
		{
			name: "should insert company successfully",
			mockCompany: &entityRepo.Company{
				Name:    "TestCompany",
				Address: "Funny mocks",
				Phone:   "322",
			},
			wantError: false,
		},
		{
			name: "should insert another company successfully",
			mockCompany: &entityRepo.Company{
				Name:    "AnotherCompany",
				Address: "Different mocks",
				Phone:   "789",
			},
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			ctx, _ := gin.CreateTestContext(nil)

			err := repo.CreateCompany(ctx, tt.mockCompany)

			if tt.wantError {
				assert.Error(t, err, "Expected an error")
			} else {
				assert.NoError(t, err, "Expected no error")
			}
		})
	}
}
