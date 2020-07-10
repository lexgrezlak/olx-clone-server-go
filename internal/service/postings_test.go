package service

import (
	"github.com/DATA-DOG/go-sqlmock"
	"testing"
)

func TestDB_CreatePosting(t *testing.T) {
	db, mock, err := sqlmock.New()


	if err != nil {
		t.Fatalf("an unexpected error when opening a stub database connection: %s", err)
	}
	defer db.Close()

	mock.ExpectExec("INSERT INTO public.posting").WillReturnResult(sqlmock.NewResult(1, 1))

}
