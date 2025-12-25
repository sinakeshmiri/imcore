package controller

import (
	"context"

	api "github.com/sinakeshmiri/imcore/api/generated"
	"github.com/sinakeshmiri/imcore/domain"
)

func (h *Handler) ListApplications(ctx context.Context, request api.ListApplicationsRequestObject) (api.ListApplicationsResponseObject, error) {
	outgoing, err := h.applicationUsecase.ListOutgoing(ctx, request.Params.User)
	if err != nil {
		return nil, err
	}
	incoming, err := h.applicationUsecase.ListIncoming(ctx, request.Params.User)
	if err != nil {
		return nil, err
	}
	resOutgoing := make([]api.Application, 0, len(outgoing))
	resIncoming := make([]api.Application, 0, len(incoming))
	for _, entity := range outgoing {
		resOutgoing = append(resOutgoing, mapApplicationFromDomain(entity))
	}
	for _, entity := range incoming {
		resIncoming = append(resIncoming, mapApplicationFromDomain(entity))
	}
	return api.ListApplications200JSONResponse{
		Incoming: &resIncoming,
		Outgoing: &resOutgoing,
	}, nil
}

func (h *Handler) CreateApplication(ctx context.Context, request api.CreateApplicationRequestObject) (api.CreateApplicationResponseObject, error) {

	app, err := h.applicationUsecase.Create(ctx, &domain.CreateApplicationRequest{
		RoleName:          request.Body.Rolename,
		ApplicantUsername: *request.Body.ApplicantUsername,
		Reason:            *request.Body.Reason,
	})
	if err != nil {
		return nil, err
	}
	status := app.Status.String()

	return api.CreateApplication200JSONResponse{
		ApplicantUsername: &app.ApplicantUsername,
		CreatedAt:         &app.CreatedAt,
		DecidedAt:         app.DecidedAt,
		DecisionNote:      &app.DecisionNote,
		Id:                &app.ID,
		OwnerUsername:     &app.OwnerUsername,
		Reason:            &app.Reason,
		Rolename:          &app.Rolename,
		Status:            &status,
	}, nil
}

func (h *Handler) GetApplication(ctx context.Context, request api.GetApplicationRequestObject) (api.GetApplicationResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (h *Handler) PatchApplication(ctx context.Context, request api.PatchApplicationRequestObject) (api.PatchApplicationResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func mapApplicationFromDomain(entity *domain.Application) api.Application {
	status := entity.Status.String()

	return api.Application{
		ApplicantUsername: &entity.ApplicantUsername,
		CreatedAt:         &entity.CreatedAt,
		DecidedAt:         entity.DecidedAt,
		DecisionNote:      &entity.DecisionNote,
		Id:                &entity.ID,
		OwnerUsername:     &entity.OwnerUsername,
		Reason:            &entity.Reason,
		Rolename:          &entity.Rolename,
		Status:            &status,
	}
}
