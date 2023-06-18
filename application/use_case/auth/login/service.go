package login

import (
	"context"
	"errors"
	"imp-backend/application/infrastructure"
	"imp-backend/application/misc"
	"imp-backend/domain"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/dgrijalva/jwt-go"
)

type LoginService struct {
	authRepository infrastructure.AuthRepository
}

func NewLoginService(
	authRepo infrastructure.AuthRepository,
) LoginService {
	return LoginService{
		authRepository: authRepo,
	}
}

func (s *LoginService) Login(ctx context.Context, params *LoginRequest) (*LoginResponse, error) {
	// getUser by username
	user, err := s.authRepository.GetUser(ctx, &domain.Users{
		Username: params.Username,
	})
	if err != nil {
		misc.LogEf("LoginService - Login error : ", err)
		return nil, err
	}
	if user == nil {
		return nil, errors.New("username is not found")
	}

	byteDBPass := []byte(user.Password)
	password := []byte(params.Password)

	// Compare password user
	if error := bcrypt.CompareHashAndPassword(byteDBPass, password); error != nil {
		return nil, errors.New("password is not correct")
	}

	// assign claim value
	now := time.Now()
	end := now.Add(time.Duration(168 * time.Hour))
	claim := domain.AccessTokenClaim{}
	claim.AccountID = user.ID
	claim.AccountName = user.FullName
	claim.ExpiresAt = end.Unix()
	claim.IssuedAt = now.Unix()

	newToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claim)
	tokenString, err := newToken.SignedString([]byte("secret"))
	if err != nil {
		return nil, err
	}

	return &LoginResponse{
		AccountID: user.ID,
		Token:     tokenString,
		CreatedAt: now,
		ExpiredAt: end,
		FullName:  user.FullName,
	}, nil
}
