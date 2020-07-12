package service

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
)

func NewTestAPI() (*api, *sql.DB, sqlmock.Sqlmock, error) {
	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		fmt.Printf("an error occurred while creating mock DB connection: %s", err)
		return nil, nil,nil, err
	}
	// DB of type *sqlx.DB is needed.
	sqlxDB := sqlx.NewDb(sqlDB, "sqlmock")

	return &api{db: sqlxDB}, sqlDB, mock, nil
}
