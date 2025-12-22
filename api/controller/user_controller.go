package controller

import (
	"context"
	"errors"

	api "github.com/sinakeshmiri/imcore/api/generated"
	"github.com/sinakeshmiri/imcore/domain"
)

func (h *Handler) CreateUser(
	ctx context.Context,
	req api.CreateUserRequestObject,
) (api.CreateUserResponseObject, error) {
	ucReq := domain.CreateUserRequest{
		Fullname: req.Body.Fullname,
		Username: req.Body.Username,
		Email:    string(req.Body.Email),
		Password: req.Body.Password,
	}
	err := h.userUsecase.Create(ctx, &ucReq)
	if err == nil {
		return api.CreateUser201Response{}, nil
	}
	switch {
	case errors.Is(err, domain.ErrUserAlreadyExists):
		return api.CreateUser400Response{}, nil
	default:
		return api.CreateUser500Response{}, err
	}
}

func (h *Handler) DeleteUser(ctx context.Context, request api.DeleteUserRequestObject) (api.DeleteUserResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (h *Handler) GetUser(ctx context.Context, request api.GetUserRequestObject) (api.GetUserResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (h *Handler) UpdateUser(ctx context.Context, request api.UpdateUserRequestObject) (api.UpdateUserResponseObject, error) {
	//TODO implement me
	panic("implement me")
}
