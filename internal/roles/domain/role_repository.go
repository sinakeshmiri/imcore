package domain

import (
	"context"
)

type RoleRepository interface {
	Create(c context.Context, role *Role) error
	FindByName(c context.Context, name string) (*Role, error)
}
