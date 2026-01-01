package repository

import (
	"context"
	"errors"
	"time"

	"github.com/sinakeshmiri/authon-core/internal/users/domain"
	"gorm.io/gorm"
)

type UserModel struct {
	Username     string    `gorm:"type:varchar(320);uniqueIndex;not null;primaryKey"`
	Email        string    `gorm:"type:varchar(320);uniqueIndex;not null"`
	Fullname     string    `gorm:"type:varchar(320);not null"`
	PasswordHash string    `gorm:"type:text;not null;column:password_hash"`
	IsActive     bool      `gorm:"not null;default:true;column:is_active"`
	CreatedAt    time.Time `gorm:"not null;default:now();column:created_at"`
	UpdatedAt    time.Time `gorm:"not null;default:now();column:updated_at"`
}

func (UserModel) TableName() string { return "users" }
func toDomain(m UserModel) domain.User {
	return domain.User{
		Username:     m.Username,
		Email:        m.Email,
		Fullname:     m.Fullname,
		PasswordHash: m.PasswordHash,
		IsActive:     m.IsActive,
		CreatedAt:    m.CreatedAt,
		UpdatedAt:    m.UpdatedAt,
	}
}
func fromDomain(d domain.User) UserModel {
	return UserModel{
		Username:     d.Username,
		Email:        d.Email,
		Fullname:     d.Fullname,
		PasswordHash: d.PasswordHash,
		IsActive:     d.IsActive,
		CreatedAt:    d.CreatedAt,
		UpdatedAt:    d.UpdatedAt,
	}
}

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
	var u UserModel
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&u).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	domainUser := toDomain(u)
	return &domainUser, nil
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
