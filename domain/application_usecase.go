package domain

import (
	"context"
	"fmt"
)

type ApplicationStatus int

const (
	Pending ApplicationStatus = iota
	Canceled
	Rejected
	Approved
)

type CreateApplicationRequest struct {
	RoleName          string
	ApplicantUsername string
	Reason            string
}

func (s ApplicationStatus) String() string {
	return [...]string{"PENDING", "CANCELED", "REJECTED", "APPROVED"}[s]
}

func ParseStatus(str string) (ApplicationStatus, error) {
	switch str {
	case "PENDING":
		return Pending, nil
	case "CANCELED":
		return Canceled, nil
	case "REJECTED":
		return Rejected, nil
	case "APPROVED":
		return Approved, nil
	default:
		return Pending, fmt.Errorf("invalid status: %s", str)
	}
}

type ApplicationUsecase interface {
	Create(c context.Context, req *CreateApplicationRequest) error
	ListIncoming(c context.Context, user string) ([]Application, error)
	ListOutgoing(c context.Context, user string) ([]Application, error)
}
