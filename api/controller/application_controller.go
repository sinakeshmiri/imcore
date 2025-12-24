package controller

import (
	"context"

	api "github.com/sinakeshmiri/imcore/api/generated"
	"github.com/sinakeshmiri/imcore/domain"
)

func (h *Handler) ListApplications(ctx context.Context, request api.ListApplicationsRequestObject) (api.ListApplicationsResponseObject, error) {
	//TODO implement me
	panic("implement me")
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
		DecidedAt:         nil,
		DecisionNote:      nil,
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
