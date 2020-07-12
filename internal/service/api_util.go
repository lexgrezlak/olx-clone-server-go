package service

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
)

func NewTestAPI() (*api, sqlmock.Sqlmock, error) {
	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}
	// DB of type *sqlx.DB is needed.
	sqlxDB := sqlx.NewDb(sqlDB, "sqlmock")

	return &api{db: sqlxDB}, mock, nil
}
