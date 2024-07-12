package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/lilpipidron/sync-service/internal/httpserver/handlers"
	"github.com/lilpipidron/sync-service/internal/models"
	"github.com/lilpipidron/sync-service/internal/storages/postgresql"
	"github.com/stretchr/testify/assert"
)

func TestDeleteClientHandler(t *testing.T) {
	db := postgresql.SetupTestDB(t)
	storage := &postgresql.PostgresqlStorage{DB: db}

	client := models.Client{
		ClientName: "test-client",
		Version:    1,
		Image:      "test-image",
		CPU:        "500m",
		Memory:     "512Mi",
	}
	if err := db.Create(&client).Error; err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	handler := handlers.DeleteClientHandler(storage)

	r := chi.NewRouter()
	r.Delete("/clients/{id}", handler)

	req, err := http.NewRequest("DELETE", "/clients/"+strconv.Itoa(int(client.ID)), nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNoContent, rr.Code)

	var deletedClient models.Client
	if err := db.First(&deletedClient, client.ID).Error; err == nil {
		t.Fatalf("Client was not deleted")
	}
}
