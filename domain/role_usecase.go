package domain

import (
	"context"
)

type CreateRoleRequest struct {
	RollName    string
	Owner       string
	Description string
}

type RoleUsecase interface {
	Create(c context.Context, req *CreateRoleRequest) error
}
