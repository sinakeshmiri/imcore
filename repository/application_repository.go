package repository

import (
	"context"

	"github.com/sinakeshmiri/imcore/domain"
	"gorm.io/gorm"
)

type ApplicationRepository struct {
	db *gorm.DB
}

func (a ApplicationRepository) Create(c context.Context, app *domain.Application) error {
	//TODO implement me
	panic("implement me")
}

func (a ApplicationRepository) GetByID(ctx context.Context, id string) (*domain.Application, error) {
	//TODO implement me
	panic("implement me")
}

func (a ApplicationRepository) ListOutGoing(c context.Context, id string) ([]*domain.Application, error) {
	//TODO implement me
	panic("implement me")
}

func (a ApplicationRepository) ListInComing(c context.Context, id string) ([]*domain.Application, error) {
	//TODO implement me
	panic("implement me")
}

func (a ApplicationRepository) UpdateStatus(c context.Context, id string, status domain.ApplicationStatus) error {
	//TODO implement me
	panic("implement me")
}

func (a ApplicationRepository) Exists(c context.Context, role string, username string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func NewApplicationRepository(db *gorm.DB) *ApplicationRepository {
	return &ApplicationRepository{db: db}
}
