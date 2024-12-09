package database

import (
	"fmt"
	"strings"

	"github.com/alessandrolattao/gosqladmin/internal/environment"
)

// generateDSN constructs the DSN string based on the SQL driver and environment
func generateDSN(env *environment.Environment) (string, error) {
	switch strings.ToLower(env.SQLDriver) {
	case "mysql":
		// MySQL DSN with connection, read, and write timeouts
		return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?timeout=%s&readTimeout=%s&writeTimeout=%s",
			env.SQLUser, env.SQLPassword, env.SQLHost, env.SQLPort, env.SQLDatabase,
			env.SQLConnTimeout.String(), env.SQLReadTimeout.String(), env.SQLWriteTimeout.String(),
		), nil
	case "postgres":
		// PostgreSQL DSN with connection, statement, and idle transaction timeouts
		return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s&connect_timeout=%d",
			env.SQLUser, env.SQLPassword, env.SQLHost, env.SQLPort, env.SQLDatabase,
			env.SSLMode,
			int(env.SQLConnTimeout.Seconds()), // Connection timeout in seconds
		), nil
	case "sqlite":
		// SQLite uses the file path as DSN, no native timeout support
		return env.SQLDatabase, nil
	case "sqlserver":
		// SQL Server DSN with connection and query timeouts
		return fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s&connection timeout=%d",
			env.SQLUser, env.SQLPassword, env.SQLHost, env.SQLPort, env.SQLDatabase,
			int(env.SQLConnTimeout.Seconds()), // Connection timeout in seconds
		), nil
	case "snowflake":
		// Snowflake DSN with login and network timeouts
		return fmt.Sprintf("%s:%s@%s:%s/%s?warehouse=%s&database=%s&schema=%s&loginTimeout=%d",
			env.SQLUser, env.SQLPassword, env.SQLHost, env.SQLPort, env.SQLDatabase,
			env.SnowflakeWarehouse, env.SQLDatabase, env.SnowflakeSchema,
			int(env.SQLConnTimeout.Seconds()), // Login timeout in seconds
		), nil
	case "clickhouse":
		// ClickHouse DSN with connection, read, and write timeouts
		return fmt.Sprintf("tcp://%s:%s?username=%s&password=%s&database=%s&dial_timeout=%s&read_timeout=%s&write_timeout=%s",
			env.SQLHost, env.SQLPort, env.SQLUser, env.SQLPassword, env.SQLDatabase,
			env.SQLConnTimeout.String(),  // Connection timeout
			env.SQLReadTimeout.String(),  // Read timeout
			env.SQLWriteTimeout.String(), // Write timeout
		), nil
	default:
		// Unsupported driver error
		return "", fmt.Errorf("unsupported SQL driver: %s", env.SQLDriver)
	}
}
