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

func (a *applicationUsecase) Approve(ctx context.Context, applicationID string, decisionNote *string) error {
	err := a.applicationRepository.Approve(ctx, applicationID, decisionNote)
	if err != nil {
		return err
	}
	return nil
}

func (a *applicationUsecase) Reject(ctx context.Context, applicationID string, decisionNote *string) error {
	//TODO implement me
	panic("implement me")
}

func (a *applicationUsecase) ListIncoming(c context.Context, user string) ([]*domain.Application, error) {
	incoming, err := a.applicationRepository.ListInComing(c, user)
	if err != nil {
		return nil, err
	}
	return incoming, nil
}

func (a *applicationUsecase) ListOutgoing(c context.Context, user string) ([]*domain.Application, error) {
	outgoing, err := a.applicationRepository.ListOutGoing(c, user)
	if err != nil {
		return nil, err
	}
	return outgoing, nil
}

func (a *applicationUsecase) Create(c context.Context, req *domain.CreateApplicationRequest) (*domain.Application, error) {
	exists, err := a.applicationRepository.ExistsPending(c, req.RoleName, req.ApplicantUsername)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("application already exists")
	}
	app, err := a.applicationRepository.Create(c, req.RoleName, req.ApplicantUsername, req.Reason)
	if err != nil {
		return nil, err
	}
	return &app, nil
}

func NewApplicationUsecase(applicationRepository domain.ApplicationRepository, timeout time.Duration) domain.ApplicationUsecase {
	return &applicationUsecase{
		applicationRepository: applicationRepository,
		contextTimeout:        timeout,
	}
}
