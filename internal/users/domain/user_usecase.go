package domain

import (
	"context"
	"time"
)

type CreateUserRequest struct {
	Email    string
	Fullname string
	Username string
	Password string
}

type UserUsecase interface {
	Create(c context.Context, req *CreateUserRequest) error
	ListRoles(c context.Context, username string) ([]string, error)
}

type User struct {
	Username     string
	Email        string
	Fullname     string
	PasswordHash string
	IsActive     bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
