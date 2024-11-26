package database

import (
	"fmt"

	"github.com/alessandrolattao/gomyadmin/internal/environment"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
)

type DB struct {
	Conn *sqlx.DB
}

// NewDB creates a new database connection
func NewDB(logger zerolog.Logger, env *environment.Environment) (*DB, error) {

	// Construct DSN
	var dsn string
	dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/?timeout=%s&readTimeout=%s&writeTimeout=%s",
		env.MySQLUser,
		env.MySQLPassword,
		env.MySQLHost,
		env.MySQLPort,
		env.MySQLConnTimeout.String(),  // Convert time.Duration to string
		env.MySQLReadTimeout.String(),  // Convert time.Duration to string
		env.MySQLWriteTimeout.String(), // Convert time.Duration to string
	)

	logger.Info().Msg("Connecting to the database")

	// Connect to the database
	db, err := sqlx.Connect("mysql", dsn)
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
