package handlers

import (
	"github.com/charmbracelet/log"
	"github.com/go-chi/render"
	"github.com/lilpipidron/sync-service/internal/httpserver/requests"
	"github.com/lilpipidron/sync-service/internal/models"
	"github.com/lilpipidron/sync-service/internal/storages/postgresql"
	"net/http"
	"reflect"
	"time"
)

func UpdateClientHandler(storage *postgresql.PostgresqlStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var updateClientRequest requests.UpdateClientRequest
		var req interface{} = &updateClientRequest

		if err := requests.Decode(w, r, &req); err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, render.M{"error": err})
			return
		}

		log.Info("Request body", "body", updateClientRequest)

		client := models.Client{
			UpdatedAt: time.Now(),
		}
		if err := storage.DB.Where("id = ?", updateClientRequest.ID).First(&client).Error; err != nil {
			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, nil)
			log.Error("Client not found", "error", err)
			return
		}

		clientVal := reflect.ValueOf(client).Elem()
		updateClientVal := reflect.ValueOf(req).Elem()
		for i := 0; i < updateClientVal.NumField(); i++ {
			field := updateClientVal.Field(i)
			fieldName := updateClientVal.Type().Field(i).Name

			if field.Kind() == reflect.String && field.String() != "" {
				clientField := clientVal.FieldByName(fieldName)

				if clientField.IsValid() && clientField.CanSet() {
					clientField.SetString(field.String())
				}
			}
		}

		if err := storage.DB.Save(&client).Error; err != nil {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, err.Error())
			log.Error("Unable to update client", "error", err)
			return
		}

		render.Status(r, http.StatusOK)
		render.JSON(w, r, client)
		log.Info("Successfully updated client", "client", client)
	}
}
