package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// InitDatabase initializes and returns a database connection
func InitDatabase(cfg *Config) (*sql.DB, error) {
	var dsn string
	
	// Use DATABASE_URL if provided, otherwise construct from individual components
	if cfg.DatabaseURL != "" {
		dsn = cfg.DatabaseURL
	} else {
		dsn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			cfg.DatabaseHost,
			cfg.DatabasePort,
			cfg.DatabaseUser,
			cfg.DatabasePass,
			cfg.DatabaseName,
		)
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	// Set connection pool settings
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)

	return db, nil
}
