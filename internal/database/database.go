package database

import (
	"fmt"

	"example.com/databases/internal/database/postgres"
	"example.com/databases/internal/database/sqlite"
)

type DBHandler interface {
	Close() error
}

func NewDatabase(driver, dsn string) (DBHandler, error) {
	switch driver {
	case "sqlite":
		return sqlite.NewSQLite(dsn)
	case "postgres":
		return postgres.NewPostgres(dsn)
	default:
		return nil, fmt.Errorf("unsupported database driver: %s", driver)
	}
}
