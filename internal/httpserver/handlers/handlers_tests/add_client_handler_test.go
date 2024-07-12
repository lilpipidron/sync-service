package handlers_test

import (
	"bytes"
	"encoding/json"
	"github.com/lilpipidron/sync-service/internal/httpserver/handlers"
	"github.com/lilpipidron/sync-service/internal/httpserver/requests"
	"github.com/lilpipidron/sync-service/internal/models"
	"github.com/lilpipidron/sync-service/internal/storages/postgresql"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddClientHandler(t *testing.T) {
	db := postgresql.SetupTestDB(t)
	storage := &postgresql.PostgresqlStorage{DB: db}

	handler := handlers.AddClientHandler(storage)

	client := requests.AddClientRequest{
		ClientName:  "test-client",
		Version:     1,
		Image:       "test-image",
		CPU:         "500m",
		Memory:      "512Mi",
		NeedRestart: false,
	}

	body, err := json.Marshal(client)
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", "/clients", bytes.NewBuffer(body))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)

	var createdClient models.Client
	if err := db.First(&createdClient).Error; err != nil {
		t.Fatalf("Failed to find created client: %v", err)
	}

	assert.Equal(t, "test-client", createdClient.ClientName)
	assert.Equal(t, 1, createdClient.Version)
	assert.Equal(t, "test-image", createdClient.Image)
}
