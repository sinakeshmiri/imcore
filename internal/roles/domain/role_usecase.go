package domain

import (
	"context"
	"time"
)

type CreateRoleRequest struct {
	RollName    string
	Owner       string
	Description string
}

type RoleUsecase interface {
	Create(c context.Context, req *CreateRoleRequest) error
}
type Role struct {
	Rolename      string
	OwnerUsername string
	Description   string
	IsActive      bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
