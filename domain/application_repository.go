package domain

import (
	"context"
	"time"
)

type ApplicationRepository interface {
	Create(c context.Context, app *Application) error
	GetByID(ctx context.Context, id string) (*Application, error)
	ListOutGoing(c context.Context, id string) ([]*Application, error)
	ListInComing(c context.Context, id string) ([]*Application, error)
	UpdateStatus(c context.Context, id string, status ApplicationStatus) error
	Exists(c context.Context, role string, username string) (bool, error)
}

type Application struct {
	ID                string `gorm:"type:uuid;primaryKey"`
	Rolename          string `gorm:"type:varchar(64);not null;index"`
	ApplicantUsername string `gorm:"type:varchar(64);not null;index"`
	OwnerUsername     string `gorm:"type:varchar(64);not null;index"`

	Status       string `gorm:"type:varchar(16);not null;index"`
	Reason       string `gorm:"type:varchar(640)"`
	DecisionNote string `gorm:"type:varchar(640)"`

	CreatedAt time.Time  `gorm:"not null;default:now();column:created_at"`
	DecidedAt *time.Time `gorm:"column:decided_at"`
}

func (Application) TableName() string { return "applications" }
