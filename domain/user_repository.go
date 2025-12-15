package domain

import (
	"context"
	"time"
)

type UserRepository interface {
	Create(c context.Context, user *User) error
	FindByEmail(c context.Context, email string) (*User, error)
}

type User struct {
	Email        string    `gorm:"type:varchar(320);uniqueIndex;not null;primaryKey"`
	PasswordHash string    `gorm:"type:text;not null;column:password_hash"`
	IsActive     bool      `gorm:"not null;default:true;column:is_active"`
	CreatedAt    time.Time `gorm:"not null;default:now();column:created_at"`
	UpdatedAt    time.Time `gorm:"not null;default:now();column:updated_at"`
}

func (User) TableName() string { return "users" }
