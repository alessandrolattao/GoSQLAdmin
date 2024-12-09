package database

import (
	"github.com/alessandrolattao/gosqladmin/internal/environment"
	"github.com/jmoiron/sqlx"
)

// configureConnectionPool sets the database connection pool parameters
func configureConnectionPool(db *sqlx.DB, env *environment.Environment) {
	db.SetMaxOpenConns(env.MaxOpenConns)
	db.SetMaxIdleConns(env.MaxIdleConns)
	db.SetConnMaxLifetime(env.ConnMaxLifetime)
}
