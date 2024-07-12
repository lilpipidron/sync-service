package postgresql

import (
	"github.com/lilpipidron/sync-service/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func SetupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to the in-memory database: %v", err)
	}

	if err := db.AutoMigrate(&models.Client{}); err != nil {
		t.Fatalf("Failed to migrate database: %v", err)
	}

	if err := db.AutoMigrate(&models.AlgorithmStatus{}); err != nil {
		t.Fatalf("Failed to migrate database: %v", err)
	}
	return db
}
