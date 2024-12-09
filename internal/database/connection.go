package database

import (
	"fmt"
	"time"

	"github.com/alessandrolattao/gosqladmin/internal/environment"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"

	// Import supported drivers
	_ "github.com/ClickHouse/clickhouse-go" // ClickHouse driver
	_ "github.com/denisenkom/go-mssqldb"    // SQL Server driver
	_ "github.com/go-sql-driver/mysql"      // MySQL driver
	_ "github.com/lib/pq"                   // PostgreSQL driver
	_ "github.com/mattn/go-sqlite3"         // SQLite driver
	_ "github.com/snowflakedb/gosnowflake"  // Snowflake driver
)

// DB wraps the sqlx.DB connection
type DB struct {
	Conn *sqlx.DB
}

// NewConnection establishes a database connection with retry logic
func NewConnection(logger zerolog.Logger, env *environment.Environment) (*DB, error) {
	logger.Info().Msgf("Initializing connection for driver: %s", env.SQLDriver)

	// Generate the DSN string based on the driver
	dsn, err := generateDSN(env)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to generate DSN")
		return nil, err
	}

	// Attempt to establish the connection with retries
	db, err := connectWithRetries(logger, env, dsn)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to connect to the database")
		return nil, err
	}

	// Configure the connection pool
	configureConnectionPool(db, env)

	logger.Info().Msg("Successfully connected to the database")
	return &DB{Conn: db}, nil
}

func connectWithRetries(logger zerolog.Logger, env *environment.Environment, dsn string) (*sqlx.DB, error) {
	maxRetries := 10
	backoff := 3 * time.Second

	var db *sqlx.DB
	var err error

	for attempts := 0; attempts < maxRetries; attempts++ {
		db, err = sqlx.Connect(env.SQLDriver, dsn)
		if err == nil {
			return db, nil
		}

		logger.Error().Err(err).Msgf("Connection failed (attempt %d/%d), retrying in %d seconds...", attempts+1, maxRetries, backoff)
		time.Sleep(backoff)
	}

	return nil, fmt.Errorf("all connection attempts failed: %w", err)
}
