package handlers

import (
	"github.com/charmbracelet/log"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/lilpipidron/sync-service/internal/models"
	"github.com/lilpipidron/sync-service/internal/storages/postgresql"
	"net/http"
)

func DeleteClientHandler(storage *postgresql.PostgresqlStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		clientID := chi.URLParam(r, "id")
		var client models.Client
		if err := storage.DB.First(&client, clientID).Error; err != nil {
			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, map[string]string{"error": err.Error()})
			log.Error("Client not found", "error", err)
			return
		}

		if err := storage.DB.Delete(&client).Error; err != nil {
			render.Status(r, http.StatusInternalServerError)
			log.Error("error deleting client", "error", err)
			return
		}

		render.Status(r, http.StatusNoContent)
		render.JSON(w, r, nil)
		log.Info("client deleted", "clientID", clientID)
	}
}
