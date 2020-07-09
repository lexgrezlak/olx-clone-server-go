package service

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"olx-clone-server/internal/config"
	"strings"
)

type DB struct {
	Pool *pgxpool.Pool
}

// Sets up the config connections using the provided configuration
func NewDB(config *config.Config) (*DB, error) {
	connStr := dbConnectionString(config)

	pool, err := pgxpool.Connect(context.Background(), connStr)
	if err != nil {
		return nil, fmt.Errorf("creating connection pool: %v", err)
	}
	return &DB{Pool: pool}, nil
}


// dbConnectionString builds a connection string suitable for the pgx
// Postgres driver, using the values of vars
func dbConnectionString(config *config.Config) string {
	vals := dbValues(config)
	var p []string
	for k, v := range vals {
		p = append(p, fmt.Sprintf("%s=%s", k, v))
	}
	return strings.Join(p, " ")
}

func dbValues(config *config.Config) map[string]string {
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