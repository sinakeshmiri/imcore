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
	return api.CreateApplication200JSONResponse(mapApplicationFromDomain(app)), nil
}

func (h *Handler) GetApplication(ctx context.Context, request api.GetApplicationRequestObject) (api.GetApplicationResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (h *Handler) PatchApplication(ctx context.Context, request api.PatchApplicationRequestObject) (api.PatchApplicationResponseObject, error) {
	if request.Body.Status == api.APPROVED {
		err := h.applicationUsecase.Approve(ctx, request.ApplicationId.String(), request.Body.Note)
		if err != nil {
			return nil, err
		}
		return api.PatchApplication201Response{}, nil
	} else if request.Body.Status == api.REJECTED {

	}
	return api.PatchApplication403Response{}, nil

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
