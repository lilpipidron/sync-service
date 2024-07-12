package handlers

import (
	"github.com/charmbracelet/log"
	"github.com/go-chi/render"
	"github.com/lilpipidron/sync-service/internal/httpserver/requests"
	"github.com/lilpipidron/sync-service/internal/models"
	"github.com/lilpipidron/sync-service/internal/storages/postgresql"
	"net/http"
	"time"
)

// AddClientHandler adds a new client to the database based on the information provided in the request body.
// If the request cannot be decoded, it logs an error and returns a 400 status code.
// Returns a 500 status code and logs an error if there is an issue adding the client to the database.
func AddClientHandler(storage *postgresql.PostgresqlStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var addClientRequest requests.AddClientRequest
		var req interface{} = &addClientRequest

		if err := requests.Decode(w, r, &req); err != nil {
			log.Error("Failed to decode request", "error", err)
			render.JSON(w, r, http.StatusBadRequest)
		}

		client := models.Client{
			ClientName:  addClientRequest.ClientName,
			Version:     addClientRequest.Version,
			Image:       addClientRequest.Image,
			CPU:         addClientRequest.CPU,
			Memory:      addClientRequest.Memory,
			NeedRestart: addClientRequest.NeedRestart,
			SpawnedAt:   time.Time{},
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		if err := storage.DB.Create(&client).Error; err != nil {
			render.JSON(w, r, http.StatusInternalServerError)
			log.Error("Failed to add client", "error", err)
			return
		}

		render.Status(r, http.StatusCreated)
		render.JSON(w, r, nil)

		log.Info("Added client", "client", client)
	}
}
