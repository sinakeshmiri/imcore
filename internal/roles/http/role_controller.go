package http

import (
	"context"
	"errors"

	api "github.com/sinakeshmiri/authon-core/api/generated"
	"github.com/sinakeshmiri/authon-core/internal/roles/domain"
)

type Handler struct {
	roleUsecase domain.RoleUsecase
}

func NewHandler(roleUsecase domain.RoleUsecase) *Handler {
	return &Handler{
		roleUsecase: roleUsecase,
	}
}
func (h *Handler) CreateRole(
	ctx context.Context,
	req api.CreateRoleRequestObject,
) (api.CreateRoleResponseObject, error) {
	ucReq := domain.CreateRoleRequest{
		Owner:       req.Body.Owner,
		RollName:    req.Body.Rolename,
		Description: req.Body.Description,
	}
	err := h.roleUsecase.Create(ctx, &ucReq)
	if err == nil {
		return api.CreateRole201Response{}, nil
	}
	switch {
	case errors.Is(err, domain.ErrRoleAlreadyExists):
		return api.CreateRole400Response{}, nil
	default:
		return api.CreateRole500Response{}, err
	}
}

func (h *Handler) DeleteRole(ctx context.Context, request api.DeleteRoleRequestObject) (api.DeleteRoleResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (h *Handler) GetRole(ctx context.Context, request api.GetRoleRequestObject) (api.GetRoleResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (h *Handler) UpdateRole(ctx context.Context, request api.UpdateRoleRequestObject) (api.UpdateRoleResponseObject, error) {
	//TODO implement me
	panic("implement me")
}
