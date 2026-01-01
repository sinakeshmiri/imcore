package repository

import (
	"context"
	"errors"
	"time"

	"github.com/sinakeshmiri/authon-core/internal/roles/domain"
	"gorm.io/gorm"
)

type RoleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{db: db}
}

func (r *RoleRepository) Create(ctx context.Context, role *domain.Role) error {
	roleModel := fromDomain(role)
	return r.db.WithContext(ctx).Create(roleModel).Error
}

func (r *RoleRepository) FindByName(ctx context.Context, name string) (*domain.Role, error) {
	var role RoleModel
	err := r.db.WithContext(ctx).Where("rolename = ?", name).First(&role).Error
	//TODO: handle when owner doesn't exist
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	domainRole := role.toDomain()
	return domainRole, nil
}

type RoleModel struct {
	Rolename      string    `gorm:"type:varchar(320);uniqueIndex;not null;primaryKey"`
	OwnerUsername string    `gorm:"type:varchar(320);not null"`
	Description   string    `gorm:"type:varchar(640);"`
	IsActive      bool      `gorm:"not null;default:true;column:is_active"`
	CreatedAt     time.Time `gorm:"not null;default:now();column:created_at"`
	UpdatedAt     time.Time `gorm:"not null;default:now();column:updated_at"`
}

func (r *RoleModel) toDomain() *domain.Role {
	return &domain.Role{
		Rolename:      r.Rolename,
		OwnerUsername: r.OwnerUsername,
		Description:   r.Description,
		IsActive:      r.IsActive,
		CreatedAt:     r.CreatedAt,
		UpdatedAt:     r.UpdatedAt,
	}
}

func fromDomain(domain *domain.Role) *RoleModel {
	return &RoleModel{
		Rolename:      domain.Rolename,
		OwnerUsername: domain.OwnerUsername,
		Description:   domain.Description,
		IsActive:      domain.IsActive,
		CreatedAt:     domain.CreatedAt,
		UpdatedAt:     domain.UpdatedAt,
	}
}

func (r *RoleModel) TableName() string { return "roles" }
