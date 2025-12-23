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
	ID                string
	Rolename          string
	ApplicantUsername string
	Status            ApplicationStatus
	Reason            string
	DecisionNote      string
	CreatedAt         time.Time
	DecidedAt         *time.Time
}
