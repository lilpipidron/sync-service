package main

import (
	"fmt"
	"github.com/charmbracelet/log"
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

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable",
		cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresUser, cfg.PostgresDB)

	_, err := postgresql.NewPostgresqlStorage(dsn, cfg.PostgresDB)
	if err != nil {
		panic(err)
	}

	log.Info("Successfully connected to postgresql storage")
}
