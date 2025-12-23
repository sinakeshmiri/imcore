package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
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

func (a applicationUsecase) Create(c context.Context, req *domain.CreateApplicationRequest) error {
	//check if it exists
	exists, err := a.applicationRepository.Exists(c, req.RoleName, req.ApplicantUsername)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("application already exists")
	}
	appId, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	err = a.applicationRepository.Create(c, &domain.Application{
		ID:                appId.String(),
		Rolename:          req.RoleName,
		ApplicantUsername: req.ApplicantUsername,
		Status:            domain.Pending,
		Reason:            req.Reason,
		DecisionNote:      "",
		CreatedAt:         time.Now(),
		DecidedAt:         nil,
	})
	if err != nil {
		return err
	}
	return nil
}

func NewApplicationUsecase(applicationRepository domain.ApplicationRepository, timeout time.Duration) domain.ApplicationUsecase {
	return &applicationUsecase{
		applicationRepository: applicationRepository,
		contextTimeout:        timeout,
	}
}
