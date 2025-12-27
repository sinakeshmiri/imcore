package domain

import (
	"context"
	"time"
)

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	FindByEmail(ctx context.Context, email string) (*User, error)
	ListRoles(ctx context.Context, username string) ([]string, error)
}

type User struct {
	Username     string    `gorm:"type:varchar(320);uniqueIndex;not null;primaryKey"`
	Email        string    `gorm:"type:varchar(320);uniqueIndex;not null"`
	Fullname     string    `gorm:"type:varchar(320);not null"`
	PasswordHash string    `gorm:"type:text;not null;column:password_hash"`
	IsActive     bool      `gorm:"not null;default:true;column:is_active"`
	CreatedAt    time.Time `gorm:"not null;default:now();column:created_at"`
	UpdatedAt    time.Time `gorm:"not null;default:now();column:updated_at"`
}

func (User) TableName() string { return "users" }
