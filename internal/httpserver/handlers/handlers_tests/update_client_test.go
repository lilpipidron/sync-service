package handlers_test

import (
	"bytes"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/lilpipidron/sync-service/internal/httpserver/handlers"
	"github.com/lilpipidron/sync-service/internal/httpserver/requests"
	"github.com/lilpipidron/sync-service/internal/models"
	"github.com/lilpipidron/sync-service/internal/storages/postgresql"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestUpdateClientHandler(t *testing.T) {
	db := postgresql.SetupTestDB(t)
	storage := &postgresql.PostgresqlStorage{DB: db}

	router := chi.NewRouter()
	router.Put("/update-client/{id}", handlers.UpdateClientHandler(storage))

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

	updateClientRequest := requests.UpdateClientRequest{
		ID:          client.ID,
		ClientName:  "Updated Client",
		Version:     2,
		Image:       "updated-image",
		CPU:         "2",
		Memory:      "2Gi",
		NeedRestart: true,
	}

	requestBody, err := json.Marshal(updateClientRequest)
	if err != nil {
		t.Fatalf("failed to marshal update request body: %v", err)
	}

	req := httptest.NewRequest("PUT", "/update-client/"+strconv.FormatUint(uint64(client.ID), 10), bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status %d; got %d", http.StatusOK, w.Code)
	}

	var updatedClient models.Client
	if err = db.First(&updatedClient, client.ID).Error; err != nil {
		t.Fatalf("failed to find updated client: %v", err)
	}

	assert.Equal(t, "Updated Client", updatedClient.ClientName)
	assert.Equal(t, "updated-image", updatedClient.Image)
	assert.Equal(t, "2", updatedClient.CPU)
	assert.Equal(t, "2Gi", updatedClient.Memory)
	assert.Equal(t, true, updatedClient.NeedRestart)
}
