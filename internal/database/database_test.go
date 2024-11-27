package database

import (
	"os"
	"testing"
)

func TestNewDatabaseSQLite(t *testing.T) {
	db, err := NewDatabase("sqlite", ":memory:") // Use in-memory SQLite for testing
	if err != nil {
		t.Fatalf("Failed to initialize SQLite through factory: %v", err)
	}
	defer db.Close()

	// Validate that the returned value implements DBHandler
	if db == nil {
		t.Fatal("Expected non-nil DBHandler")
	}
}

func TestNewDatabasePostgres(t *testing.T) {
	dsn := os.Getenv("POSTGRES_TEST_DSN")
	if dsn == "" {
		t.Skip("Skipping PostgreSQL test: POSTGRES_TEST_DSN not set")
	}

	db, err := NewDatabase("postgres", dsn)
	if err != nil {
		t.Fatalf("Failed to initialize PostgreSQL through factory: %v", err)
	}
	defer db.Close()

	// Validate that the returned value implements DBHandler
	if db == nil {
		t.Fatal("Expected non-nil DBHandler")
	}
}

func TestNewDatabaseInvalidDriver(t *testing.T) {
	_, err := NewDatabase("invalid", "dummy_dsn")
	if err == nil {
		t.Fatal("Expected error for invalid driver, but got none")
	}
}
