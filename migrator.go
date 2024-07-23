package main

import (
	"database/sql"
	"fmt"
	"github.com/FoodMoodOTG/examplearch/services/config"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"log/slog"
	"os"
)

func migrationsUp(dbInstance *migrate.Migrate) error {
	if err := dbInstance.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to apply migrations due [%s]", err)
	}
	return nil
}

func migrationsDown(dbInstance *migrate.Migrate) error {
	if err := dbInstance.Down(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to rollback migrations due [%s]", err)
	}
	return nil
}

func getArgs() (string, error) {
	args := os.Args[1:]

	if args[0] != "up" && args[0] != "down" {
		return "", fmt.Errorf("invalid arguments, must be 'up' or 'down'")
	}

	return args[0], nil
}

func main() {
	cfg := config.Make()

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.PostgresUser(),
		cfg.PostgresPassword(),
		cfg.PostgresHost(),
		cfg.PostgresPort(),
		cfg.PostgresName(),
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		slog.Error("Failed to connect to the database", err)
		return
	}
	defer db.Close()

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		slog.Error("Failed to create database driver", err)
		return
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://_db/migrations",
		"postgres", driver)
	if err != nil {
		slog.Error("Failed to create migrate instance", err)
		return
	}

	arg, err := getArgs()
	if err != nil {
		slog.Error("failed to get argument due", err)
		return
	}

	switch arg {
	case "up":
		err = migrationsUp(m)
	case "down":
		err = migrationsDown(m)
	}
	if err != nil {
		slog.Error("Failed to migrate database instance", err)
		return
	}

	slog.Info("Migrations applied successfully!")
}
