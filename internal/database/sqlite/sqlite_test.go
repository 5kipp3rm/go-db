package sqlite

import "testing"

func TestSQLiteConnection(t *testing.T) {
	db, err := NewSQLite(":memory:")
	if err != nil {
		t.Fatalf("Failed to initialize SQLite: %v", err)
	}
	defer db.Close()

	_, err = db.DB.Exec("CREATE TABLE IF NOT EXISTS test (id INTEGER PRIMARY KEY, name TEXT)")
	if err != nil {
		t.Fatalf("Failed to execute query on SQLite: %v", err)
	}
}
