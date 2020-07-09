package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"strings"
)

type DB struct {
	Pool *pgxpool.Pool
}


// Sets up the database connections using the provided configuration
func NewFromEnv(ctx context.Context, config *Config) (*DB, error) {
	connStr := dbConnectionString(config)

	pool, err := pgxpool.Connect(ctx, connStr)
	if err != nil {
		return nil, fmt.Errorf("creating connection pool: %v", err)
	}
	return &DB{Pool: pool}, nil
}

// Close releases database connections
func (db *DB) Close(ctx context.Context) {
	db.Pool.Close()
}

// dbConnectionString builds a connection string suitable for the pgx
// Postgres driver, using the values of vars
func dbConnectionString(config *Config) string {
	vals := dbValues(config)
	var p []string
	for k, v := range vals {
		p = append(p, fmt.Sprintf("%s=%s", k, v))
	}
	return strings.Join(p, " ")
}

func dbValues(config *Config) map[string]string {
	p := map[string]string{}
	setIfNotEmpty(p, "dbname", config.Name)
	setIfNotEmpty(p, "host", config.Host)
	setIfNotEmpty(p, "user", config.User)
	setIfNotEmpty(p, "password", config.Password)
	setIfNotEmpty(p, "port", config.Port)
	setIfNotEmpty(p, "sslmode", config.SSLMode)
	return p
}

func setIfNotEmpty(m map[string]string, key, val string) {
	if val != "" {
		m[key] = val
	}
}