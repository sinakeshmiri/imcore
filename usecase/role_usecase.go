package usecase

import (
	"context"
	"log"
	"time"

	"github.com/sinakeshmiri/imcore/domain"
)

type roleUsecase struct {
	roleRepository domain.RoleRepository
	contextTimeout time.Duration
}

func (ru roleUsecase) Create(ctx context.Context, req *domain.CreateRoleRequest) error {
	byName, err := ru.roleRepository.FindByName(ctx, req.RollName)
	if err != nil {
		log.Printf("failed to check if the role already exists or not %s\n", err)
		return domain.ErrDatabaseQueryFailed
	}
	if byName != nil {
		return domain.ErrRoleAlreadyExists
	}

	role := domain.Role{
		Rolename:      req.RollName,
		Description:   req.Description,
		OwnerUsername: req.Owner,
		IsActive:      true,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
	err = ru.roleRepository.Create(ctx, &role)
	if err != nil {
		log.Printf("failed to insert role: %s\n", err)
		return domain.ErrDatabaseQueryFailed
	}
	return nil
}

func NewRoleUsecase(roleRepository domain.RoleRepository, timeout time.Duration) domain.RoleUsecase {
	return &roleUsecase{
		roleRepository: roleRepository,
		contextTimeout: timeout,
	}
}
