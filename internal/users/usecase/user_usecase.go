package usecase

import (
	"context"
	"log"
	"time"

	domain "github.com/sinakeshmiri/authon-core/internal/users/domain"
	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func (cu *userUsecase) Create(ctx context.Context, req *domain.CreateUserRequest) error {
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
		Fullname:     req.Fullname,
		Username:     req.Username,
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
func (cu *userUsecase) ListRoles(c context.Context, username string) ([]string, error) {
	roles, err := cu.userRepository.ListRoles(c, username)
	if err != nil {
		return nil, err
	}
	return roles, nil
}

func NewUserUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}
