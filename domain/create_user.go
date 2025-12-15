package domain

import (
	"context"
	"errors"
)

type CreateUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type CreateUserUsecase interface {
	Create(c context.Context, req *CreateUserRequest) error
}

var (
	ErrDatabaseQueryFailed        = errors.New("failed to execute the query")
	ErrUserAlreadyExists          = errors.New("user already exists")
	ErrPasswordHashCreationFailed = errors.New("failed to calculate password hash")
	ErrInvalidCredentials         = errors.New("invalid credentials")
)
