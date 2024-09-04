package main

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/yugabytedb"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq" // PostgreSQL driver for YugabyteDB
)

func main() {
	// Open the database connection
	db, err := sql.Open("postgres", "postgres://yugabyte:yugabyte@localhost:5433/test_db?sslmode=disable")
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	// Initialize the migration driver specific for YugabyteDB
	driver, err := yugabytedb.WithInstance(db, &yugabytedb.Config{})
	if err != nil {
		log.Fatalf("Failed to create YugabyteDB migration driver: %v", err)
	}

	// Create a new migrate instance
	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		"yugabytedb", driver)
	if err != nil {
		log.Fatalf("Failed to create migrate instance: %v", err)
	}

	// Apply the migrations
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Migration failed: %v", err)
	}

	log.Println("Migrations applied successfully!")
}
