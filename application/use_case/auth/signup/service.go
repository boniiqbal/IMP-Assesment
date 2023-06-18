package signup

import (
	"context"
	"errors"
	"imp-backend/application/infrastructure"
	"imp-backend/application/misc"
	"imp-backend/domain"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type SignupService struct {
	authRepository infrastructure.AuthRepository
}

func NewSignupService(
	authRepo infrastructure.AuthRepository,
) SignupService {
	return SignupService{
		authRepository: authRepo,
	}
}

func (s *SignupService) Signup(ctx context.Context, params *SignupRequest) error {
	// GetUser by username
	user, err := s.authRepository.GetUser(ctx, &domain.Users{
		Username: params.Username,
	})
	if err != nil {
		misc.LogEf("SignupService - Signup error : ", err)
		return err
	}

	if user != nil {
		return errors.New("account already exist")
	}

	// Hashing password user
	hashedPassword, errHash := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if errHash != nil {
		return errHash
	}

	now := time.Now()
	_, err = s.authRepository.CreateUser(ctx, &domain.Users{
		CreatedAt: now,
		Password:  string(hashedPassword),
		Username:  params.Username,
		FullName:  params.Fullname,
	})
	if err != nil {
		misc.LogEf("SignupService - Login error : ", err)
		return err
	}

	return nil
}
