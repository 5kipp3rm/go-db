package postgres

import (
	"os"
	"testing"
)

func TestPostgresConnection(t *testing.T) {
	dsn := os.Getenv("POSTGRES_TEST_DSN")
	if dsn == "" {
		t.Skip("Skipping PostgreSQL test: POSTGRES_TEST_DSN not set")
	}

	db, err := NewPostgres(dsn)
	if err != nil {
		t.Fatalf("Failed to initialize PostgreSQL: %v", err)
	}
	defer db.Close()

	_, err = db.DB.Exec("CREATE TABLE IF NOT EXISTS test (id SERIAL PRIMARY KEY, name TEXT)")
	if err != nil {
		t.Fatalf("Failed to execute query on PostgreSQL: %v", err)
	}
}
