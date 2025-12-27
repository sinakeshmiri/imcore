package repository

import (
	"context"
	"errors"

	"github.com/sinakeshmiri/imcore/domain"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, u *domain.User) error {
	return r.db.WithContext(ctx).Create(u).Error
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	var u domain.User
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&u).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &u, nil
}

func (r *UserRepository) ListRoles(ctx context.Context, username string) ([]string, error) {
	var roles []string

	err := r.db.WithContext(ctx).
		Table("user_roles").
		Select("rolename").
		Where("username = ?", username).
		Order("created_at DESC").
		Scan(&roles).Error

	if err != nil {
		return nil, err
	}

	return roles, nil
}
