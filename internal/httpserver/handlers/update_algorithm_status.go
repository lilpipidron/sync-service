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

func UpdateAlgorithmStatusHandler(storage *postgresql.PostgresqlStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var updateAlgorithmStatusRequest requests.UpdateAlgorithmStatusRequest
		var req interface{} = &updateAlgorithmStatusRequest

		if err := requests.Decode(w, r, &req); err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, render.M{"error": err})
			return
		}

		log.Info("Request body", "body", updateAlgorithmStatusRequest)

		algorithmStatus := models.Client{
			UpdatedAt: time.Now(),
		}
		if err := storage.DB.Where("id = ?", updateAlgorithmStatusRequest.ID).First(&algorithmStatus).Error; err != nil {
			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, nil)
			log.Error("Algorithm not found", "error", err)
			return
		}

		algorithmVal := reflect.ValueOf(algorithmStatus).Elem()
		updateAlgorithmVal := reflect.ValueOf(req).Elem()
		for i := 0; i < updateAlgorithmVal.NumField(); i++ {
			field := updateAlgorithmVal.Field(i)
			fieldName := updateAlgorithmVal.Type().Field(i).Name

			if field.Kind() == reflect.String && field.String() != "" {
				clientField := algorithmVal.FieldByName(fieldName)

				if clientField.IsValid() && clientField.CanSet() {
					clientField.SetString(field.String())
				}
			}
		}

		if err := storage.DB.Save(&algorithmStatus).Error; err != nil {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, err.Error())
			log.Error("Unable to update algorithm status", "error", err)
			return
		}

		render.Status(r, http.StatusOK)
		render.JSON(w, r, algorithmStatus)
		log.Info("Successfully updated algorithm status", "algorithmStatus", algorithmStatus)
	}
}
