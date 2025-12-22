package domain

import "context"

type ApplicationStatus int

const (
	PENDING ApplicationStatus = iota
	CANCELED
	REJECTED
	APPROVED
)

type CreateApplicationRequest struct {
	RoleName          string
	ApplicantUsername string
	Reason            string
}
type ApplicationUsecase interface {
	Create(c context.Context, req *CreateApplicationRequest) error
}
