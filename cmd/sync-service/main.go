package main

import (
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/lilpipidron/sync-service/internal/config"
	"github.com/lilpipidron/sync-service/internal/storages/postgresql"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	cfg := config.MustLoad()

	log.Info("Successfully loaded config", "config", *cfg)

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDB)

	_, err := postgresql.NewPostgresqlStorage(dsn, cfg.PostgresDB)
	if err != nil {
		panic(err)
	}

	log.Info("Successfully connected to postgresql storage")

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	router.Post("/addClient", nil)
	router.Put("/updateClient", nil)
	router.Delete("/deleteClient", nil)
	router.Put("/updateAlgorithmStatus", nil)

}
