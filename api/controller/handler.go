package controller

import (
	"github.com/sinakeshmiri/imcore/domain"
)

func NewHandler(userUsecase domain.UserUsecase, roleUsecase domain.RoleUsecase, applicationUsecase domain.ApplicationUsecase) *Handler {
	return &Handler{
		userUsecase:        userUsecase,
		roleUsecase:        roleUsecase,
		applicationUsecase: applicationUsecase,
	}
}

type Handler struct {
	userUsecase        domain.UserUsecase
	roleUsecase        domain.RoleUsecase
	applicationUsecase domain.ApplicationUsecase
}
