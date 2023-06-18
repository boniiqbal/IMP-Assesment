package list_user

import (
	"context"
	"imp-backend/application/infrastructure"
	"imp-backend/application/misc"
	"imp-backend/domain"
)

type ListUserService struct {
	authRepository infrastructure.AuthRepository
}

func NewListUserService(
	accountBankRepo infrastructure.AuthRepository,
) ListUserService {
	return ListUserService{
		authRepository: accountBankRepo,
	}
}

func (s *ListUserService) ListUser(ctx context.Context, params *ListUserRequest) ([]ListUserResponse, error) {
	// get user
	res, err := s.authRepository.SelectUser(ctx, &domain.UserParams{
		BasedFilter: domain.BasedFilter{
			Page:  params.Page,
			Limit: params.Limit,
		},
	})
	if err != nil {
		misc.LogEf("ListUserService - ListUser error : ", err)
		return nil, err
	}

	return MappingResponse(res), nil
}
