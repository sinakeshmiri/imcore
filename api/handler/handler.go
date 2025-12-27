package handler

import (
	"context"

	api "github.com/sinakeshmiri/authon-core/api/generated"
	appshttp "github.com/sinakeshmiri/authon-core/internal/applications/http"
	roleshttp "github.com/sinakeshmiri/authon-core/internal/roles/http"
	usershttp "github.com/sinakeshmiri/authon-core/internal/users/http"
)

type APIHandler struct {
	Users        *usershttp.Handler
	Roles        *roleshttp.Handler
	Applications *appshttp.Handler
}

func (h *APIHandler) CreateUser(ctx context.Context, req api.CreateUserRequestObject) (api.CreateUserResponseObject, error) {
	return h.Users.CreateUser(ctx, req)
}
func (h *APIHandler) GetUser(ctx context.Context, req api.GetUserRequestObject) (api.GetUserResponseObject, error) {
	return h.Users.GetUser(ctx, req)
}

func (h *APIHandler) CreateRole(ctx context.Context, req api.CreateRoleRequestObject) (api.CreateRoleResponseObject, error) {
	return h.Roles.CreateRole(ctx, req)
}
func (h *APIHandler) GetRole(ctx context.Context, req api.GetRoleRequestObject) (api.GetRoleResponseObject, error) {
	return h.Roles.GetRole(ctx, req)
}

func (h *APIHandler) CreateApplication(ctx context.Context, req api.CreateApplicationRequestObject) (api.CreateApplicationResponseObject, error) {
	return h.Applications.CreateApplication(ctx, req)
}
func (h *APIHandler) ListApplications(ctx context.Context, req api.ListApplicationsRequestObject) (api.ListApplicationsResponseObject, error) {
	return h.Applications.ListApplications(ctx, req)
}
func (h *APIHandler) GetApplication(ctx context.Context, req api.GetApplicationRequestObject) (api.GetApplicationResponseObject, error) {
	return h.Applications.GetApplication(ctx, req)
}
func (h *APIHandler) PatchApplication(ctx context.Context, req api.PatchApplicationRequestObject) (api.PatchApplicationResponseObject, error) {
	return h.Applications.PatchApplication(ctx, req)
}
func (h *APIHandler) DeleteRole(ctx context.Context, request api.DeleteRoleRequestObject) (api.DeleteRoleResponseObject, error) {
	return h.Roles.DeleteRole(ctx, request)
}

func (h *APIHandler) UpdateRole(ctx context.Context, request api.UpdateRoleRequestObject) (api.UpdateRoleResponseObject, error) {
	return h.Roles.UpdateRole(ctx, request)
}

func (h *APIHandler) DeleteUser(ctx context.Context, request api.DeleteUserRequestObject) (api.DeleteUserResponseObject, error) {
	return h.Users.DeleteUser(ctx, request)
}

func (h *APIHandler) UpdateUser(ctx context.Context, request api.UpdateUserRequestObject) (api.UpdateUserResponseObject, error) {
	return h.Users.UpdateUser(ctx, request)
}

func (h *APIHandler) ListUserRoles(ctx context.Context, request api.ListUserRolesRequestObject) (api.ListUserRolesResponseObject, error) {
	return h.Users.ListUserRoles(ctx, request)
}
