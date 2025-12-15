package usecase

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"

	"github.com/sinakeshmiri/imcore/domain"
)

type creatUserUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func (cu creatUserUsecase) Create(ctx context.Context, req *domain.CreateUserRequest) error {
	byEmail, err := cu.userRepository.FindByEmail(ctx, req.Email)
	if err != nil {
		log.Printf("failed to check if the user already exists or not %s\n", err)
		return domain.ErrDatabaseQueryFailed
	}
	if byEmail != nil {
		return domain.ErrUserAlreadyExists
	}
	// TODO: check password format
	hashBytes, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		log.Printf("failed to calculate password hash for user %s : %s\n", req.Email, err)
		return domain.ErrPasswordHashCreationFailed
	}
	passwordHash := string(hashBytes)
	user := domain.User{
		Email:        req.Email,
		PasswordHash: passwordHash,
		IsActive:     true,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	err = cu.userRepository.Create(ctx, &user)
	if err != nil {
		log.Printf("failed to insert user: %s\n", err)
		return domain.ErrDatabaseQueryFailed
	}
	return nil
}

func NewCreateUserUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.CreateUserUsecase {
	return &creatUserUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}
