package checker

import (
	"github.com/charmbracelet/log"
	"github.com/lilpipidron/sync-service/internal/deployer"
	"github.com/lilpipidron/sync-service/internal/models"
	"github.com/lilpipidron/sync-service/internal/storages/postgresql"
	"k8s.io/apimachinery/pkg/api/errors"
	"strconv"
	"time"
)

func StatusChecker(podDeployer deployer.Deployer, storage *postgresql.PostgresqlStorage) {
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			var algorithms []models.AlgorithmStatus
			if err := storage.DB.First(&algorithms).Error; err != nil {
				log.Error("Failed to get algorithm statuses", "error", err)
				return
			}

			for _, algorithm := range algorithms {
				var client models.Client
				if err := storage.DB.Where("id = ?", algorithm.ClientID).First(&client).Error; err != nil {
					log.Error("Failed to get client", "error", err)
					continue
				}
				if algorithm.VWAP {
					if err := podDeployer.CreatePod(strconv.Itoa(int(client.ID)) + "-vwap"); err != nil && !errors.IsAlreadyExists(err) {
						log.Error("Failed to create pod", "error", err)
					}
				} else {
					if err := podDeployer.DeletePod(strconv.Itoa(int(client.ID)) + "-vwap"); err != nil && !errors.IsNotFound(err) {
						log.Error("Failed to delete pod", "error", err)
					}
				}

				if algorithm.TWAP {
					if err := podDeployer.CreatePod(strconv.Itoa(int(client.ID)) + "-twap"); err != nil && !errors.IsAlreadyExists(err) {
						log.Error("Failed to create pod", "error", err)
					}
				} else {
					if err := podDeployer.DeletePod(strconv.Itoa(int(client.ID)) + "-twap"); err != nil && !errors.IsNotFound(err) {
						log.Error("Failed to delete pod", "error", err)
					}
				}

				if algorithm.HFT {
					if err := podDeployer.CreatePod(strconv.Itoa(int(client.ID)) + "-hft"); err != nil && !errors.IsAlreadyExists(err) {
						log.Error("Failed to create pod", "error", err)
					}
				} else {
					if err := podDeployer.DeletePod(strconv.Itoa(int(client.ID)) + "-hft"); err != nil && !errors.IsNotFound(err) {
						log.Error("Failed to delete pod", "error", err)
					}
				}
			}
		}
	}
}
