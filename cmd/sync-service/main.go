package main

import (
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/lilpipidron/sync-service/internal/config"
	"github.com/lilpipidron/sync-service/internal/httpserver/handlers"
	"github.com/lilpipidron/sync-service/internal/storages/postgresql"
	"net/http"
	"strconv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	cfg := config.MustLoad()

	log.Info("Successfully loaded config", "config", *cfg)

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDB)

	storage, err := postgresql.NewPostgresqlStorage(dsn, cfg.PostgresDB)
	if err != nil {
		panic(err)
	}

	log.Info("Successfully connected to postgresql storage")

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	router.Post("/addClient", handlers.AddClientHandler(storage))
	router.Put("/updateClient", handlers.UpdateClientHandler(storage))
	router.Delete("/deleteClient/{id}", handlers.DeleteClientHandler(storage))
	router.Put("/updateAlgorithmStatus", nil)

	addr := cfg.ServiceHost + ":" + strconv.Itoa(cfg.ServicePort)
	if err = http.ListenAndServe(addr, router); err != nil {
		panic(err)
	}

}
