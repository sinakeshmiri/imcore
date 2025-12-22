package usecase

import (
	"context"
	"time"

	"github.com/sinakeshmiri/imcore/domain"
)

type applicationUsecase struct {
	applicationRepository domain.ApplicationRepository
	contextTimeout        time.Duration
}

func (a applicationUsecase) Create(c context.Context, req *domain.CreateApplicationRequest) error {
	//TODO implement me
	panic("implement me")
}

func NewApplicationUsecase(applicationRepository domain.ApplicationRepository, timeout time.Duration) domain.ApplicationUsecase {
	return &applicationUsecase{
		applicationRepository: applicationRepository,
		contextTimeout:        timeout,
	}
}
