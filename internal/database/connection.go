package database

import (
	"fmt"
	"strings"

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

type DB struct {
	Conn *sqlx.DB
}

// NewConnection creates a new database connection
func NewConnection(logger zerolog.Logger, env *environment.Environment) (*DB, error) {

	// Construct DSN based on driver
	var dsn string
	logger.Info().Msgf("Initializing connection for driver: %s", env.SQLDriver)

	switch strings.ToLower(env.SQLDriver) {
	case "mysql":
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?timeout=%s&readTimeout=%s&writeTimeout=%s",
			env.SQLUser,
			env.SQLPassword,
			env.SQLHost,
			env.SQLPort,
			env.SQLDatabase,
			env.SQLConnTimeout.String(),
			env.SQLReadTimeout.String(),
			env.SQLWriteTimeout.String(),
		)
	case "postgres":
		dsn = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
			env.SQLUser,
			env.SQLPassword,
			env.SQLHost,
			env.SQLPort,
			env.SQLDatabase,
			env.SSLMode,
		)
	case "sqlite":
		dsn = env.SQLDatabase // SQLite uses the file path as DSN
	case "sqlserver":
		dsn = fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
			env.SQLUser,
			env.SQLPassword,
			env.SQLHost,
			env.SQLPort,
			env.SQLDatabase,
		)
	case "snowflake":
		dsn = fmt.Sprintf("%s:%s@%s:%s/%s?warehouse=%s&database=%s&schema=%s",
			env.SQLUser,
			env.SQLPassword,
			env.SQLHost,
			env.SQLPort,
			env.SQLDatabase,
			env.SnowflakeWarehouse,
			env.SQLDatabase,
			env.SnowflakeSchema,
		)
	case "clickhouse":
		dsn = fmt.Sprintf("tcp://%s:%s?username=%s&password=%s&database=%s",
			env.SQLHost,
			env.SQLPort,
			env.SQLUser,
			env.SQLPassword,
			env.SQLDatabase,
		)
	default:
		err := fmt.Errorf("unsupported SQL driver: %s", env.SQLDriver)
		logger.Error().Err(err).Msg("Failed to configure DSN")
		return nil, err
	}

	// Log connection info (without credentials for security)
	logger.Info().Msg("Connecting to the database...")

	// Connect to the database
	db, err := sqlx.Connect(env.SQLDriver, dsn)
	if err != nil {
		logger.Error().Err(err).Msg("Error connecting to the database")
		return nil, err
	}

	// Configure the connection pool
	db.SetMaxOpenConns(env.MaxOpenConns)
	db.SetMaxIdleConns(env.MaxIdleConns)
	db.SetConnMaxLifetime(env.ConnMaxLifetime)

	logger.Info().Msg("Successfully connected to the database")
	return &DB{Conn: db}, nil
}
