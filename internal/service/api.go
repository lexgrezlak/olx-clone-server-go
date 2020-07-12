package service

import "github.com/jmoiron/sqlx"

type api struct {
	db *sqlx.DB
}

type UserDatastore interface {
	ValidateUser(email, password string) (*User, error)
	CreateUser(input SignUpInput) error
}

type PostingDatastore interface {
	CreatePosting(input CreatePostingInput) error
	GetAllPostings() ([]*PostingPreview, error)
}

func NewAPI(db *sqlx.DB) *api {
	return &api{db: db}
}