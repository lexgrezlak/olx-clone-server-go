package database

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/jmoiron/sqlx"
	"github.com/ory/dockertest"
	"github.com/sethvargo/go-retry"
	"net"
	"net/url"
	"path/filepath"
	"runtime"
	"testing"
	"time"
)

func NewTestDatabaseWithConfig(tb testing.TB) (*sqlx.DB, *Config) {
	tb.Helper()

	ctx := context.Background()

	// Create the pool (docker instance)
	pool, err := dockertest.NewPool("")
	if err != nil {
		tb.Fatalf("failed to create Docker pool: %s", err)
	}

	// Start the container.
	dbname, user, password := "test-olx-clone-go", "test-user", "test-password"
	container, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "12-alpine",
		Env: []string{
			"LANG=C",
			"POSTGRES_DB=" + dbname,
			"POSTGRES_USER=" + user,
			"POSTGRES_PASSWORD=" + password,
		},
	})
	if err != nil {
		tb.Fatalf("failed to start postgres container: %s", err)
	}

	// Ensure container is cleaned up.
	tb.Cleanup(func() {
		if err := pool.Purge(container); err != nil {
			tb.Fatalf("failed to cleanup postgres container: %s", err)
		}
	})

	// Get the host. On Mac, Docker runs in a VM.
	host := container.Container.NetworkSettings.IPAddress
	if runtime.GOOS == "darwin" {
		host = net.JoinHostPort(container.GetBoundIP("5432/tcp"), container.GetPort("5432/tcp"))
	}

	// Build the connection URL.
	connURL := &url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(user, password),
		Host:   host,
		Path:   dbname,
	}
	q := connURL.Query()
	q.Add("sslmode", "disable")
	connURL.RawQuery = q.Encode()

	// Wait for the container to start - we'll retry connections in a loop
	// below, but there's no point in trying immediately.
	time.Sleep(1 * time.Second)

	b, err := retry.NewFibonacci(500 * time.Millisecond)
	if err != nil {
		tb.Fatalf("failed to configure backoff: %v", err)
	}
	b = retry.WithMaxRetries(10, b)
	b = retry.WithCappedDuration(10*time.Second, b)

	// Establish a connection to the database. Use a Fibonacci backoff
	// instead of exponential so wait times scale appropriately.
	var db *sqlx.DB

	if err := retry.Do(ctx, b, func(ctx context.Context) error {
		var err error
		// Create the db instance.
		db, err = sqlx.Connect("postgres", connURL.String())
		if err != nil {
			return retry.RetryableError(err)
		}
		return nil
	}); err != nil {
		tb.Fatalf("failed to start postgres: %s", err)
	}

	// Run the migrations.
	if err := dbMigrate(connURL.String()); err != nil {
		tb.Fatalf("failed to migrate database: %s", err)
	}

	// Close db when done.
	tb.Cleanup(func() {
		db.Close()
	})

	return db, &Config{
		Name:     dbname,
		User:     user,
		Password: password,
		Host:     container.GetBoundIP("5432/tcp"),
		Port:     container.GetPort("5432/tcp"),
		SSLMode:  "disable",
	}
}

func NewTestDatabase(tb testing.TB) *sqlx.DB {
	tb.Helper()

	db, _ := NewTestDatabaseWithConfig(tb)
	return db
}

// Runs the migrations. u is the connection URL (e.g. postgres://...).
func dbMigrate(u string) error {
	migrationsDir := fmt.Sprintf("file://%s", dbMigrationsDir())
	fmt.Errorf(migrationsDir)
	m, err := migrate.New(migrationsDir, u)
	if err != nil {
		return fmt.Errorf("failed create migrate: %w", err)
	}
	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("failed to run migrate: %w", err)
	}
	srcErr, dbErr := m.Close()
	if srcErr != nil {
		return fmt.Errorf("migrate source error: %w", srcErr)
	}
	if dbErr != nil {
		return fmt.Errorf("migrate database error: %w", dbErr)
	}
	return nil
}

// Return the path on disk to the migrations.
func dbMigrationsDir() string {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		return ""
	}

	return filepath.Join(filepath.Dir(filename), "../../migrations")
}
