package postgresql

import (
	"errors"
	"github.com/charmbracelet/log"
	"github.com/golang-migrate/migrate/v4"
	migratePostgres "github.com/golang-migrate/migrate/v4/database/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type PostgresqlStorage struct {
	DB *gorm.DB
}

func NewPostgresqlStorage(dsn, dbname string) (*PostgresqlStorage, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	log.Info("Successfully opened postgresql connection")

	postgresDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	driver, err := migratePostgres.WithInstance(postgresDB, &migratePostgres.Config{})
	if err != nil {
		return nil, err
	}

	migration, err := migrate.NewWithDatabaseInstance("file://internal/storages/postgresql/migrations", dbname, driver)
	if err != nil {
		return nil, err
	}

	if err = migration.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return nil, err
	}

	return &PostgresqlStorage{DB: db}, nil
}
