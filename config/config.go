package config

type DatabaseConfig struct {
	Driver string // e.g., "sqlite" or "postgres"
	DSN    string // Connection string
}
