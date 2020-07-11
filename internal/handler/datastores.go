package handler

import (
	"context"
	"olx-clone-server/internal/service"
)

type UserDatastore interface {
	ValidateUser(ctx context.Context, email, password string) (*service.User, error)
	CreateUser(ctx context.Context, input service.SignUpInput) error
}

type PostingDatastore interface {
	CreatePosting(ctx context.Context, input service.CreatePostingInput) error
	GetAllPostings(ctx context.Context) ([]*service.PostingPreview, error)
}
