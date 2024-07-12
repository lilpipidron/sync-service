package handlers_test

import (
	"bytes"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/lilpipidron/sync-service/internal/httpserver/handlers"
	"github.com/lilpipidron/sync-service/internal/httpserver/requests"
	"github.com/lilpipidron/sync-service/internal/models"
	"github.com/lilpipidron/sync-service/internal/storages/postgresql"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestUpdateAlgorithmStatusHandler(t *testing.T) {
	db := postgresql.SetupTestDB(t)
	defer db.Migrator().DropTable(&models.Client{})

	storage := &postgresql.PostgresqlStorage{DB: db}

	router := chi.NewRouter()
	router.Put("/updateAlgorithmStatus", handlers.UpdateAlgorithmStatusHandler(storage))

	client := models.Client{
		ClientName:  "Test Client",
		Version:     1,
		Image:       "test-image",
		CPU:         "1",
		Memory:      "1Gi",
		NeedRestart: false,
	}

	if err := db.Create(&client).Error; err != nil {
		t.Fatalf("failed to create test client: %v", err)
	}

	algorithm := models.AlgorithmStatus{
		ID:       1,
		ClientID: client.ID,
		VWAP:     false,
		TWAP:     false,
		HFT:      false,
	}

	if err := db.Create(&algorithm).Error; err != nil {
		t.Fatalf("failed to create test algorithm status: %v", err)
	}

	updateAlgorithmStatusRequest := requests.UpdateAlgorithmStatusRequest{
		ID:       1,
		ClientID: 1,
		VWAP:     true,
		TWAP:     false,
		HFT:      true,
	}
	requestBody, err := json.Marshal(updateAlgorithmStatusRequest)
	if err != nil {
		t.Fatalf("failed to marshal update request body: %v", err)
	}

	req := httptest.NewRequest("PUT", "/updateAlgorithmStatus", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status %d; got %d", http.StatusOK, w.Code)
	}

	var updatedAlgorithmStatus models.AlgorithmStatus
	if err := db.First(&updatedAlgorithmStatus, 1).Error; err != nil {
		t.Fatalf("failed to find updated Algorithm: %v", err)
	}

	expectedAlgorithm := models.AlgorithmStatus{
		ID:       1,
		ClientID: client.ID,
		VWAP:     true,
		TWAP:     false,
		HFT:      true,
	}
	if !reflect.DeepEqual(expectedAlgorithm, updatedAlgorithmStatus) {
		t.Errorf("expected updated algorithm status %+v; got %+v", expectedAlgorithm, updatedAlgorithmStatus)
	}
}
