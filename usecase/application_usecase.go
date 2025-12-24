package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/sinakeshmiri/imcore/domain"
)

type applicationUsecase struct {
	applicationRepository domain.ApplicationRepository
	contextTimeout        time.Duration
}

func (a applicationUsecase) ListIncoming(c context.Context, user string) ([]domain.Application, error) {
	//TODO implement me
	panic("implement me")
}

func (a applicationUsecase) ListOutgoing(c context.Context, user string) ([]domain.Application, error) {
	//TODO implement me
	panic("implement me")
}

func (a applicationUsecase) Create(c context.Context, req *domain.CreateApplicationRequest) (domain.Application, error) {
	//check if it exists
	exists, err := a.applicationRepository.ExistsPending(c, req.RoleName, req.ApplicantUsername)
	if err != nil {
		return domain.Application{}, err
	}
	if exists {
		return domain.Application{}, errors.New("application already exists")
	}

	app, err := a.applicationRepository.Create(c, req.RoleName, req.ApplicantUsername, req.Reason)
	if err != nil {
		return domain.Application{}, err
	}
	return app, nil
}

func NewApplicationUsecase(applicationRepository domain.ApplicationRepository, timeout time.Duration) domain.ApplicationUsecase {
	return &applicationUsecase{
		applicationRepository: applicationRepository,
		contextTimeout:        timeout,
	}
}
