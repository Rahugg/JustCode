package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	entityRepo "hw8/internal/crm_core/entity"
	"hw8/pkg/crm_core/logger"
	"net/http/httptest"
	"testing"
)

func TestGetContacts(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	testConfig := getTestConfig()
	logger := logger.New("1")

	repo := New(testConfig, logger)
	testMigrate(repo, logger)

	contacts, err := repo.GetContacts(ctx)

	// Assertions
	assert.Nil(t, err, "Expected no error")
	assert.NotNil(t, contacts, "Expected contacts to be not nil")

	assert.Len(t, *contacts, 0, "Unexpected length of contacts slice")
}

func TestCreateContact(t *testing.T) {
	testConfig := getTestConfig()
	logger := logger.New("1")
	repo := New(testConfig, logger)
	testMigrate(repo, logger)

	tests := []struct {
		name        string
		mockContact *entityRepo.Contact
		wantError   bool
	}{
		{
			name: "should insert contact successfully",
			mockContact: &entityRepo.Contact{
				FirstName: "TestCompany",
				LastName:  "Funny mocks",
				Email:     "testcompany@gmail.com",
				Phone:     "322",
			},
			wantError: false,
		},
		{
			name: "should insert another contact successfully",
			mockContact: &entityRepo.Contact{
				FirstName: "AnotherCompany",
				LastName:  "Different mocks",
				Email:     "testcompany2@gmail.com",
				Phone:     "789",
			},
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			ctx, _ := gin.CreateTestContext(nil)

			err := repo.CreateContact(ctx, tt.mockContact)

			if tt.wantError {
				assert.Error(t, err, "Expected an error")
			} else {
				assert.NoError(t, err, "Expected no error")
			}
		})
	}
}
