package main

import (
	"api-env-example/cmd"
	"api-env-example/infra"
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	env := infra.NewConfig()

	runMigrations()

	ctx := context.Background()
	cmd.StartHttp(ctx, env)
}

func runMigrations() {
	dbURL := "postgres://root:6S492Mm4uwZR@localhost:5432/db_test?sslmode=disable" // substitua pelo seu URL de conexão

	// Verifique se a URL de conexão está no formato correto
	if dbURL == "" {
		log.Fatalf("Database URL is empty")
	}
	fmt.Println("Connecting to database at URL:", dbURL)

	// Open a database connection to check if the URL is correct
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Failed to open database connection: %v", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Initialize the migrate instance
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("Failed to create migration driver: %v", err)
	}

	// Get absolute path for migrations directory
	pwd, _ := os.Getwd()
	migrationsPath := filepath.Join(pwd, "migration")
	migrationsURL := fmt.Sprintf("file://%s", migrationsPath)
	fmt.Println("Migrations URL:", migrationsURL)

	m, err := migrate.NewWithDatabaseInstance(
		migrationsURL,
		"postgres", driver)
	if err != nil {
		log.Fatalf("Failed to create migration instance: %v", err)
	}

	// Run the migrations
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Failed to apply migrations: %v", err)
	}

	log.Println("Migrations applied successfully!")
}
