package domain

import (
	"context"
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
