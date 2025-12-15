package controller

import (
	"context"
	"errors"
	api "github.com/sinakeshmiri/imcore/api/generated"
	"github.com/sinakeshmiri/imcore/domain"
)

type Handler struct {
	uc domain.CreateUserUsecase
}

func NewHandler(uc domain.CreateUserUsecase) *Handler {
	return &Handler{
		uc: uc,
	}
}
func (h *Handler) CreateUser(
	ctx context.Context,
	req api.CreateUserRequestObject,
) (api.CreateUserResponseObject, error) {
	ucReq := domain.CreateUserRequest{
		Email:    string(req.Body.Email),
		Password: req.Body.Password,
	}
	err := h.uc.Create(ctx, &ucReq)
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
