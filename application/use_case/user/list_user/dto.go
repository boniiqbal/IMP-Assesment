package list_user

import (
	"imp-backend/domain"
	"time"

	"github.com/go-playground/validator/v10"
)

type ListUserRequest struct {
	Page  int `json:"page" form:"page" query:"page"`
	Limit int `json:"limit" form:"limit" query:"limit"`
}

type ListUserResponse struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Fullname  string    `json:"fullname"`
	CreatedAt time.Time `json:"created_at"`
}

func (c *ListUserRequest) Validate() error {
	validate := validator.New()
	err := validate.Struct(c)
	if err != nil {
		return err
	}
	return nil
}

func MappingResponse(s []domain.Users) []ListUserResponse {
	result := []ListUserResponse{}
	for _, val := range s {
		result = append(result, ListUserResponse{
			ID:        val.ID,
			Fullname:  val.FullName,
			Username:  val.Username,
			CreatedAt: val.CreatedAt,
		})
	}
	return result
}
