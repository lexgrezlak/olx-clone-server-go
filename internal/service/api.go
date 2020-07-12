package service

import "github.com/jmoiron/sqlx"

type API struct {
	Db *sqlx.DB
}

type UserDatastore interface {
	ValidateUser(email, password string) (*User, error)
	CreateUser(input SignUpInput) error
}

type PostingDatastore interface {
	CreatePosting(input CreatePostingInput) error
	GetAllPostings() ([]*PostingPreview, error)
}
