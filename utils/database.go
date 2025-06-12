package utils

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sort"
	"strings"
)

// RunMigrations executes database migration scripts
func RunMigrations(db *sql.DB) error {
	// Create migrations table if it doesn't exist
	err := createMigrationsTable(db)
	if err != nil {
		return fmt.Errorf("failed to create migrations table: %w", err)
	}

	// Get list of applied migrations
	appliedMigrations, err := getAppliedMigrations(db)
	if err != nil {
		return fmt.Errorf("failed to get applied migrations: %w", err)
	}

	// Get list of migration files
	migrationFiles, err := getMigrationFiles("migrations")
	if err != nil {
		return fmt.Errorf("failed to get migration files: %w", err)
	}

	// Apply new migrations
	for _, file := range migrationFiles {
		migrationName := strings.TrimSuffix(filepath.Base(file), ".sql")
		
		// Skip if already applied
		if contains(appliedMigrations, migrationName) {
			continue
		}

		err = applyMigration(db, file, migrationName)
		if err != nil {
			return fmt.Errorf("failed to apply migration %s: %w", migrationName, err)
		}

		fmt.Printf("Applied migration: %s\n", migrationName)
	}

	return nil
}

func createMigrationsTable(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS schema_migrations (
			id SERIAL PRIMARY KEY,
			version VARCHAR(255) UNIQUE NOT NULL,
			applied_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`
	_, err := db.Exec(query)
	return err
}

func getAppliedMigrations(db *sql.DB) ([]string, error) {
	query := "SELECT version FROM schema_migrations ORDER BY version"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var migrations []string
	for rows.Next() {
		var version string
		if err := rows.Scan(&version); err != nil {
			return nil, err
		}
		migrations = append(migrations, version)
	}

	return migrations, nil
}

func getMigrationFiles(dir string) ([]string, error) {
	files, err := filepath.Glob(filepath.Join(dir, "*.sql"))
	if err != nil {
		return nil, err
	}

	sort.Strings(files)
	return files, nil
}

func applyMigration(db *sql.DB, filePath, migrationName string) error {
	// Read migration file
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	// Start transaction
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Split content by statements (simple approach)
	statements := strings.Split(string(content), ";")
	
	for _, stmt := range statements {
		stmt = strings.TrimSpace(stmt)
		if stmt == "" {
			continue
		}

		_, err = tx.Exec(stmt)
		if err != nil {
			return fmt.Errorf("failed to execute statement: %s, error: %w", stmt, err)
		}
	}

	// Record migration as applied
	_, err = tx.Exec("INSERT INTO schema_migrations (version) VALUES ($1)", migrationName)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

// DatabaseHealthCheck checks if database is accessible
func DatabaseHealthCheck(db *sql.DB) error {
	return db.Ping()
}

// GetDatabaseStats returns database connection statistics
func GetDatabaseStats(db *sql.DB) map[string]interface{} {
	stats := db.Stats()
	return map[string]interface{}{
		"open_connections":     stats.OpenConnections,
		"in_use":              stats.InUse,
		"idle":                stats.Idle,
		"wait_count":          stats.WaitCount,
		"wait_duration":       stats.WaitDuration.String(),
		"max_idle_closed":     stats.MaxIdleClosed,
		"max_idle_time_closed": stats.MaxIdleTimeClosed,
		"max_lifetime_closed": stats.MaxLifetimeClosed,
	}
}

// CreateIndex creates a database index if it doesn't exist
func CreateIndex(db *sql.DB, tableName, indexName string, columns []string) error {
	query := fmt.Sprintf(
		"CREATE INDEX IF NOT EXISTS %s ON %s (%s)",
		indexName,
		tableName,
		strings.Join(columns, ", "),
	)
	
	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create index %s: %w", indexName, err)
	}
	
	return nil
}

// DropIndex drops a database index if it exists
func DropIndex(db *sql.DB, indexName string) error {
	query := fmt.Sprintf("DROP INDEX IF EXISTS %s", indexName)
	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to drop index %s: %w", indexName, err)
	}
	
	return nil
}

// TableExists checks if a table exists in the database
func TableExists(db *sql.DB, tableName string) (bool, error) {
	query := `
		SELECT EXISTS (
			SELECT FROM information_schema.tables 
			WHERE table_schema = 'public'
			AND table_name = $1
		)
	`
	
	var exists bool
	err := db.QueryRow(query, tableName).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("failed to check if table exists: %w", err)
	}
	
	return exists, nil
}

// ExecuteInTransaction executes a function within a database transaction
func ExecuteInTransaction(db *sql.DB, fn func(*sql.Tx) error) error {
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	
	err = fn(tx)
	return err
}
